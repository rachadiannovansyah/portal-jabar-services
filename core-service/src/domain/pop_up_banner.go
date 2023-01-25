package domain

import (
	"context"
	"time"
)

type PopUpBanner struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	ButtonLabel string     `json:"button_label,omitempty"`
	Image       NullString `json:"image,omitempty"`
	Link        string     `json:"link"`
	Status      string     `json:"status"`
	Duration    int64      `json:"duration,omitempty"`
	StartDate   *time.Time `json:"start_date,omitempty"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type ListPopUpBannerResponse struct {
	ID        int64       `json:"id"`
	Title     string      `json:"title"`
	Image     ImageBanner `json:"image,omitempty"`
	Link      string      `json:"link"`
	Duration  int64       `json:"duration,omitempty"`
	StartDate *time.Time  `json:"start_date,omitempty"`
	Status    string      `json:"status"`
}

type DetailPopUpBannerResponse struct {
	ID          int64       `json:"id"`
	Title       string      `json:"title"`
	ButtonLabel string      `json:"button_label"`
	Image       ImageBanner `json:"image,omitempty"`
	Link        string      `json:"link"`
	Status      string      `json:"status"`
	Duration    int64       `json:"duration"`
	StartDate   *time.Time  `json:"start_date"`
	UpdateAt    time.Time   `json:"updated_at"`
}

type ImageBanner struct {
	Desktop string `json:"desktop,omitempty"`
	Mobile  string `json:"mobile,omitempty"`
}

type PopUpBannerUsecase interface {
	Fetch(ctx context.Context, auth *JwtCustomClaims, params *Request) (res []PopUpBanner, total int64, err error)
	GetByID(ctx context.Context, id int64) (res PopUpBanner, err error)
}

type PopUpBannerRepository interface {
	Fetch(ctx context.Context, params *Request) (res []PopUpBanner, total int64, err error)
	GetByID(ctx context.Context, id int64) (res PopUpBanner, err error)
}
