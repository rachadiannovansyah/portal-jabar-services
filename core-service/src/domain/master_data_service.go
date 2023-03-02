package domain

import (
	"context"
	"database/sql"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
)

type MasterDataService struct {
	ID                    int64                 `json:"id"`
	MainService           MainService           `json:"main_service"`
	Application           Application           `json:"application"`
	AdditionalInformation AdditionalInformation `json:"additional_information"`
}

type StoreMasterDataService struct {
	ID       int64 `json:"id"`
	Services struct {
		ID          int64 `json:"id"`
		Information struct {
			OpdName             int64    `json:"opd_name"`
			GovernmentAffair    string   `json:"government_affair"`
			SubGovernmentAffair string   `json:"sub_government_affair"`
			ServiceForm         string   `json:"form"`
			ServiceType         string   `json:"type"`
			SubServiceType      string   `json:"sub_service_type"`
			ServiceName         string   `json:"name"`
			ProgramName         string   `json:"program_name"`
			Description         string   `json:"description"`
			ServiceUser         string   `json:"user"`
			SubServiceSpbe      string   `json:"sub_service_spbe"`
			OperationalStatus   string   `json:"operational_status"`
			Technical           string   `json:"technical"`
			Benefits            []string `json:"benefits"`
			Facilities          []string `json:"facilities"`
			Website             string   `json:"website"`
			Links               []struct {
				Tautan string `json:"tautan"`
				Type   string `json:"type"`
				Label  string `json:"label"`
			} `json:"links"`
		} `json:"information"`
		ServiceDetail struct {
			TermsAndConditions []string `json:"terms_and_conditions"`
			ServiceProcedures  []string `json:"service_procedures"`
			ServiceFee         string   `json:"service_fee"`
			OperationalTime    []struct {
				Day   string `json:"day"`
				Start string `json:"start"`
				End   string `json:"end"`
			} `json:"operational_time"`
			HotlineNumber string `json:"hotline_number"`
			HotlineMail   string `json:"hotline_mail"`
		} `json:"service_detail"`
		Location []struct {
			Type         string `json:"type"`
			Organization string `json:"organization"`
			Name         string `json:"name"`
			Address      string `json:"address"`
			PhoneNumber  string `json:"phone_number"`
		} `json:"location"`
	} `json:"services" validate:"required"`
	Application struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		Status   string `json:"status"`
		Features []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"features"`
	} `json:"application" validate:"required"`
	AdditionalInformation struct {
		ID              int64  `json:"id"`
		ResponsibleName string `json:"responsible_name"`
		PhoneNumber     string `json:"phone_number"`
		Email           string `json:"email"`
		SocialMedia     []struct {
			Name string `json:"name"`
			Type string `json:"type" validate:"required,eq=GOOGLE_FORM|eq=|eq=GOOGLE_PLAYSTORE|eq=APP_STORE|eq=WEBSITE"`
			Link string `json:"link"`
		} `json:"social_media"`
	} `json:"additional_information" validate:"required"`
}

type MasterDataServiceUsecaseArgs struct {
	MdsRepo        MasterDataServiceRepository
	MsRepo         MainServiceRepository
	ApRepo         ApplicationRepository
	AiRepo         AdditionalInformationRepository
	Cfg            *config.Config
	ContextTimeout time.Duration
}

type MasterDataServiceUsecase interface {
	Store(ctx context.Context, au *JwtCustomClaims, body *StoreMasterDataService) (err error)
}

type MasterDataServiceRepository interface {
	Store(ctx context.Context, body *StoreMasterDataService, tx *sql.Tx) (err error)
	GetTx(context.Context) (*sql.Tx, error)
}
