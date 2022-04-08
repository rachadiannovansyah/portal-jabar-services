package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

type userUsecase struct {
	userRepo          domain.UserRepository
	unitRepo          domain.UnitRepository
	roleRepo          domain.RoleRepository
	mailTemplateRepo  domain.TemplateRepository
	regInvitationRepo domain.RegistrationInvitationRepository
	contextTimeout    time.Duration
}

// NewUserUsecase creates a new user usecase
func NewUserUsecase(u domain.UserRepository, un domain.UnitRepository, r domain.RoleRepository,
	m domain.TemplateRepository, i domain.RegistrationInvitationRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:          u,
		unitRepo:          un,
		roleRepo:          r,
		mailTemplateRepo:  m,
		regInvitationRepo: i,
		contextTimeout:    timeout,
	}
}

func encryptPassword(password string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}

	return string(encryptedPassword), nil
}

func (n *userUsecase) isValidUser(ctx context.Context, req *domain.CheckPasswordRequest, id uuid.UUID) (ok bool, err error) {
	user, err := n.userRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return
	}

	return
}

func (u *userUsecase) Store(c context.Context, usr *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	// generate uuid v4
	usr.ID = uuid.New()

	encryptedPassword, err := encryptPassword(usr.Password)
	if err != nil {
		return err
	}

	usr.Password = string(encryptedPassword)

	err = u.userRepo.Store(ctx, usr)
	return
}

func (u *userUsecase) UpdateProfile(c context.Context, req *domain.User) (user domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err = u.GetByID(ctx, req.ID)
	if err != nil {
		return
	}

	if req.Nip != nil && *user.Nip != *req.Nip {
		if res, _ := u.CheckIfNipExists(ctx, req.Nip); res {
			return user, domain.ErrDuplicateNIP
		}
	}

	// FIXME: make some utility function to separate this code
	if req.Name != "" {
		user.Name = req.Name
	}

	if req.Username != "" {
		user.Username = req.Username
	}

	if req.Email != "" {
		user.Email = req.Email
	}

	if req.Nip != nil {
		user.Nip = req.Nip
	}

	if req.Occupation != nil {
		user.Occupation = req.Occupation
	}

	err = u.userRepo.Update(ctx, &user)
	return
}

func (n *userUsecase) GetByID(c context.Context, id uuid.UUID) (res domain.User, err error) {

	ctx, cancel := context.WithTimeout(c, n.contextTimeout)

	res, err = n.userRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	unit, _ := n.unitRepo.GetByID(ctx, res.Unit.ID)
	role, _ := n.roleRepo.GetByID(ctx, res.Role.ID)

	res.Unit = helpers.GetUnitInfo(unit)
	res.Role = helpers.GetRoleInfo(role)

	defer cancel()

	return
}

func (n *userUsecase) ChangePassword(c context.Context, id uuid.UUID, req *domain.ChangePasswordRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	user, err := n.userRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.CurrentPassword))
	if err != nil {
		return err
	}

	encryptedPassword, err := encryptPassword(req.NewPassword)
	if err != nil {
		return err
	}

	currentTime := time.Now()
	user.Password = string(encryptedPassword)
	user.LastPasswordChanged = &currentTime
	err = n.userRepo.Update(ctx, &user)

	return
}

func (n *userUsecase) AccountSubmission(c context.Context, id uuid.UUID, key string) (res domain.AccountSubmission, err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	user, err := n.userRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	template, err := n.mailTemplateRepo.GetByTemplate(ctx, key)
	if err != nil {
		return
	}

	go func() {
		err = helpers.SendEmail(user.Email, template, []string{user.Name, user.UnitName})
		if err != nil {
			return
		}
	}()

	return
}

func (u *userUsecase) RegisterByInvitation(c context.Context, req *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	regInvitation, err := u.regInvitationRepo.GetByToken(ctx, req.Token)
	if err != nil {
		return
	}

	if err = helpers.IsInvitationTokenValid(regInvitation, req.Token); err != nil {
		return
	}

	// check nip is not used
	nip := *req.Nip
	if exists, _ := u.CheckIfNipExists(ctx, req.Nip); exists {
		return domain.ErrDuplicateNIP
	}

	encryptedPassword, err := encryptPassword(req.Password)
	if err != nil {
		return err
	}

	occupation := *req.Occupation
	payload := &domain.User{
		ID:         uuid.New(),
		Name:       req.Name,
		Username:   regInvitation.Email,
		Email:      regInvitation.Email,
		Nip:        &nip,
		Occupation: &occupation,
		Unit:       domain.UnitInfo{ID: regInvitation.UnitID},
		Role:       domain.RoleInfo{ID: domain.RoleContributor},
		Password:   string(encryptedPassword),
	}

	err = u.userRepo.Store(ctx, payload)
	if err != nil {
		return
	}

	err = u.regInvitationRepo.Delete(ctx, *regInvitation.ID)

	return
}

func (u *userUsecase) UserList(ctx context.Context, params *domain.Request) (res []domain.User, total int64, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	res, total, err = u.userRepo.UserList(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	// looping user to assign roles object
	for key, user := range res {
		role, err := u.roleRepo.GetByID(ctx, user.Role.ID)
		if err != nil {
			return nil, 0, err
		}

		res[key].Role = helpers.GetRoleInfo(role)
	}

	return
}

// GetUserByID will find an object by given id
func (u *userUsecase) GetUserByID(c context.Context, id uuid.UUID) (res domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return
	}

	role, _ := u.roleRepo.GetByID(ctx, res.Role.ID)
	res.Role = helpers.GetRoleInfo(role)

	return
}

func (u *userUsecase) CheckIfNipExists(c context.Context, nip *string) (res bool, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err := u.userRepo.GetByNip(ctx, nip)

	if err != nil {
		return
	}

	if user.Nip != nil {
		res = true
	}

	return
}

func (n *userUsecase) SetAsAdmin(c context.Context, id uuid.UUID, req *domain.CheckPasswordRequest, userID uuid.UUID) (err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	_, err = n.isValidUser(ctx, req, id)
	if err != nil {
		return
	}

	err = n.userRepo.SetAsAdmin(ctx, userID, domain.RoleAdministrator)

	return
}

func (n *userUsecase) ChangeEmail(c context.Context, id uuid.UUID, req *domain.CheckPasswordRequest, userID uuid.UUID) (err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	_, err = n.isValidUser(ctx, req, id)
	if err != nil {
		return
	}

	newEmail := req.NewEmail
	if _, ok := helpers.IsValidMailAddress(newEmail); !ok {
		return errors.New("invalid email format")
	}

	if u, _ := n.userRepo.GetByEmail(ctx, newEmail); u.Email != "" {
		return errors.New("email already registered")
	}

	if reg, _ := n.regInvitationRepo.GetByEmail(ctx, newEmail); reg.Email != "" {
		return errors.New("email already registered on invitation")
	}

	err = n.userRepo.ChangeEmail(ctx, userID, newEmail)

	return
}

func (n *userUsecase) ChangeStatus(c context.Context, id uuid.UUID, req *domain.CheckPasswordRequest, userID uuid.UUID) (err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	_, err = n.isValidUser(ctx, req, id)
	if err != nil {
		return
	}

	err = n.userRepo.ChangeStatus(ctx, userID, req.Status)

	return
}
