package domain

import (
	"context"
	"time"
)

// PublicService ...
type PublicService struct {
	ID          int64            `json:"id"`
	Name        string           `json:"name"`
	Description NullString       `json:"description"`
	Excerpt     NullString       `json:"excerpt"`
	Unit        NullString       `json:"unit"`
	Url         NullString       `json:"url"`
	Category    NullString       `json:"category"`
	IsActive    NullString       `json:"is_active"`
	Slug        NullString       `json:"slug"`
	ServiceType NullString       `json:"service_type"`
	Video       NullString       `json:"video"`
	Logo        NullString       `json:"logo"`
	Website     NullString       `json:"website"`
	SocialMedia SocialMedia      `json:"social_media"`
	Images      JSONStringSlices `json:"images"`
	Purposes    JSONStringSlices `json:"purposes"`
	Facilities  JSONStringSlices `json:"facilities"`
	Info        JSONStringSlices `json:"info"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

type ListPublicServiceResponse struct {
	ID      int64      `json:"id"`
	Name    string     `json:"name"`
	Logo    NullString `json:"logo"`
	Excerpt NullString `json:"excerpt"`
	Slug    NullString `json:"slug"`
}

type StorePserviceRequest struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Unit        string    `json:"unit,omitempty"`
	Url         string    `json:"url"`
	Image       string    `json:"image"`
	Category    string    `json:"category"`
	IsActive    int8      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// PublicServiceUsecase ...
type PublicServiceUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]PublicService, error)
	MetaFetch(ctx context.Context, params *Request) (int64, string, error)
	GetBySlug(ctx context.Context, slug string) (PublicService, error)
}

// PublicServiceRepository ...
type PublicServiceRepository interface {
	Fetch(ctx context.Context, params *Request) (ps []PublicService, err error)
	MetaFetch(ctx context.Context, params *Request) (int64, string, error)
	GetBySlug(ctx context.Context, slug string) (PublicService, error)
}
