package usecase

import (
	"context"
	"fmt"
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

	var RoleContributor int8 = 4
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
		Role:       domain.RoleInfo{ID: RoleContributor},
		Password:   string(encryptedPassword),
	}

	err = u.userRepo.Store(ctx, payload)
	if err != nil {
		return
	}

	err = u.regInvitationRepo.Delete(ctx, *regInvitation.ID)

	return
}

func (u *userUsecase) CheckIfNipExists(c context.Context, nip *string) (res bool, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err := u.userRepo.GetByNip(ctx, nip)
	fmt.Println("user", user)
	if err != nil {
		return
	}

	if user.Nip != nil {
		res = true
	}

	return
}
