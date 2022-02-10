package domain

import (
	"context"
	"time"
)

// Event ...
type Event struct {
	ID           int64      `json:"id"`
	Title        string     `json:"title"`
	Priority     string     `json:"priority"`
	Date         time.Time  `json:"date"`
	StartHour    string     `json:"start_hour"`
	EndHour      string     `json:"end_hour"`
	Image        NullString `json:"image"`
	Type         string     `json:"type"`
	Status       string     `json:"status"`
	Address      NullString `json:"address"`
	URL          NullString `json:"url"`
	Category     string     `json:"category"`
	Tags         []DataTag  `json:"tags"`
	ProvinceCode Area       `json:"province_code"`
	CityCode     Area       `json:"city_code"`
	DistrictCode Area       `json:"district_code"`
	VillageCode  Area       `json:"village_code"`
	CreatedBy    User       `json:"created_by"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

// StoreRequestEvent ..
type StoreRequestEvent struct {
	ID        int64    `json:"id"`
	Title     string   `json:"title" validate:"required"`
	Type      string   `json:"type" validate:"required"`
	URL       string   `json:"url"`
	Address   string   `json:"address"`
	Date      string   `json:"date" validate:"required"`
	StartHour string   `json:"start_hour" validate:"required"`
	EndHour   string   `json:"end_hour" validate:"required"`
	Category  string   `json:"category" validate:"required"`
	Tags      []string `json:"tags"`
}

// UpdateRequestEvent ..
type UpdateRequestEvent struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Type      string    `json:"type" validate:"required"`
	URL       string    `json:"url"`
	Address   string    `json:"address"`
	Date      string    `json:"date" validate:"required"`
	StartHour string    `json:"start_hour" validate:"required"`
	EndHour   string    `json:"end_hour" validate:"required"`
	Category  string    `json:"category" validate:"required"`
	Tags      []string  `json:"tags"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ListEventResponse model ..
type ListEventResponse struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title" validate:"required"`
	Date      time.Time  `json:"date"`
	StartHour string     `json:"start_hour"`
	EndHour   string     `json:"end_hour"`
	Priority  string     `json:"priority"`
	Type      string     `json:"type"`
	Status    string     `json:"status"`
	Address   NullString `json:"address"`
	URL       NullString `json:"url"`
	Category  string     `json:"category" validate:"required"`
	Tags      []DataTag  `json:"tags"`
}

type DetailEventResponse struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Date      time.Time  `json:"date"`
	StartHour string     `json:"start_hour"`
	EndHour   string     `json:"end_hour"`
	Type      string     `json:"type"`
	Status    string     `json:"status"`
	Address   NullString `json:"address"`
	URL       NullString `json:"url"`
	Category  string     `json:"category"`
	Tags      []DataTag  `json:"tags"`
	CreatedBy Author     `json:"created_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

//ListEventCalendarReponse ..
type ListEventCalendarReponse struct {
	ID    int64     `json:"id"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}

// EventUsecase ..
type EventUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]Event, int64, error)
	Store(context.Context, *StoreRequestEvent) error
	GetByID(ctx context.Context, id int64) (Event, error)
	GetByTitle(ctx context.Context, title string) (Event, error)
	Delete(ctx context.Context, id int64) error
	Update(context.Context, int64, *UpdateRequestEvent) error
	ListCalendar(ctx context.Context, params *Request) ([]Event, error)
	AgendaPortal(ctx context.Context, params *Request) ([]Event, int64, error)
}

// EventRepository ..
type EventRepository interface {
	Fetch(ctx context.Context, params *Request) (new []Event, total int64, err error)
	Store(ctx context.Context, body *StoreRequestEvent) error
	GetByID(ctx context.Context, id int64) (Event, error)
	GetByTitle(ctx context.Context, title string) (Event, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, body *UpdateRequestEvent) error
	ListCalendar(ctx context.Context, params *Request) ([]Event, error)
	AgendaPortal(ctx context.Context, params *Request) (new []Event, total int64, err error)
}
