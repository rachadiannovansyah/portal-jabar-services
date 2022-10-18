package domain

import "context"

type GeneralInformation struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Slug             string `json:"slug"`
	Category         string `json:"category"`
	Address          string `json:"address"`
	Unit             string `json:"unit"`
	Phone            string `json:"phone"`
	Logo             string `json:"logo"`
	OperationalHours string `json:"operationalHours"`
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
