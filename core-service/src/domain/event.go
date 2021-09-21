package domain

import (
	"context"
	"time"
)

// Event model ..
type Event struct {
	ID        int64      `json:"id"`
	Category  Category   `json:"category" validate:"required"`
	Title     NullString `json:"title" validate:"required"`
	Priority  NullString `json:"priority"`
	Address   NullString `json:"address"`
	StartHour NullString `json:"start_hour,omitempty"`
	EndHour   NullString `json:"end_hour,omitempty"`
	Weeks     NullString `json:"weeks"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// ListEvent ...
type ListEvent struct {
	ID        int64      `json:"id"`
	Category  Category   `json:"category"`
	Title     NullString `json:"title"`
	Priority  NullString `json:"priority"`
	Address   NullString `json:"address"`
	StartHour NullString `json:"start_hour,omitempty"`
	EndHour   NullString `json:"end_hour,omitempty"`
	Week      int64      `json:"week"`
	Month     int64      `json:"month"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// EventUsecase ..
type EventUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]Event, int64, error)
	GetByID(ctx context.Context, id int64) (Event, error)
}

// EventRepository ..
type EventRepository interface {
	Fetch(ctx context.Context, params *Request) (new []Event, total int64, err error)
	GetByID(ctx context.Context, id int64) (Event, error)
}
