package domain

import (
	"context"
	"time"
)

type ServicePublic struct {
	ID                 int64              `json:"id"`
	GeneralInformation GeneralInformation `json:"general_information"`
	Purpose            NullString         `json:"purpose"`
	Facility           NullString         `json:"facility"`
	Requirement        NullString         `json:"requirement"`
	Procedure          NullString         `json:"procedure"`
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
	Procedure          Procedure             `json:"procedure"`
	InfoGraphic        InfoGraphic           `json:"info_graphic"`
	FAQ                FAQ                   `json:"faq"`
	CreatedAt          time.Time             `json:"created_at"`
	UpdatedAt          time.Time             `json:"updated_at"`
}

type GeneralInformationRes struct {
	ID               int64       `json:"id"`
	Name             string      `json:"name"`
	Description      string      `json:"description"`
	Slug             string      `json:"slug"`
	Category         string      `json:"category"`
	Address          string      `json:"address"`
	Unit             string      `json:"unit"`
	Logo             string      `json:"logo"`
	Type             string      `json:"type"`
	Phone            []string    `json:"phone"`
	OperationalHours []string    `json:"operational_hours"`
	Media            Media       `json:"media"`
	SocialMedia      SocialMedia `json:"social_media"`
}

type Media struct {
	Video  string   `json:"video"`
	Images []string `json:"images"`
}

type Purpose struct {
	Title string  `json:"title"`
	Items []Items `json:"items"`
}

type FacilityService struct {
	Title string  `json:"title"`
	Items []Items `json:"items"`
}

type Requirement struct {
	Title string  `json:"title"`
	Items []Items `json:"items"`
}

type Procedure struct {
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

type ServicePublicUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]ServicePublic, error)
	MetaFetch(ctx context.Context, params *Request) (int64, string, error)
	GetBySlug(ctx context.Context, slug string) (ServicePublic, error)
}

type ServicePublicRepository interface {
	Fetch(ctx context.Context, params *Request) (sp []ServicePublic, err error)
	MetaFetch(ctx context.Context, params *Request) (int64, string, error)
	GetBySlug(ctx context.Context, slug string) (ServicePublic, error)
}
