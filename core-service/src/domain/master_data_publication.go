package domain

import (
	"context"
	"database/sql"
	"time"
)

type MasterDataPublication struct {
	ID                    int64                  `json:"id"`
	DefaultInformation    DefaultInformation     `json:"default_information"`
	ServiceDescription    ServiceDescription     `json:"service_description"`
	AdditionalInformation PublicationInformation `json:"additional_information"`
	Status                string                 `json:"status"`
	UpdatedAt             time.Time              `json:"updated_at"`
	CreatedAt             time.Time              `json:"created_at"`
}

type DefaultInformation struct {
}

type ServiceDescription struct {
}

type PublicationInformation struct {
}

type StoreMasterDataPublication struct {
	ID                 int64 `json:"id"`
	DefaultInformation struct {
		MdsID          int64     `json:"mds_id"`
		PortalCategory string    `json:"portal_category"`
		Slug           string    `json:"slug"`
		Benefits       MdsObject `json:"benefits"`
		Facilities     MdsObject `json:"facilities"`
	} `json:"default_information" validate:"required"`
	ServiceDescription struct {
		Cover              CoverPublication       `json:"cover"`
		Images             []DetailMetaDataImage  `json:"images"`
		InfoGraphics       PublicationInfographic `json:"infographics"`
		TermsAndConditions MdsObjectCover         `json:"terms_and_conditions"`
		ServiceProcedures  MdsObjectCover         `json:"service_procedures"`
	} `json:"service_description" validate:"required"`
	AdditionalInformation struct {
		Keywords []string       `json:"keywords"`
		FAQ      PublicationFAQ `json:"faq"`
	} `json:"additional_information" validate:"required"`
	Status    string    `json:"status" validate:"required,eq=PUBLISH|eq=ARCHIVE"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CoverPublication struct {
	Video string              `json:"video"`
	Image DetailMetaDataImage `json:"image"`
}

type PublicationInfographic struct {
	IsActive int8                  `json:"is_active"`
	Images   []DetailMetaDataImage `json:"images"`
}

type PublicationFAQ struct {
	IsActive int8 `json:"is_active"`
	Items    []struct {
		Question string `json:"question"`
		Answer   string `json:"answer"`
	}
}
type MasterDataPublicationUsecase interface {
	Store(ctx context.Context, body *StoreMasterDataPublication) (err error)
}

type MasterDataPublicationRepository interface {
	Store(ctx context.Context, body *StoreMasterDataPublication) (err error)
	GetTx(context.Context) (*sql.Tx, error)
}
