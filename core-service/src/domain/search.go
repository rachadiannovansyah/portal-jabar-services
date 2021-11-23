package domain

import (
	"context"
	"time"
)

// Search ...
type Search struct {
	ID        int       `json:"id"`
	Domain    string    `json:"domain"`
	Title     string    `json:"title"`
	Excerpt   string    `json:"excerpt"`
	Content   string    `json:"content"`
	Slug      string    `json:"slug"`
	Category  string    `json:"category" validate:"required"`
	Thumbnail string    `json:"thumbnail"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SearchListResponse ...
type SearchListResponse struct {
	ID        int       `json:"id"`
	Domain    string    `json:"domain"`
	Title     string    `json:"title"`
	Excerpt   string    `json:"excerpt"`
	Slug      string    `json:"slug"`
	Category  string    `json:"category" validate:"required"`
	Thumbnail string    `json:"thumbnail"`
	CreatedAt time.Time `json:"created_at"`
}

// SearchUsecase represent the search usecases
type SearchUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]SearchListResponse, int64, error)
}

// SearchRepository represent the search repository contract
type SearchRepository interface {
	Fetch(ctx context.Context, params *Request) (new []SearchListResponse, total int64, err error)
}
