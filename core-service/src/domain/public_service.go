package domain

import (
	"context"
	"database/sql"
	"time"
)

// PublicService ...
type PublicService struct {
	ID          int64            `json:"id"`
	Name        string           `json:"name"`
	Facilities  string           `json:"facilities"`
	Description NullString       `json:"description"`
	Excerpt     NullString       `json:"excerpt"`
	Unit        NullString       `json:"unit"`
	Url         NullString       `json:"url"`
	Category    NullString       `json:"category"`
	IsActive    NullString       `json:"is_active"`
	Slug        NullString       `json:"slug"`
	ServiceType NullString       `json:"service_type"`
	Video       NullString       `json:"video"`
	Logo        NullString       `json:"logo"`
	Website     NullString       `json:"website"`
	SocialMedia SocialMedia      `json:"social_media"`
	Images      JSONStringSlices `json:"images"`
	Purposes    JSONStringSlices `json:"purposes"`
	Info        PosterInfo       `json:"info"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

type ServicePublic struct {
	ID                 int64            `json:"id"`
	GeneralInformation int64            `json:"general_information_id"`
	Purpose            JSONStringSlices `json:"purpose"`
	Facility           JSONStringSlices `json:"facility"`
	Requirement        JSONStringSlices `json:"requirement"`
	ToS                JSONStringSlices `json:"tos"`
	Info_graphic       JSONStringSlices `json:"info_graphic"`
	Faq                JSONStringSlices `json:"faq"`
	CreatedAt          time.Time        `json:"created_at"`
	UpdatedAt          time.Time        `json:"updated_at"`
}

type Facility struct {
	Image       string `json:"image"`
	Description string `json:"description"`
}

// PosterInfo ...
type PosterInfo struct {
	Requirements JSONStringSlices `json:"requirements"`
	Posters      JSONStringSlices `json:"posters"`
}

type ListPublicServiceResponse struct {
	ID      int64      `json:"id"`
	Name    string     `json:"name"`
	Logo    NullString `json:"logo"`
	Excerpt NullString `json:"excerpt"`
	Slug    NullString `json:"slug"`
}

// PublicService ...
type DetailPublicServiceResponse struct {
	ID          int64            `json:"id"`
	Name        string           `json:"name"`
	Description NullString       `json:"description"`
	Excerpt     NullString       `json:"excerpt"`
	Unit        NullString       `json:"unit"`
	Url         NullString       `json:"url"`
	Category    NullString       `json:"category"`
	IsActive    NullString       `json:"is_active"`
	Slug        NullString       `json:"slug"`
	ServiceType NullString       `json:"service_type"`
	Video       NullString       `json:"video"`
	Logo        NullString       `json:"logo"`
	Website     NullString       `json:"website"`
	SocialMedia SocialMedia      `json:"social_media"`
	Images      JSONStringSlices `json:"images"`
	Purposes    JSONStringSlices `json:"purposes"`
	Facilities  []Facility       `json:"facilities"`
	Info        PosterInfo       `json:"info"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

type StorePserviceRequest struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Unit        string    `json:"unit,omitempty"`
	Url         string    `json:"url"`
	Image       string    `json:"image"`
	Category    string    `json:"category"`
	IsActive    int8      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type StorePublicService struct {
	GeneralInformation GeneralInformation `json:"general_information"`
	Purpose            Purpose            `json:"purpose"`
	Facility           FacilityStore      `json:"facility"`
	Requirement        Requirement        `json:"requirement"`
	Tos                Tos                `json:"tos"`
	Infographic        Infographic        `json:"infographic"`
	Faq                Faq                `json:"faq"`
}
type Media struct {
	Video  string   `json:"video"`
	Images []string `json:"images"`
}

type GeneralInformation struct {
	Name             string      `json:"name"`
	Description      string      `json:"description"`
	Slug             string      `json:"slug"`
	Category         string      `json:"category"`
	Address          string      `json:"address"`
	Unit             string      `json:"unit"`
	Phone            []string    `json:"phone"`
	Logo             string      `json:"logo"`
	OperationalHours []string    `json:"operational_hours"`
	Media            Media       `json:"media"`
	SocialMedia      SocialMedia `json:"social_media"`
	Type             string      `json:"type"`
}
type Purpose struct {
	Title string   `json:"title"`
	Items []string `json:"items"`
}
type Items struct {
	Image       string `json:"image"`
	Description string `json:"description"`
}
type FacilityStore struct {
	Title string          `json:"title"`
	Items []FacilityItems `json:"items"`
}
type FacilityItems struct {
	Link        string `json:"link"`
	Description string `json:"description"`
}
type Requirement struct {
	Title string  `json:"title"`
	Items []Items `json:"items"`
}
type Tos struct {
	Title string     `json:"title"`
	Items []TosItems `json:"items"`
	Image string     `json:"image"`
}
type Infographic struct {
	Images []string `json:"images"`
}
type TosItems struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
type Faq struct {
	Items []Items `json:"items"`
}

// PublicServiceUsecase ...
type PublicServiceUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]PublicService, error)
	MetaFetch(ctx context.Context, params *Request) (int64, string, error)
	GetBySlug(ctx context.Context, slug string) (PublicService, error)
	Store(context.Context, StorePublicService) error
}

// PublicServiceRepository ...
type PublicServiceRepository interface {
	Fetch(ctx context.Context, params *Request) (ps []PublicService, err error)
	MetaFetch(ctx context.Context, params *Request) (int64, string, error)
	GetBySlug(ctx context.Context, slug string) (PublicService, error)
	Store(context.Context, StorePublicService) (err error)
	StoreGeneralInformation(context.Context, *sql.Tx, StorePublicService) (id int64, err error)
}
