package domain

import (
	"context"
	"time"
)

// Unit ...
type Unit struct {
	ID          int64      `json:"id"`
	ParentID    NullInt64  `json:"parent_id"`
	Name        NullString `json:"name" validate:"required"`
	Description NullString `json:"description"`
	Logo        NullString `json:"logo"`
	Website     NullString `json:"website"`
	Phone       NullString `json:"phone"`
	Address     NullString `json:"address"`
	Chief       NullString `json:"chief"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// UnitInfo ...
type UnitInfo struct {
	ID   int64      `json:"id"`
	Name NullString `json:"name" validate:"required"`
}

// UnitUsecase represent the unit usecases
type UnitUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]Unit, int64, error)
	GetByID(ctx context.Context, id int64) (Unit, error)
}

// UnitRepository represent the unit repository contract
type UnitRepository interface {
	Fetch(ctx context.Context, params *Request) (new []Unit, total int64, err error)
	GetByID(ctx context.Context, id int64) (Unit, error)
}
