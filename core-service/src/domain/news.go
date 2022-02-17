package domain

import (
	"context"
	"time"
)

// News ...
type News struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title" validate:"required"`
	Excerpt   string     `json:"excerpt" validate:"required"`
	Content   string     `json:"content" validate:"required"`
	Slug      string     `json:"slug"`
	Image     NullString `json:"image"`
	Video     NullString `json:"video"`
	Source    NullString `json:"source"`
	Status    string     `json:"status,omitempty"`
	Views     int64      `json:"views"`
	Shared    int64      `json:"shared"`
	Highlight int8       `json:"highlight,omitempty"`
	Type      string     `json:"type"`
	Tags      []DataTag  `json:"tags"`
	Category  string     `json:"category" validate:"required"`
	Author    User       `json:"author" validate:"required"`
	Area      Area       `json:"area" validate:"required"`
	StartDate time.Time  `json:"start_date"`
	EndDate   time.Time  `json:"end_date"`
	IsLive    int8       `json:"is_live"`
	CreatedBy User       `json:"created_by"`
	UpdatedBy User       `json:"updated_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type StoreNewsRequest struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Excerpt   string    `json:"excerpt"`
	Content   string    `json:"content"`
	Slug      string    `json:"slug"`
	Image     string    `json:"image"`
	Video     string    `json:"video"`
	Source    string    `json:"source"`
	Status    string    `json:"status"`
	Type      string    `json:"type"`
	Category  string    `json:"category"`
	Author    User      `json:"author"`
	StartDate string    `json:"start_date"`
	EndDate   string    `json:"end_date"`
	Tags      []string  `json:"tags"`
	AreaID    int64     `json:"area_id"`
	IsLive    int8      `json:"is_live"`
	CreatedBy User      `json:"created_by"`
	UpdatedBy User      `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UpdateNewsStatusRequest ...
type UpdateNewsStatusRequest struct {
	Status string `json:"status" validate:"required,alpha,eq=DRAFT|eq=REVIEW|eq=PUBLISHED|eq=ARCHIVED"`
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
	Tags      []DataTag  `json:"tags"`
	Area      Area       `json:"area"`
	Status    string     `json:"status"`
	IsLive    int8       `json:"is_live"`
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
	Title     string     `json:"title"`
	Excerpt   string     `json:"excerpt"`
	Content   string     `json:"content"`
	Slug      string     `json:"slug"`
	Image     NullString `json:"image"`
	Video     NullString `json:"video"`
	Source    NullString `json:"source"`
	Status    string     `json:"status"`
	Views     int64      `json:"views"`
	Shared    int64      `json:"shared"`
	Highlight int8       `json:"highlight,omitempty"`
	Type      string     `json:"type"`
	Tags      []DataTag  `json:"tags"`
	Category  string     `json:"category"`
	Author    Author     `json:"author"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type TabStatusResponse struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}

// NewsUsecase represent the news usecases
type NewsUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]News, int64, error)
	FetchNewsBanner(ctx context.Context) ([]NewsBanner, error)
	FetchNewsHeadline(ctx context.Context) ([]News, error)
	GetByID(ctx context.Context, id int64) (News, error)
	GetBySlug(ctx context.Context, slug string) (News, error)
	AddShare(ctx context.Context, id int64) error
	Store(context.Context, *StoreNewsRequest) error
	Update(context.Context, int64, *StoreNewsRequest) error
	UpdateStatus(context.Context, int64, string) error
	TabStatus(context.Context) ([]TabStatusResponse, error)
}

// NewsRepository represent the news repository contract
type NewsRepository interface {
	Fetch(ctx context.Context, params *Request) (new []News, total int64, err error)
	FetchNewsBanner(ctx context.Context) (news []News, err error)
	FetchNewsHeadline(ctx context.Context) (news []News, err error)
	GetByID(ctx context.Context, id int64) (News, error)
	GetBySlug(ctx context.Context, slug string) (News, error)
	AddView(ctx context.Context, id int64) (err error)
	AddShare(ctx context.Context, id int64) (err error)
	Store(ctx context.Context, n *StoreNewsRequest) error
	Update(ctx context.Context, id int64, n *StoreNewsRequest) error
	UpdateStatus(ctx context.Context, id int64, stat string) error
	TabStatus(ctx context.Context) (res []TabStatusResponse, err error)
}
