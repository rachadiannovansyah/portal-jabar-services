package domain

import (
	"context"
	"time"
)

type RegistrationInvitation struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegistrationInvitationClaim struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type RegistrationInvitationRepository interface {
	GetByEmail(ctx context.Context, email string) (RegistrationInvitation, error)
	GetByToken(ctx context.Context, token string) (RegistrationInvitation, error)
	Store(ctx context.Context, invitation *RegistrationInvitation) error
	Update(ctx context.Context, id int64, invitation *RegistrationInvitation) error
}

type RegistrationInvitationUsecase interface {
	Invite(ctx context.Context, email string) (RegistrationInvitation, error)
	Authorize(ctx context.Context, token string) (RegistrationInvitationClaim, error)
}
