package domain

import "context"

// Category ...
type Category struct {
	ID          int64      `json:"id"`
	Title       NullString `json:"title" validate:"required"`
	Description NullString `json:"description,omitempty"`
	Type        NullString `json:"type" validate:"required"`
}

// CategoryUsecase ...
type CategoryUsecase interface {
	GetByID(ctx context.Context, id int64) (Category, error)
}

// CategoryRepository ...
type CategoryRepository interface {
	GetByID(ctx context.Context, id int64) (Category, error)
}
