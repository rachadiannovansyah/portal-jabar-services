package domain

import (
	"context"
	"time"
)

// Information ...
type Information struct {
	ID        int64      `json:"id"`
	Category  Category   `json:"category" validate:"required"`
	Title     NullString `json:"title" validate:"required"`
	Content   NullString `json:"content" validate:"required"`
	Slug      NullString `json:"slug"`
	Image     NullString `json:"image"`
	ShowDate  NullString `json:"show_date,omitempty"`
	EndDate   NullString `json:"end_date,omitempty"`
	Status    NullString `json:"status"`
	CreatedBy NullString `json:"created_by"`
	UpdatedBy NullString `json:"updated_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt time.Time  `json:"deleted_at"`
}

// ListInformation ...
type ListInformation struct {
	ID        int64      `json:"id"`
	Title     NullString `json:"title"`
	Excerpt   NullString `json:"excerpt"`
	Category  Category   `json:"category"`
	Slug      NullString `json:"slug"`
	Image     NullString `json:"image"`
	Author    NullString `json:"author"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// InformationUsecase ...
type InformationUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]Information, int64, error)
	GetByID(ctx context.Context, id int64) (Information, error)
}

// InformationRepository ...
type InformationRepository interface {
	Fetch(ctx context.Context, params *Request) (new []Information, total int64, err error)
	GetByID(ctx context.Context, id int64) (Information, error)
}
