package domain

import (
	"context"
	"time"
)

// News ...
type News struct {
	ID             int64     `json:"id"`
	NewsCategoryID int64     `json:"newsCategoryId" validate:"required"`
	Title          string    `json:"title" validate:"required"`
	Content        string    `json:"content" validate:"required"`
	Slug           string    `json:"slug"`
	ImagePath      string    `json:"imagePath"`
	VideoUrl       string    `json:"videoUrl"`
	NewsSource     string    `json:"newsSource"`
	ShowDate       string    `json:"showDate"`
	EndDate        string    `json:"endDate"`
	IsPublished    bool      `json:"isPublished"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// NewsUsecase represent the content's usecases
type NewsUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]News, string, error)
	GetByID(ctx context.Context, id int64) (News, error)
	Update(ctx context.Context, ar *News) error
	GetBySlug(ctx context.Context, slug string) (News, error)
	Store(context.Context, *News) error
	Delete(ctx context.Context, id int64) error
}

// NewsRepository represent the content's repository contract
type NewsRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []News, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (News, error)
	GetBySlug(ctx context.Context, slug string) (News, error)
	Update(ctx context.Context, ar *News) error
	Store(ctx context.Context, a *News) error
	Delete(ctx context.Context, id int64) error
}
