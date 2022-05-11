package domain

import (
	"context"
	"time"
)

// Award ...
type Award struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Logo        NullString `json:"logo"`
	Appreciator string     `json:"appreciator"`
	Description string     `json:"description"`
	Category    string     `json:"category"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// Categories returns all categories
type AwardCategoryAggregation struct {
	Category string `json:"category"`
	Count    int    `json:"count"`
}

// AwardUsecase represent the award usecases
type AwardUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]Award, int64, error)
	GetByID(ctx context.Context, id int64) (Award, error)
	FetchCategories(ctx context.Context) ([]AwardCategoryAggregation, error)
}

// AwardRepository represent the award repository contract
type AwardRepository interface {
	Fetch(ctx context.Context, params *Request) (res []Award, total int64, err error)
	GetByID(ctx context.Context, id int64) (res Award, err error)
	FetchCategories(ctx context.Context) (res []AwardCategoryAggregation, err error)
}
