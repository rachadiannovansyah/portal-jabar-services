package domain

import (
	"context"
	"time"
)

type Agenda struct {
	ID          int64      `json:"id"`
	Category    Category   `json:"category" validate:"required"`
	Title       NullString `json:"title" validate:"required"`
	Description NullString `json:"description" validate:"required"`
	Date        NullString `json:"date"`
	Address     NullString `json:"address"`
	StartHour   NullString `json:"start_hour,omitempty"`
	EndHour     NullString `json:"end_hour,omitempty"`
	Image       NullString `json:"image"`
	PublishedBy NullString `json:"published_by"`
	Province    Area       `json:"province" validate:"required"`
	City        Area       `json:"city" validate:"required"`
	District    Area       `json:"district" validate:"required"`
	Village     Area       `json:"village" validate:"required"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   time.Time  `json:"deleted_at"`
}

type ListAgenda struct {
	ID          int64      `json:"id"`
	Title       NullString `json:"title"`
	Description NullString `json:"description"`
	Category    Category   `json:"category"`
	Date        NullString `json:"date"`
	Image       NullString `json:"image"`
	PublishedBy NullString `json:"published_by"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type AgendaUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]Agenda, int64, error)
	GetByID(ctx context.Context, id int64) (Agenda, error)
}

type AgendaRepository interface {
	Fetch(ctx context.Context, params *Request) (new []Agenda, total int64, err error)
	GetByID(ctx context.Context, id int64) (Agenda, error)
}
