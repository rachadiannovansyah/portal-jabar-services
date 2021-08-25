package domain

import (
	"context"
	"time"
)

type Informations struct {
	ID        int64      `json:"id"`
	Category  Category   `json:"category_id" validate:"required"`
	Title     string     `json:"title" validate:"required"`
	Content   string     `json:"content" validate:"required"`
	Slug      string     `json:"slug"`
	Image     NullString `json:"image"`
	ShowDate  string     `json:"show_date,omitempty"`
	EndDate   string     `json:"end_date,omitempty"`
	Status    string     `json:"status"`
	CreatedBy string     `json:"created_by"`
	UpdatedBy string     `json:"updated_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt time.Time  `json:"deleted_at"`
}

type ListInformations struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Excerpt   string     `json:"excerpt"`
	Category  Category   `json:"category_id"`
	Slug      string     `json:"slug"`
	Image     NullString `json:"image"`
	Author    string     `json:"author"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type FetchInformationsRequest struct {
	Keyword string
	Type    string
	PerPage int64
	Offset  int64
	OrderBy string
	SortBy  string
}

type InformationsUcase interface {
	FetchAll(ctx context.Context, params *FetchInformationsRequest) ([]Informations, int64, error)
	GetByID(ctx context.Context, id int64) (Informations, error)
}

type InformationsRepo interface {
	FetchAll(ctx context.Context, params *FetchInformationsRequest) (new []Informations, total int64, err error)
	GetByID(ctx context.Context, id int64) (Informations, error)
}
