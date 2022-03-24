package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type RegistrationInvitation struct {
	ID        *uuid.UUID `json:"id"`
	Email     string     `json:"email"`
	Token     string     `json:"token"`
	UnitID    int64      `json:"unit_id"`
	InvitedBy uuid.UUID  `json:"created_by"`
	InvitedAt time.Time  `json:"created_at"`
}

type RegistrationInvitationClaim struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type RegistrationInvitationRepository interface {
	GetByEmail(ctx context.Context, email string) (RegistrationInvitation, error)
	GetByToken(ctx context.Context, token string) (RegistrationInvitation, error)
	Store(ctx context.Context, invitation *RegistrationInvitation) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, invitation *RegistrationInvitation) error
}

type RegistrationInvitationUsecase interface {
	Invite(context.Context, RegistrationInvitation) (RegistrationInvitation, error)
	Authorize(ctx context.Context, token string) (RegistrationInvitationClaim, error)
}
