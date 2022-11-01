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
	ID                 int64                    `json:"id"`
	GeneralInformation DetailGeneralInformation `json:"general_information"`
	Purpose            Purpose                  `json:"purpose"`
	Facility           FacilityService          `json:"facility"`
	Requirement        Requirement              `json:"requirement"`
	ToS                TermsOfService           `json:"terms_of_service"`
	InfoGraphic        InfoGraphic              `json:"infographic"`
	FAQ                FAQ                      `json:"faq"`
	CreatedAt          time.Time                `json:"created_at"`
	UpdatedAt          time.Time                `json:"updated_at"`
}

type DetailGeneralInformation struct {
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
	Title string   `json:"title,omitempty"`
	Items []string `json:"items,omitempty"`
}

type FacilityService struct {
	Title string          `json:"title,omitempty"`
	Items []ItemsFacility `json:"items,omitempty"`
}

type ItemsFacility struct {
	Title string `json:"title,omitempty"`
	Image string `json:"image,omitempty"`
}

type Requirement struct {
	Title string  `json:"title,omitempty"`
	Items []Items `json:"items,omitempty"`
}

type TermsOfService struct {
	Title string  `json:"title,omitempty"`
	Items []Items `json:"items,omitempty"`
	Image string  `json:"image,omitempty"`
}

type Items struct {
	Description string `json:"description,omitempty"`
	Image       string `json:"link,omitempty"`
}

type InfoGraphic struct {
	Images []string `json:"images,omitempty"`
}

type FAQ struct {
	Items []QuestionAnswer `json:"items,omitempty"`
}

type QuestionAnswer struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type StorePublicService struct {
	GeneralInformation struct {
		ID               int64    `json:"id"`
		Name             string   `json:"name" validate:"required"`
		Alias            string   `json:"alias" validate:"required"`
		Description      string   `json:"description"`
		Category         string   `json:"category" validate:"required"`
		Addresses        []string `json:"addresses"`
		Unit             string   `json:"unit" validate:"required"`
		Phone            []string `json:"phone"`
		Email            string   `json:"email"`
		Logo             string   `json:"logo"`
		OperationalHours []struct {
			Start string `json:"start"`
			End   string `json:"end"`
		} `json:"operational_hours"`
		Media struct {
			Video  string   `json:"video"`
			Images []string `json:"images"`
		} `json:"media"`
		SocialMedia struct {
			Facebook  string `json:"facebook" validate:"omitempty,url"`
			Instagram string `json:"instagram" validate:"omitempty,url"`
			Twitter   string `json:"twitter" validate:"omitempty,url"`
			Tiktok    string `json:"tiktok" validate:"omitempty,url"`
			Youtube   string `json:"youtube" validate:"omitempty,url"`
		} `json:"social_media"`
		Link struct {
			Website     string `json:"website" validate:"omitempty,url"`
			Google_play string `json:"google_play" validate:"omitempty,url"`
			Google_form string `json:"google_form" validate:"omitempty,url"`
			App_store   string `json:"app_store" validate:"omitempty,url"`
		} `json:"link"`
		Type string `json:"type"`
	} `json:"general_information"`
	Purpose struct {
		Title string   `json:"title"`
		Items []string `json:"items"`
	} `json:"purpose"`
	Facility struct {
		Title string `json:"title"`
		Items []struct {
			Image string `json:"image"`
			Title string `json:"title"`
		} `json:"items"`
	} `json:"facility"`
	Requirement struct {
		Title string `json:"title"`
		Items []struct {
			Link        string `json:"link" validate:"omitemty,url"`
			Description string `json:"description"`
		} `json:"items"`
	} `json:"requirement"`
	Tos struct {
		Title string `json:"title"`
		Items []struct {
			Link        string `json:"link" validate:"omitemty,url"`
			Description string `json:"description"`
		} `json:"items"`
		Image string `json:"image" validate:"omitempty,url"`
	} `json:"tos"`
	Infographic struct {
		Images []string `json:"images" validate:"omitempty,min=1"`
	} `json:"infographic"`
	Faq struct {
		Items []struct {
			Question string `json:"question"`
			Answer   string `json:"answer"`
		} `json:"items"`
	} `json:"faq"`
}

type ServicePublicRepository interface {
	Fetch(ctx context.Context, params *Request) (sp []ServicePublic, err error)
	MetaFetch(ctx context.Context, params *Request) (int64, string, int64, error)
	GetBySlug(ctx context.Context, slug string) (ServicePublic, error)
	Store(context.Context, StorePublicService, *sql.Tx) (err error)
}

type ServicePublicUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]ServicePublic, error)
	MetaFetch(ctx context.Context, params *Request) (int64, string, int64, error)
	GetBySlug(ctx context.Context, slug string) (ServicePublic, error)
	Store(context.Context, StorePublicService) error
}
