package domain

import (
	"context"
	"time"
)

// Event ...
type Event struct {
	ID           int64      `json:"id"`
	Title        NullString `json:"title"`
	Description  NullString `json:"description"`
	Date         NullString `json:"date"`
	Priority     NullString `json:"priority"`
	StartHour    NullString `json:"start_hour,omitempty"`
	EndHour      NullString `json:"end_hour,omitempty"`
	Image        NullString `json:"image"`
	PublishedBy  NullString `json:"published_by"`
	Type         NullString `json:"type"`
	Address      NullString `json:"address"`
	URL          NullString `json:"url"`
	Category     Category   `json:"category"`
	ProvinceCode Area       `json:"province_code"`
	CityCode     Area       `json:"city_code"`
	DistrictCode Area       `json:"district_code"`
	VillageCode  Area       `json:"village_code"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

// ListEventResponse model ..
type ListEventResponse struct {
	ID        int64      `json:"id"`
	Title     NullString `json:"title" validate:"required"`
	Date      NullString `json:"date"`
	StartHour NullString `json:"start_hour,omitempty"`
	EndHour   NullString `json:"end_hour,omitempty"`
	Priority  NullString `json:"priority"`
	Type      NullString `json:"type"`
	Address   NullString `json:"address"`
	URL       NullString `json:"url"`
	Category  Category   `json:"category" validate:"required"`
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
