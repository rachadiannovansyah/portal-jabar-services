package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type authUsecase struct {
	config         *config.Config
	userRepo       domain.UserRepository
	unitRepo       domain.UnitRepository
	roleRepo       domain.RoleRepository
	contextTimeout time.Duration
}

// NewAuthUsecase will create new an authUsecase object representation of domain.AuthUsecase interface
func NewAuthUsecase(c *config.Config, u domain.UserRepository, un domain.UnitRepository, r domain.RoleRepository, timeout time.Duration) domain.AuthUsecase {
	return &authUsecase{
		config:         c,
		userRepo:       u,
		unitRepo:       un,
		roleRepo:       r,
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

func (n *authUsecase) createAccessToken(user *domain.User) (accessToken string, exp int64, err error) {
	exp = time.Now().Add(time.Second * n.config.JWT.TTL).Unix()
	claims := &domain.JwtCustomClaims{
		ID:    user.ID,
		Email: user.Email,
		Unit:  user.Unit,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * n.config.JWT.RefreshTTL).Unix(),
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
			ExpiresAt: time.Now().Add(time.Second * n.config.JWT.RefreshTTL).Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)

	t, err = refreshToken.SignedString([]byte(n.config.JWT.RefreshSecret))

	return
}

func (n *authUsecase) Login(c context.Context, req *domain.LoginRequest) (res domain.LoginResponse, err error) {

	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	user, err := n.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return domain.LoginResponse{}, domain.ErrInvalidCredentials
	}

	accessToken, exp, err := n.createAccessToken(&user)
	if err != nil {
		return domain.LoginResponse{}, err
	}

	refreshToken, err := n.createRefreshToken(&user)

	// write last active
	timeNow := time.Now()
	err = n.userRepo.WriteLastActive(c, timeNow, &user)
	if err != nil {
		return domain.LoginResponse{}, err
	}

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

func (n *authUsecase) UserProfile(c context.Context, id uuid.UUID) (res domain.User, err error) {

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
