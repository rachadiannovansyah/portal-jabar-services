package domain

import (
	"context"
	"time"
)

// News ...
type News struct {
	ID        int64      `json:"id"`
	Title     NullString `json:"title" validate:"required"`
	Excerpt   NullString `json:"excerpt"`
	Content   NullString `json:"content" validate:"required"`
	Slug      NullString `json:"slug"`
	Image     NullString `json:"image"`
	Video     NullString `json:"video"`
	Source    NullString `json:"source"`
	Status    NullString `json:"status,omitempty"`
	Views     int64      `json:"views"`
	Highlight int8       `json:"highlight,omitempty"`
	Category  Category   `json:"category" validate:"required"`
	CreatedBy NullString `json:"created_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// NewsListResponse ...
type NewsListResponse struct {
	ID        int64      `json:"id"`
	Title     NullString `json:"title"`
	Excerpt   NullString `json:"excerpt"`
	Slug      NullString `json:"slug"`
	Image     NullString `json:"image"`
	Category  Category   `json:"category"`
	CreatedBy NullString `json:"created_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// NewsUsecase represent the news usecases
type NewsUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]News, int64, error)
	GetByID(ctx context.Context, id int64) (News, error)
}

// NewsRepository represent the news repository contract
type NewsRepository interface {
	Fetch(ctx context.Context, params *Request) (new []News, total int64, err error)
	GetByID(ctx context.Context, id int64) (News, error)
	AddView(ctx context.Context, id int64) (err error)
}
