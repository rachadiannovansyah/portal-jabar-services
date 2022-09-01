package domain

import (
	"context"
	"time"
)

// PublicService ...
type PublicService struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Description NullString `json:"description"`
	Unit        NullString `json:"unit"`
	Url         NullString `json:"url"`
	Image       NullString `json:"image"`
	Category    NullString `json:"category"`
	IsActive    NullString `json:"is_active"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type StorePserviceRequest struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
	Image       string    `json:"image"`
	Category    string    `json:"category"`
	IsActive    int8      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// PublicServiceUsecase ...
type PublicServiceUsecase interface {
	Store(context.Context, *StorePserviceRequest) error
}

// PublicServiceRepository ...
type PublicServiceRepository interface {
	Fetch(ctx context.Context, params *Request) ([]PublicService, error)
	Store(ctx context.Context, params *StorePserviceRequest) error
}
