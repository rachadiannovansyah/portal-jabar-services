package domain

import (
	"context"
	"time"
)

// News ...
type News struct {
	ID             int64      `json:"id"`
	NewsCategoryID int64      `json:"newsCategoryId" validate:"required"`
	Title          string     `json:"title" validate:"required"`
	Content        string     `json:"content" validate:"required"`
	Slug           string     `json:"slug"`
	Image          NullString `json:"image"`
	Video          NullString `json:"video"`
	Source         NullString `json:"source"`
	ShowDate       string     `json:"showDate"`
	EndDate        string     `json:"endDate"`
	Status         string     `json:"status"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
}

// ListNews ...
type ListNews struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Excerpt   string     `json:"excerpt"`
	Category  int64      `json:"category"`
	Slug      string     `json:"slug"`
	Image     NullString `json:"image"`
	Video     NullString `json:"video"`
	Author    string     `json:"author"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

// FetchNewsRequest penggunaan pointer ini agar dapat memberikan value nil jika tidak digunakan
type FetchNewsRequest struct {
	Keyword string
	Type    string
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
