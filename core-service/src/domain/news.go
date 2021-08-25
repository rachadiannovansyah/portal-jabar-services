package domain

import (
	"context"
	"time"
)

// News ...
type News struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title" validate:"required"`
	Excerpt   string     `json:"excerpt"`
	Content   string     `json:"content" validate:"required"`
	Slug      string     `json:"slug"`
	Image     NullString `json:"image"`
	Video     NullString `json:"video"`
	Source    NullString `json:"source"`
	Status    string     `json:"status,omitempty"`
	Category  Category   `json:"category_id" validate:"required"`
	CreatedBy NullString `json:"created_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// NewsListResponse ...
type NewsListResponse struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Excerpt   string     `json:"excerpt"`
	Slug      string     `json:"slug"`
	Image     NullString `json:"image"`
	Category  Category   `json:"category_id"`
	CreatedBy NullString `json:"created_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// FetchNewsRequest penggunaan pointer ini agar dapat memberikan value nil jika tidak digunakan
type FetchNewsRequest struct {
	Keyword string
	PerPage int64
	Offset  int64
	OrderBy string
	SortBy  string
}

// NewsUsecase represent the news usecases
type NewsUsecase interface {
	Fetch(ctx context.Context, params *FetchNewsRequest) ([]News, int64, error)
	GetByID(ctx context.Context, id int64) (News, error)
}

// NewsRepository represent the news repository contract
type NewsRepository interface {
	Fetch(ctx context.Context, params *FetchNewsRequest) (new []News, total int64, err error)
	GetByID(ctx context.Context, id int64) (News, error)
}
