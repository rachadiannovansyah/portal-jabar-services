package domain

import (
	"context"
	"time"
)

// News ...
type News struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title" validate:"required"`
	Excerpt   string     `json:"excerpt"`
	Content   string     `json:"content" validate:"required"`
	Slug      NullString `json:"slug"`
	Image     NullString `json:"image"`
	Video     NullString `json:"video"`
	Source    NullString `json:"source"`
	Status    string     `json:"status,omitempty"`
	Views     int64      `json:"views"`
	Highlight int8       `json:"highlight,omitempty"`
	Type      string     `json:"type"`
	Tags      string     `json:"tags"`
	Category  string     `json:"category" validate:"required"`
	Author    User       `json:"author" validate:"required"`
	CreatedBy User       `json:"created_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// NewsListResponse ...
type NewsListResponse struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Excerpt   string     `json:"excerpt"`
	Slug      NullString `json:"slug"`
	Image     NullString `json:"image"`
	Category  string     `json:"category"`
	Author    Author     `json:"author"`
	Video     NullString `json:"video"`
	Source    NullString `json:"source"`
	Tags      string     `json:"tags"`
	CreatedBy NullString `json:"created_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// NewsBanner ...
type NewsBanner struct {
	ID          int64        `json:"id"`
	Title       string       `json:"title"`
	Category    string       `json:"category"`
	Image       NullString   `json:"image"`
	Slug        NullString   `json:"slug"`
	Author      Author       `json:"author,omitempty"`
	CreatedAt   time.Time    `json:"created_at"`
	RelatedNews []NewsBanner `json:"related_news,omitempty"`
}

// DetailNewsResponse ...
type DetailNewsResponse struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title" validate:"required"`
	Excerpt   string     `json:"excerpt"`
	Content   string     `json:"content" validate:"required"`
	Slug      NullString `json:"slug"`
	Image     NullString `json:"image"`
	Video     NullString `json:"video"`
	Source    NullString `json:"source"`
	Status    string     `json:"status,omitempty"`
	Views     int64      `json:"views"`
	Highlight int8       `json:"highlight,omitempty"`
	Type      string     `json:"type"`
	Tags      string     `json:"tags"`
	Category  string     `json:"category" validate:"required"`
	Author    Author     `json:"author" validate:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// NewsUsecase represent the news usecases
type NewsUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]News, int64, error)
	FetchNewsBanner(ctx context.Context) ([]NewsBanner, error)
	FetchNewsHeadline(ctx context.Context) ([]News, error)
	GetByID(ctx context.Context, id int64) (News, error)
}

// NewsRepository represent the news repository contract
type NewsRepository interface {
	Fetch(ctx context.Context, params *Request) (new []News, total int64, err error)
	FetchNewsBanner(ctx context.Context) (news []News, err error)
	FetchNewsHeadline(ctx context.Context) (news []News, err error)
	GetByID(ctx context.Context, id int64) (News, error)
	AddView(ctx context.Context, id int64) (err error)
}
