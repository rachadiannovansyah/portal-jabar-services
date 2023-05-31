package domain

import (
	"context"
	"database/sql"
	"time"
)

type InfographicBanner struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Sequence  int8      `json:"sequence"`
	Link      string    `json:"link"`
	IsActive  bool      `json:"is_active"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SyncSequence struct {
	ID       int64 `json:"id"`
	Sequence int8  `json:"sequence"`
}
type StoreInfographicBanner struct {
	Title string      `json:"title" validate:"required,max=255"`
	Link  string      `json:"link"`
	Image ImageBanner `json:"image" validate:"required"`
}

type InfographicBannerResponse struct {
	ID        int64       `json:"id"`
	Title     string      `json:"title"`
	Sequence  int8        `json:"sequence"`
	Link      string      `json:"link"`
	IsActive  bool        `json:"is_active"`
	Image     ImageBanner `json:"image"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type InfographicBannerUsecase interface {
	Store(ctx context.Context, body *StoreInfographicBanner) (err error)
	Fetch(ctx context.Context, params Request) (res []InfographicBanner, total int64, err error)
	Delete(ctx context.Context, ID int64) (err error)
	GetByID(ctx context.Context, ID int64) (res InfographicBanner, err error)
}

type InfographicBannerRepository interface {
	GetTx(ctx context.Context) (*sql.Tx, error)
	Fetch(ctx context.Context, params Request) (res []InfographicBanner, total int64, err error)
	Store(ctx context.Context, body *StoreInfographicBanner, tx *sql.Tx) (err error)
	GetLastSequence(ctx context.Context) (count int64)
	SyncSequence(ctx context.Context, sequence int64, tx *sql.Tx) (err error)
	Delete(ctx context.Context, ID int64, tx *sql.Tx) (err error)
	GetByID(ctx context.Context, ID int64, tx *sql.Tx) (res InfographicBanner, err error)
}
