package usecase

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type authUsecase struct {
	config         *config.Config
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

// NewAuthUsecase will create new an authUsecase object representation of domain.AuthUsecase interface
func NewAuthUsecase(cfg *config.Config, u domain.UserRepository, timeout time.Duration) domain.AuthUsecase {
	return &authUsecase{
		config:         cfg,
		userRepo:       u,
		contextTimeout: timeout,
	}
}

func (n *authUsecase) Login(c context.Context, req *domain.LoginRequest) (res domain.LoginResponse, err error) {

	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	user, err := n.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	accessToken, exp, err := n.createAccessToken(&user)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	refreshToken, err := n.createRefreshToken(&user)

	res = domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Exp:          exp,
	}

	return
}

func (n *authUsecase) createAccessToken(user *domain.User) (accessToken string, exp int64, err error) {
	exp = time.Now().Add(time.Hour * n.config.JWT.ExpireCount).Unix()
	claims := &domain.JwtCustomClaims{
		user.Name,
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = token.SignedString([]byte(n.config.JWT.AccessSecret))

	return
}

func (n *authUsecase) createRefreshToken(user *domain.User) (t string, err error) {
	claimsRefresh := &domain.JwtCustomRefreshClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * n.config.JWT.ExpireRefreshCount).Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)

	t, err = refreshToken.SignedString([]byte(n.config.JWT.RefreshSecret))

	return
}
