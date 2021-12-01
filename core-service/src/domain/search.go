package domain

import (
	"context"
)

// Search ...
type Search struct {
	ID        int    `json:"id"`
	Domain    string `json:"domain"`
	Title     string `json:"title"`
	Excerpt   string `json:"excerpt"`
	Content   string `json:"content"`
	Slug      string `json:"slug"`
	Category  string `json:"category"`
	Thumbnail string `json:"thumbnail"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// SearchListResponse ...
type SearchListResponse struct {
	ID        int    `json:"id"`
	Domain    string `json:"domain"`
	Title     string `json:"title"`
	Excerpt   string `json:"excerpt"`
	Slug      string `json:"slug"`
	Category  string `json:"category"`
	Thumbnail string `json:"thumbnail"`
	CreatedAt string `json:"created_at" mapstructure:"created_at"`
}

// SuggestResponse ..
type SuggestResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// SearchUsecase represent the search usecases
type SearchUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]SearchListResponse, int64, interface{}, error)
	SearchSuggestion(ctx context.Context, params *Request) ([]SuggestResponse, error)
}

// SearchRepository represent the search repository contract
type SearchRepository interface {
	Fetch(ctx context.Context, params *Request) (docs []SearchListResponse, total int64, aggs interface{}, err error)
	SearchSuggestion(ctx context.Context, params *Request) (res []SuggestResponse, err error)
}
