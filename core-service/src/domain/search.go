package domain

import (
	"context"
	"time"
)

// Search ...
type Search struct {
	ID        int       `json:"id"`
	Domain    string    `json:"domain"`
	Title     string    `json:"title"`
	Excerpt   string    `json:"excerpt"`
	Content   string    `json:"content"`
	Slug      string    `json:"slug"`
	Category  string    `json:"category"`
	Thumbnail string    `json:"thumbnail"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
}

// SearchListResponse ...
type SearchListResponse struct {
	ID        int       `json:"id"`
	Domain    string    `json:"domain"`
	Title     string    `json:"title"`
	Excerpt   string    `json:"excerpt"`
	Slug      string    `json:"slug"`
	Category  string    `json:"category"`
	Thumbnail string    `json:"thumbnail"`
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"`
}

// SuggestResponse ..
type SuggestResponse struct {
	Value string `json:"value" mapstructure:"title"`
}

// SearchUsecase represent the search usecases
type SearchUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]SearchListResponse, int64, interface{}, error)
	Store(ctx context.Context, indices string, body *Search) error
	Update(ctx context.Context, indices string, id int, body *Search) error
	Delete(ctx context.Context, indices string, id int, domain string) error
	SearchSuggestion(ctx context.Context, params *Request) ([]SuggestResponse, error)
}

// SearchRepository represent the search repository contract
type SearchRepository interface {
	Fetch(ctx context.Context, indices string, params *Request) (docs []SearchListResponse, total int64, aggs interface{}, err error)
	Store(ctx context.Context, indices string, body *Search) (err error)
	Update(ctx context.Context, indices string, id int, body *Search) (err error)
	Delete(ctx context.Context, indices string, id int, domain string) (err error)
	SearchSuggestion(ctx context.Context, indices string, params *Request) (res []SuggestResponse, err error)
}
