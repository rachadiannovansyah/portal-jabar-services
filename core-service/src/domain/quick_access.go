package domain

import (
	"context"
	"database/sql"
	"time"
)

type QuickAccess struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	IsActive    int8      `json:"is_active"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateStatusQuickAccess struct {
	IsActive *int8 `json:"is_active" validate:"required,eq=1|eq=0"`
}

type StoreQuickAccess struct {
	Title       string      `json:"title" validate:"required,max=80"`
	Description string      `json:"description" validate:"required,max=200"`
	Link        string      `json:"link"`
	Image       ImageBanner `json:"image" validate:"required"`
	IsActive    *int8       `json:"is_active" validate:"required,eq=1|eq=0"`
}

type QuickAccessResponse struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	IsActive    bool      `json:"is_active"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type QuickAccessUsecase interface {
	Store(ctx context.Context, body *StoreQuickAccess) (err error)
	Fetch(ctx context.Context, params Request) (res []QuickAccess, total int64, err error)
	Delete(ctx context.Context, ID int64) (err error)
	GetByID(ctx context.Context, ID int64) (res QuickAccess, err error)
	Update(c context.Context, ID int64, body *StoreQuickAccess) (err error)
	UpdateStatus(ctx context.Context, ID int64, body *UpdateStatusQuickAccess) (err error)
}

type QuickAccessRepository interface {
	GetTx(ctx context.Context) (*sql.Tx, error)
	Fetch(ctx context.Context, params Request) (res []QuickAccess, total int64, err error)
	Store(ctx context.Context, body *StoreQuickAccess, tx *sql.Tx) (err error)
	Update(ctx context.Context, ID int64, body *StoreQuickAccess, tx *sql.Tx) (err error)
	UpdateStatus(ctx context.Context, ID int64, body *UpdateStatusQuickAccess, tx *sql.Tx) (err error)
	Delete(ctx context.Context, ID int64, tx *sql.Tx) (err error)
	GetByID(ctx context.Context, ID int64, tx *sql.Tx) (res QuickAccess, err error)
}
