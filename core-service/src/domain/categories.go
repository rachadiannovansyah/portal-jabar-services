package domain

import "context"

type Category struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title" validate:"required"`
	Description NullString `json:"description,omitempty"`
	Type        string     `json:"type" validate:"required"`
}

type CategoriesUsecase interface {
	GetByID(ctx context.Context, id int64) (Category, error)
}

type CategoriesRepository interface {
	GetByID(ctx context.Context, id int64) (Category, error)
}
