package domain

import (
	"context"
	"time"
)

type Informations struct {
	ID                    int64      `json:"id"`
	InformationCategoryID int64      `json:"informationCategoryID" validate:"required"`
	Title                 string     `json:"title" validate:"required"`
	Description           string     `json:"description" validate:"required"`
	Slug                  string     `json:"slug"`
	Image                 NullString `json:"image"`
	ShowDate              string     `json:"showDate"`
	EndDate               string     `json:"endDate"`
	Status                string     `json:"status"`
	CreatedBy             string     `json:"createdBy"`
	UpdatedBy             string     `json:"updatedBy"`
	CreatedAt             time.Time  `json:"createdAt"`
	UpdatedAt             time.Time  `json:"updatedAt"`
	DeletedAt             time.Time  `json:"deletedAt"`
}

type ListInformations struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Excerpt   string     `json:"excerpt"`
	Category  int64      `json:"category"`
	Slug      string     `json:"slug"`
	Image     NullString `json:"image"`
	Author    string     `json:"author"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
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
	FetchOne(ctx context.Context, id int64) (Informations, error)
}

type InformationsRepo interface {
	FetchAll(ctx context.Context, params *FetchInformationsRequest) (new []Informations, total int64, err error)
	FetchOne(ctx context.Context, id int64) (Informations, error)
}
