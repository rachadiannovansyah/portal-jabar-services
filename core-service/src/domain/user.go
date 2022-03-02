package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// User ...
type User struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Nip        *string   `json:"nip" validate:"omitempty,len=18"`
	Occupation *string   `json:"occupation" validate:"omitempty,max=35"`
	Photo      *string   `json:"photo"`
	Unit       UnitInfo  `json:"unit"`
	UnitName   string    `json:"unit_name"`
	Role       RoleInfo  `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

// UserInfo ...
type UserInfo struct {
	ID         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	Username   string     `json:"username"`
	Email      string     `json:"email"`
	Nip        NullString `json:"nip"`
	Occupation NullString `json:"occupation"`
	Photo      NullString `json:"photo"`
	Unit       UnitInfo   `json:"unit"`
	Role       RoleInfo   `json:"role"`
}

// Author ...
type Author struct {
	Name     string `json:"name"`
	UnitName string `json:"unit_name"`
}

// UserRepository represent the unit repository contract
type UserRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	Store(context.Context, *User) error
	Update(context.Context, *User) error
}

// UserUsecase ...
type UserUsecase interface {
	Store(context.Context, *User) error
	UpdateProfile(context.Context, *User) error
}
