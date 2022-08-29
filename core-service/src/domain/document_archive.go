package domain

import (
	"context"
	"time"
)

// DocumentArchive Struct ...
type DocumentArchive struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Excerpt     string    `json:"excerpt" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Source      string    `json:"source"`
	Mimetype    string    `json:"mimetype"`
	Category    string    `json:"category" validate:"required"`
	CreatedBy   User      `json:"created_by"`
	UpdatedBy   User      `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ListDocumentArchive ...
type ListDocumentArchive struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Excerpt   string    `json:"excerpt" validate:"required"`
	Source    string    `json:"source"`
	Mimetype  string    `json:"mimetype"`
	Category  string    `json:"category" validate:"required"`
	CreatedBy Author    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DocumentArchiveUsecase ...
type DocumentArchiveUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]DocumentArchive, int64, error)
}

// DocumentArchiveRepository ...
type DocumentArchiveRepository interface {
	Fetch(ctx context.Context, params *Request) ([]DocumentArchive, int64, error)
}
