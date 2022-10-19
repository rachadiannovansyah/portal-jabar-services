package domain

import "context"

type GeneralInformation struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	Alias            string `json:"alias"`
	Description      string `json:"description"`
	Slug             string `json:"slug"`
	Category         string `json:"category"`
	Addresses        string `json:"addresses"`
	Unit             string `json:"unit"`
	Phone            string `json:"phone"`
	Email            string `json:"email"`
	Logo             string `json:"logo"`
	OperationalHours string `json:"operationalHours"`
	Link             string `json:"link"`
	Media            string `json:"media"`
	SocialMedia      string `json:"socialMedia"`
	Type             string `json:"type"`
}

type GeneralInformationRepository interface {
	GetByID(ctx context.Context, id int64) (GeneralInformation, error)
}

type GeneralInformationUsecase interface {
	GetByID(ctx context.Context, id int64) (GeneralInformation, error)
}
