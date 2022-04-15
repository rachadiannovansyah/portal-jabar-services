package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Constanta User Status
const (
	ActiveUser   string = "ACTIVE"
	InactiveUser string = "INACTIVE"
	PendingUser  string = "PENDING"
)

// User ...
type User struct {
	ID                  uuid.UUID  `json:"id"`
	Name                string     `json:"name" validate:"omitempty,required,max=100"`
	Username            string     `json:"username" validate:"omitempty,required,max=62"`
	Email               string     `json:"email" validate:"omitempty,required,max=64"`
	Password            string     `json:"password"`
	LastPasswordChanged *time.Time `json:"last_password_changed"`
	LastActive          *time.Time `json:"last_active"`
	Status              string     `json:"status"`
	Nip                 *string    `json:"nip" validate:"omitempty,len=0|len=18"`
	Occupation          *string    `json:"occupation" validate:"omitempty,max=35"`
	Photo               *string    `json:"photo" validate:"omitempty,max=255"`
	Unit                UnitInfo   `json:"unit"`
	UnitName            string     `json:"unit_name"`
	Role                RoleInfo   `json:"role"`
	Token               string     `json:"token"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	DeletedAt           time.Time  `json:"deleted_at"`
}

// UserInfo ...
type UserInfo struct {
	ID                  uuid.UUID  `json:"id"`
	Name                string     `json:"name"`
	Username            string     `json:"username"`
	Email               string     `json:"email"`
	Nip                 *string    `json:"nip"`
	Occupation          *string    `json:"occupation"`
	Photo               *string    `json:"photo"`
	Unit                UnitInfo   `json:"unit"`
	Role                RoleInfo   `json:"role"`
	Status              string     `json:"status"`
	LastActive          *time.Time `json:"last_active"`
	LastPasswordChanged *time.Time `json:"last_password_changed"`
}

// UserListResponse ...
type UserListResponse struct {
	ID         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Role       RoleInfo   `json:"role"`
	LastActive *time.Time `json:"last_active"`
	Status     string     `json:"status"`
}

// AccountSubmission ...
type AccountSubmission struct {
	ID   int64  `json:"id"`
	Role string `json:"role"`
}

// Author ...
type Author struct {
	Name     string `json:"name"`
	UnitName string `json:"unit_name"`
}

// ChangePasswordRequest ...
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required"`
}

type CheckPasswordRequest struct {
	Password string `json:"password" validate:"required"`
	NewEmail string `json:"new_email" validate:"omitempty,required"`
	Status   string `json:"status" validate:"omitempty,required,eq=ACTIVE|eq=INACTIVE|eq=PENDING"`
}

type CheckNipExistRequest struct {
	Nip *string `json:"nip" validate:"len=18"`
}

// UserRepository represent the unit repository contract
type UserRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (User, error)
	GetByNip(ctx context.Context, nip *string) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	Store(context.Context, *User) error
	Update(context.Context, *User) error
	WriteLastActive(context.Context, time.Time, *User) error
	UserList(context.Context, *Request) ([]User, int64, error)
	SetAsAdmin(context.Context, uuid.UUID, int8) error
	ChangeEmail(context.Context, uuid.UUID, string) error
	ChangeStatus(context.Context, uuid.UUID, string) error
}

// UserUsecase ...
type UserUsecase interface {
	Store(context.Context, *User) error
	GetByID(ctx context.Context, id uuid.UUID) (User, error)
	CheckIfNipExists(ctx context.Context, nip *string) (bool, error)
	UpdateProfile(context.Context, *User) (User, error)
	ChangePassword(context.Context, uuid.UUID, *ChangePasswordRequest) error
	AccountSubmission(context.Context, uuid.UUID, string) (AccountSubmission, error)
	RegisterByInvitation(ctx context.Context, user *User) error // domain mana yach?
	UserList(ctx context.Context, params *Request) (res []User, total int64, err error)
	SetAsAdmin(context.Context, uuid.UUID, *CheckPasswordRequest, uuid.UUID) error
	ChangeEmail(context.Context, uuid.UUID, *CheckPasswordRequest, uuid.UUID) error
	ChangeStatus(context.Context, uuid.UUID, *CheckPasswordRequest, uuid.UUID) error
}
