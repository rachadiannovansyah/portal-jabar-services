package domain

import (
	"context"
)

type Logo struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}

type LogoUsecase interface {
	Fetch(ctx context.Context, params Request) (res []Logo, total int64, err error)
}

type LogoRepository interface {
	Fetch(ctx context.Context, params Request) (res []Logo, total int64, err error)
}
