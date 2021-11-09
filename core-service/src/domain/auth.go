package domain

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JwtCustomClaims struct {
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID uuid.UUID `json:"id"`
	jwt.StandardClaims
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Exp          int64  `json:"exp"`
}

// AuthUsecase ...
type AuthUsecase interface {
	Login(ctx context.Context, req *LoginRequest) (LoginResponse, error)
}
