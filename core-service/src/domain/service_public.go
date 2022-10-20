package domain

import (
	"context"
	"database/sql"
	"time"
)

type ServicePublic struct {
	ID                 int64              `json:"id"`
	GeneralInformation GeneralInformation `json:"general_information"`
	Purpose            NullString         `json:"purpose"`
	Facility           NullString         `json:"facility"`
	Requirement        NullString         `json:"requirement"`
	ToS                NullString         `json:"tos"`
	InfoGraphic        NullString         `json:"info_graphic"`
	FAQ                NullString         `json:"faq"`
	CreatedAt          time.Time          `json:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at"`
}

type ListServicePublicResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
}

type DetailServicePublicResponse struct {
	ID                 int64                 `json:"id"`
	GeneralInformation GeneralInformationRes `json:"general_information"`
	Purpose            Purpose               `json:"purpose"`
	Facility           FacilityService       `json:"facility"`
	Requirement        Requirement           `json:"requirement"`
	ToS                TermsOfService        `json:"terms_of_service"`
	InfoGraphic        InfoGraphic           `json:"infographic"`
	FAQ                FAQ                   `json:"faq"`
	CreatedAt          time.Time             `json:"created_at"`
	UpdatedAt          time.Time             `json:"updated_at"`
}

type GeneralInformationRes struct {
	ID               int64              `json:"id"`
	Name             string             `json:"name"`
	Alias            string             `json:"alias"`
	Description      string             `json:"description"`
	Slug             string             `json:"slug"`
	Category         string             `json:"category"`
	Addresses        []string           `json:"addresses"`
	Unit             string             `json:"unit"`
	Logo             string             `json:"logo"`
	Type             string             `json:"type"`
	Phone            []string           `json:"phone"`
	Email            string             `json:"email"`
	Link             Link               `json:"link"`
	OperationalHours []OperationalHours `json:"operational_hours"`
	Media            Media              `json:"media"`
	SocialMedia      SocialMedia        `json:"social_media"`
}

type OperationalHours struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type Link struct {
	Website    string `json:"website"`
	GooglePlay string `json:"google_play"`
	GoogleForm string `json:"google_form"`
	AppStore   string `json:"app_store"`
}

type Media struct {
	Video  string   `json:"video"`
	Images []string `json:"images"`
}

type Purpose struct {
	Title string   `json:"title"`
	Items []string `json:"items"`
}

type FacilityService struct {
	Title string  `json:"title"`
	Items []Items `json:"items"`
}

type Requirement struct {
	Title string  `json:"title"`
	Items []Items `json:"items"`
}

type TermsOfService struct {
	Title string  `json:"title"`
	Items []Items `json:"items"`
}

type Items struct {
	Description string `json:"description"`
	Image       string `json:"link"`
}

type InfoGraphic struct {
	Images []string `json:"images"`
}

type FAQ struct {
	Items []QuestionAnswer `json:"items"`
}

type QuestionAnswer struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type StorePublicService struct {
	GeneralInformation struct {
		Name             string   `json:"name" validate:"required"`
		Alias            string   `json:"alias" validate:"required"`
		Description      string   `json:"description" validate:"required"`
		Category         string   `json:"category" validate:"required"`
		Addresses        []string `json:"addresses" validate:"required"`
		Unit             string   `json:"unit" validate:"required"`
		Phone            []string `json:"phone" validate:"required,min=1"`
		Email            string   `json:"email" validate:"omitempty,url"`
		Logo             string   `json:"logo" validate:"required"`
		OperationalHours []struct {
			Start string `json:"start" validate:"required,ISO8601date"`
			End   string `json:"end" validate:"required,ISO8601date"`
		} `json:"operational_hours" validate:"required,min=1"`
		Media struct {
			Video  string   `json:"video" validate:"required"`
			Images []string `json:"images" validate:"required,min=1"`
		} `json:"media" validate:"required"`
		SocialMedia struct {
			Facebook  string `json:"facebook" validate:"omitempty,url"`
			Instagram string `json:"instagram" validate:"omitempty,url"`
			Twitter   string `json:"twitter" validate:"omitempty,url"`
			Tiktok    string `json:"tiktok" validate:"omitempty,url"`
			Youtube   string `json:"youtube" validate:"omitempty,url"`
		} `json:"social_media" validate:"required"`
		Link struct {
			Website     string `json:"website" validate:"omitempty,url"`
			Google_play string `json:"google_play" validate:"omitempty,url"`
			Google_form string `json:"google_form" validate:"omitempty,url"`
			App_store   string `json:"app_store" validate:"omitempty,url"`
		} `json:"link" validate:"required"`
		Type string `json:"type" validate:"required"`
	} `json:"general_information"`
	Purpose struct {
		Title string   `json:"title" validate:"required"`
		Items []string `json:"items" validate:"required,min=1"`
	} `json:"purpose" validate:"required"`
	Facility struct {
		Title string `json:"title" validate:"required"`
		Items []struct {
			Link        string `json:"link" validate:"required,url"`
			Description string `json:"description" validate:"required"`
		} `json:"items" validate:"required,min=1"`
	} `json:"facility" validate:"required"`
	Requirement struct {
		Title string `json:"title" validate:"required"`
		Items []struct {
			Link        string `json:"link" validate:"required,url"`
			Description string `json:"description" validate:"required"`
		} `json:"items" validate:"required,min=1"`
	} `json:"requirement" validate:"required"`
	Tos struct {
		Title string `json:"title" validate:"required"`
		Items []struct {
			Link        string `json:"link" validate:"required,url"`
			Description string `json:"description" validate:"required"`
		} `json:"items" validate:"required,min=1"`
		Image string `json:"image" validate:"required,url"`
	} `json:"tos" validate:"required"`
	Infographic struct {
		Images []string `json:"images" validate:"required,min=1"`
	} `json:"infographic" validate:"required"`
	Faq struct {
		Items []struct {
			Question string `json:"question"`
			Answer   string `json:"answer"`
		} `json:"items" validate:"required,min=1"`
	} `json:"faq" validate:"required"`
}

type ServicePublicUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]ServicePublic, error)
	MetaFetch(ctx context.Context, params *Request) (int64, string, error)
	GetBySlug(ctx context.Context, slug string) (ServicePublic, error)
	Store(context.Context, StorePublicService) error
}

type ServicePublicRepository interface {
	Fetch(ctx context.Context, params *Request) (sp []ServicePublic, err error)
	MetaFetch(ctx context.Context, params *Request) (int64, string, error)
	GetBySlug(ctx context.Context, slug string) (ServicePublic, error)
	Store(context.Context, StorePublicService) (err error)
	StoreGeneralInformation(context.Context, *sql.Tx, StorePublicService) (id int64, err error)
	UpdateSlugGeneralInformation(context.Context, *sql.Tx, StorePublicService, int64) (err error)
}
