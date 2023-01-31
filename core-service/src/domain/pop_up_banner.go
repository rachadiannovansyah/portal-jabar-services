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

type StorePopUpBannerRequest struct {
	ID           int64                `json:"id"`
	Title        string               `json:"title" validate:"required,max=255"`
	CustomButton CustomButtonlabel    `json:"custom_button,omitempty"`
	Scheduler    SchedulerPopUpBanner `json:"scheduler,omitempty"`
	Image        ImageBanner          `json:"image" validate:"required"`
}

type UpdateStatusPopUpBannerRequest struct {
	Status string `json:"status" validate:"required,eq=ACTIVE|eq=NON-ACTIVE"`
}

type CustomButtonlabel struct {
	Label string `json:"label"`
	Link  string `json:"link"`
}

type SchedulerPopUpBanner struct {
	Duration  int64   `json:"duration"`
	StartDate *string `json:"start_date"`
}

type PopUpBannerUsecase interface {
	Fetch(ctx context.Context, auth *JwtCustomClaims, params *Request) (res []PopUpBanner, total int64, err error)
	GetByID(ctx context.Context, id int64) (res PopUpBanner, err error)
	Store(ctx context.Context, auth *JwtCustomClaims, body StorePopUpBannerRequest) (err error)
	Delete(ctx context.Context, id int64) (err error)
	UpdateStatus(ctx context.Context, id int64, status string) (err error)
}

type PopUpBannerRepository interface {
	Fetch(ctx context.Context, params *Request) (res []PopUpBanner, total int64, err error)
	GetByID(ctx context.Context, id int64) (res PopUpBanner, err error)
	Store(ctx context.Context, body StorePopUpBannerRequest) (err error)
	Delete(ctx context.Context, id int64) (err error)
	UpdateStatus(ctx context.Context, id int64, status string) (err error)
}
