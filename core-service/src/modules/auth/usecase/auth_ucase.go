package usecase

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type authUsecase struct {
	config         *config.Config
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

// NewAuthUsecase will create new an authUsecase object representation of domain.AuthUsecase interface
func NewAuthUsecase(c *config.Config, u domain.UserRepository, timeout time.Duration) domain.AuthUsecase {
	return &authUsecase{
		config:         c,
		userRepo:       u,
		contextTimeout: timeout,
	}
}

func newLoginResponse(token, refreshToken string, exp int64) domain.LoginResponse {
	return domain.LoginResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
		Exp:          exp,
	}
}

func (n *authUsecase) Login(c context.Context, req *domain.LoginRequest) (res domain.LoginResponse, err error) {

	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	user, err := n.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return domain.LoginResponse{}, fmt.Errorf("Invalid credentials")
	}

	accessToken, exp, err := n.createAccessToken(&user)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	refreshToken, err := n.createRefreshToken(&user)

	res = newLoginResponse(accessToken, refreshToken, exp)

	return
}

func (n *authUsecase) RefreshToken(c context.Context, req *domain.RefreshRequest) (res domain.LoginResponse, err error) {

	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	// claim refresh token first
	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(n.config.JWT.RefreshSecret), nil
	})

	if err != nil {
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return domain.LoginResponse{}, fmt.Errorf("invalid token")
	}

	user, err := n.userRepo.GetByEmail(ctx, claims["email"].(string))
	if err != nil {
		return domain.LoginResponse{}, err
	}

	accessToken, exp, err := n.createAccessToken(&user)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	refreshToken, err := n.createRefreshToken(&user)

	res = newLoginResponse(accessToken, refreshToken, exp)

	return
}

func (n *authUsecase) createAccessToken(user *domain.User) (accessToken string, exp int64, err error) {
	exp = time.Now().Add(time.Hour * n.config.JWT.ExpireCount).Unix()
	claims := &domain.JwtCustomClaims{
		user.ID,
		user.Name,
		user.Email,
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
		ID:    user.ID,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * n.config.JWT.ExpireRefreshCount).Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)

	t, err = refreshToken.SignedString([]byte(n.config.JWT.RefreshSecret))

	return
}
