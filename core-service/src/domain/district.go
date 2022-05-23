package domain

import (
	"context"
	"time"
)

// District ...
type District struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Chief     string     `json:"chief"`
	Address   string     `json:"address"`
	Website   string     `json:"website"`
	Logo      NullString `json:"logo"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// District Usecase ...
type DistrictUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]District, int64, error)
}

// District Repository ...
type DistrictRepository interface {
	Fetch(ctx context.Context, params *Request) ([]District, int64, error)
}
