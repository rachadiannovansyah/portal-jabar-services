package domain

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JwtCustomClaims struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Unit        UnitInfo  `json:"unit"`
	Role        RoleInfo  `json:"role"`
	Permissions []string  `json:"permissions"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	jwt.StandardClaims
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshRequest struct {
	Token string `json:"token" validate:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Exp          int64  `json:"exp"`
}

// AuthUsecase ...
type AuthUsecase interface {
	Login(ctx context.Context, req *LoginRequest) (LoginResponse, error)
	RefreshToken(ctx context.Context, req *RefreshRequest) (LoginResponse, error)
	GetPermissionsByRoleID(ctx context.Context, id int8) ([]string, error)
}
