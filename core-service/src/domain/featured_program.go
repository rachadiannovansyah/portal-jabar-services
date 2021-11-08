package domain

import (
	"context"
	"time"
)

// FeaturedProgram ...
type FeaturedProgram struct {
	ID           int64            `json:"id"`
	Title        string           `json:"title"`
	Excerpt      string           `json:"excerpt"`
	Description  string           `json:"description"`
	Organization string           `json:"organization"`
	Categories   JSONStringSlices `json:"categories"`
	ServiceType  string           `json:"service_type"`
	Websites     JSONStringSlices `json:"websites"`
	SocialMedia  SocialMedia      `json:"social_media"`
	Logo         NullString       `json:"logo"`
	CreatedAt    *time.Time       `json:"created_at,omitempty"`
	UpdatedAt    *time.Time       `json:"updated_at,omitempty"`
}

// SocialMedia ...
type SocialMedia struct {
	Facebook  NullString `json:"facebook"`
	Instagram NullString `json:"instagram"`
	Twitter   NullString `json:"twitter"`
	Tiktok    NullString `json:"tiktok"`
	Youtube   NullString `json:"youtube"`
}

// FeaturedProgramUsecase represent the featured program usecases
type FeaturedProgramUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]FeaturedProgram, error)
}

// FeaturedProgramRepository represent the featured program repository contract
type FeaturedProgramRepository interface {
	Fetch(ctx context.Context, params *Request) (fp []FeaturedProgram, err error)
}
