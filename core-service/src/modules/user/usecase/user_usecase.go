package usecase

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type userUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

// NewUserUsecase creates a new user usecase
func NewUserkUsecase(u domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepo:       u,
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
