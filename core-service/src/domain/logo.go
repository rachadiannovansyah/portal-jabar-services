package domain

import (
	"context"
)

type Logo struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}

type StoreLogoRequest struct {
	Title string `json:"title" validate:"required,max=100"`
	Image string `json:"image" validate:"required"`
}

type LogoUsecase interface {
	Store(ctx context.Context, body *StoreLogoRequest) (err error)
	Fetch(ctx context.Context, params Request) (res []Logo, total int64, err error)
}

type LogoRepository interface {
	Fetch(ctx context.Context, params Request) (res []Logo, total int64, err error)
	Store(ctx context.Context, body *StoreLogoRequest) (err error)
}
