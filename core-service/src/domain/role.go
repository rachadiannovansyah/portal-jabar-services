package domain

import (
	"context"
	"time"
)

// Role ...
type Role struct {
	ID          int8      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RoleInfo struct {
	ID   int8   `json:"id"`
	Name string `json:"name"`
}

// RoleRepository represent the unit repository contract
type RoleRepository interface {
	GetByID(ctx context.Context, id int8) (Role, error)
}
