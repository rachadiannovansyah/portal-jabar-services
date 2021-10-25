package domain

import "context"

type Auth struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

// AuthUsecase ...
type AuthUsecase interface {
	Login(ctx context.Context, req *LoginRequest) (Auth, error)
}
