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
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type PublicServiceRepository interface {
	Fetch(ctx context.Context, params *Request) ([]PublicService, error)
}
