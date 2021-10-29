package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type authUsecase struct {
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

// NewAuthUsecase will create new an authUsecase object representation of domain.AuthUsecase interface
func NewAuthUsecase(u domain.UserRepository, timeout time.Duration) domain.AuthUsecase {
	return &authUsecase{
		userRepo:       u,
		contextTimeout: timeout,
	}
}

func (n *authUsecase) Login(c context.Context, req *domain.LoginRequest) (res domain.Auth, err error) {

	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	user, err := n.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return domain.Auth{}, err
	}

	res, err = createAccessToken(user)

	return
}

func createAccessToken(u domain.User) (res domain.Auth, err error) {
	res = domain.Auth{
		AccessToken: "23843294u2xn32nxb23yx48732y8$",
		TokenType:   "bearer",
		ExpiresIn:   20,
	}
	return
}
