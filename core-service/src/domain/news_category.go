package domain

import "context"

// NewsCategory ...
type NewsCategory struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title" validate:"required"`
	Description NullString `json:"description,omitempty"`
}

// NewsCategoryUsecase represent the news category usecases
type NewsCategoryUsecase interface {
	GetByID(ctx context.Context, id int64) (NewsCategory, error)
}

// NewsCategoryRepository represent the news repository contract
type NewsCategoryRepository interface {
	GetByID(ctx context.Context, id int64) (NewsCategory, error)
}
