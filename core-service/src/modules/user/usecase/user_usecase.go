package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	unitRepo       domain.UnitRepository
	roleRepo       domain.RoleRepository
	contextTimeout time.Duration
}

// NewUserUsecase creates a new user usecase
func NewUserkUsecase(u domain.UserRepository, un domain.UnitRepository, r domain.RoleRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       u,
		unitRepo:       un,
		roleRepo:       r,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) Store(c context.Context, usr *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	// generate uuid v4
	usr.ID = uuid.New()

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(usr.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}
	usr.Password = string(encryptedPassword)

	err = u.userRepo.Store(ctx, usr)
	return
}

func (u *userUsecase) UpdateProfile(c context.Context, req *domain.User) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err := u.userRepo.GetByID(ctx, req.ID)
	if err != nil {
		return err
	}

	user.Name = req.Name
	user.Nip = req.Nip
	user.Occupation = req.Occupation

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
