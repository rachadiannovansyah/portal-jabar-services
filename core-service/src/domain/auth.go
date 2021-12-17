package domain

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JwtCustomClaims struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	jwt.StandardClaims
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
	UserProfile(ctx context.Context, id uuid.UUID) (User, error)
}
