package domain

import (
	"context"
	"database/sql"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
)

const (
	ArchiveStatus = "ARCHIVE"
)

type MasterDataService struct {
	ID                    int64                 `json:"id"`
	MainService           MainService           `json:"main_service"`
	Application           Application           `json:"application"`
	AdditionalInformation AdditionalInformation `json:"additional_information"`
	Status                string                `json:"status"`
	HasPublication        int8                  `json:"has_publication"`
	CreatedBy             User                  `json:"created_by"`
	UpdatedAt             time.Time             `json:"updated_at"`
	CreatedAt             time.Time             `json:"created_at"`
}

type ListMasterDataResponse struct {
	ID             int64     `json:"id"`
	ServiceName    string    `json:"service_name"`
	OpdName        string    `json:"opd_name"`
	ServiceUser    string    `json:"service_user"`
	Technical      string    `json:"technical"`
	UpdatedAt      time.Time `json:"updated_at"`
	Status         string    `json:"status"`
	HasPublication int8      `json:"has_publication"`
}

type StoreMasterDataService struct {
	ID       int64 `json:"id"`
	Services struct {
		ID          int64 `json:"id"`
		Information struct {
			OpdName             int64     `json:"opd_name"`
			GovernmentAffair    string    `json:"government_affair"`
			SubGovernmentAffair string    `json:"sub_government_affair"`
			ServiceForm         string    `json:"form"`
			ServiceType         string    `json:"type"`
			ServiceName         string    `json:"name"`
			ProgramName         string    `json:"program_name"`
			Description         string    `json:"description"`
			ServiceUser         string    `json:"user"`
			SubServiceSpbe      string    `json:"sub_service_spbe"`
			OperationalStatus   string    `json:"operational_status"`
			Technical           string    `json:"technical"`
			Benefits            MdsObject `json:"benefits"`
			Facilities          MdsObject `json:"facilities"`
			Website             string    `json:"website"`
			Links               []struct {
				Tautan string `json:"tautan"`
				Type   string `json:"type"`
				Label  string `json:"label"`
			} `json:"links"`
		} `json:"information"`
		ServiceDetail struct {
			TermsAndConditions MdsObjectCover `json:"terms_and_conditions"`
			ServiceProcedures  MdsObjectCover `json:"service_procedures"`
			ServiceFee         MdsServiceFee  `json:"service_fee"`
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
	Application           MdsApplication `json:"application" validate:"required"`
	AdditionalInformation struct {
		ID              int64  `json:"id"`
		ResponsibleName string `json:"responsible_name"`
		PhoneNumber     string `json:"phone_number"`
		Email           string `json:"email"`
		SocialMedia     []struct {
			Name string `json:"name"`
			Type string `json:"type" validate:"required,eq=GOOGLE_FORM|eq=|eq=GOOGLE_PLAYSTORE|eq=APP_STORE|eq=WEBSITE|eq=TIKTOK"`
			Link string `json:"link"`
		} `json:"social_media"`
	} `json:"additional_information" validate:"required"`
	Status    string `json:"status" validate:"required,eq=DRAFT|eq=ARCHIVE"`
	CreatedBy User   `json:"created_by"`
}

type MdsItems struct {
	Name  string              `json:"name,omitempty"`
	Image DetailMetaDataImage `json:"image,omitempty"`
}

type MdsItemCovers struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type MdsObject struct {
	Title    string     `json:"title"`
	IsActive int8       `json:"is_active"`
	Items    []MdsItems `json:"items"`
}

type MdsObjectCover struct {
	Cover    DetailMetaDataImage `json:"cover"`
	Title    string              `json:"title"`
	IsActive int8                `json:"is_active"`
	Items    []MdsItemCovers     `json:"items"`
}

type MdsApplication struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Title    string `json:"title"`
	Features []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"features"`
}

type DetailMasterDataServiceResponse struct {
	ID                    int64                       `json:"id"`
	MainService           MainServiceDetail           `json:"services"`
	Application           ApplicationDetail           `json:"application"`
	AdditionalInformation AdditionalInformationDetail `json:"additional_information"`
	Status                string                      `json:"status"`
	HasPublication        int8                        `json:"has_publication"`
	UpdatedAt             time.Time                   `json:"updated_at"`
	CreatedAt             time.Time                   `json:"created_at"`
}

type AdditionalInformationDetail struct {
	ID              int64            `json:"id"`
	ResponsibleName string           `json:"responsible_name"`
	PhoneNumber     string           `json:"phone_number"`
	Email           string           `json:"email"`
	SocialMedia     []SocialMediaMds `json:"social_media"`
}

type SocialMediaMds struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Link string `json:"link"`
}

type ApplicationDetail struct {
	ID       int64         `json:"id"`
	Name     string        `json:"name"`
	Status   string        `json:"status"`
	Title    string        `json:"title"`
	Features []FeaturesMds `json:"features"`
}

type FeaturesMds struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type MainServiceDetail struct {
	ID                  int64                `json:"id"`
	OpdName             string               `json:"opd_name"`
	GovernmentAffair    string               `json:"government_affair"`
	SubGovernmentAffair string               `json:"sub_government_affair"`
	ServiceForm         string               `json:"form"`
	ServiceType         string               `json:"type"`
	ServiceName         string               `json:"name"`
	ProgramName         string               `json:"program_name"`
	Description         string               `json:"description"`
	ServiceUser         string               `json:"user"`
	SubServiceSpbe      string               `json:"sub_service_spbe"`
	OperationalStatus   string               `json:"operational_status"`
	Technical           string               `json:"technical"`
	Benefits            MdsObject            `json:"benefits"`
	Facilities          MdsObject            `json:"facilities"`
	Website             string               `json:"website"`
	Links               []LinkMds            `json:"links"`
	TermsAndConditions  MdsObjectCover       `json:"terms_and_conditions"`
	ServiceProcedures   MdsObjectCover       `json:"service_procedures"`
	ServiceFee          MdsServiceFee        `json:"service_fee"`
	OperationalTimes    []OperationalTimeMds `json:"operational_times"`
	HotlineNumber       string               `json:"hotline_number"`
	HotlineMail         string               `json:"hotline_mail"`
	Locations           []LocationMds        `json:"locations"`
}

type MdsServiceFee struct {
	HasRange       int8   `json:"has_range"`
	MinimunFee     int64  `json:"minimum_fee"`
	MaximumFee     int64  `json:"maximum_fee"`
	HasDescription int8   `json:"has_description"`
	Description    string `json:"description"`
}

type LinkMds struct {
	Link  string `json:"tautan"`
	Type  string `json:"type"`
	Label string `json:"label"`
}

type OperationalTimeMds struct {
	Day   string `json:"day"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type LocationMds struct {
	Type         string `json:"type"`
	Organization string `json:"organization"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	PhoneNumber  string `json:"phone_number"`
}

type MasterDataServiceUsecaseArgs struct {
	MdsRepo        MasterDataServiceRepository
	MsRepo         MainServiceRepository
	ApRepo         ApplicationRepository
	AiRepo         AdditionalInformationRepository
	UserRepo       UserRepository
	Cfg            *config.Config
	ContextTimeout time.Duration
}

type MasterDataServiceEntityID struct {
	ID                      int64 `json:"mds_id"`
	MainServiceID           int64 `json:"main_service_id"`
	ApplicationID           int64 `json:"application_id"`
	AdditionalInformationID int64 `json:"additional_information_id"`
}

type MasterDataServiceUsecase interface {
	Store(ctx context.Context, au *JwtCustomClaims, body *StoreMasterDataService) (err error)
	Fetch(ctx context.Context, au *JwtCustomClaims, params *Request) (res []MasterDataService, total int64, err error)
	Delete(ctx context.Context, ID int64) (err error)
	GetByID(ctx context.Context, ID int64) (res MasterDataService, err error)
	Update(context.Context, *StoreMasterDataService, int64) (err error)
	TabStatus(ctx context.Context, au *JwtCustomClaims, params *Request) ([]TabStatusResponse, error)
	Archive(ctx context.Context, au *JwtCustomClaims, params *Request) (res []MasterDataService, err error)
}

type MasterDataServiceRepository interface {
	Store(ctx context.Context, body *StoreMasterDataService, tx *sql.Tx) (err error)
	GetTx(context.Context) (*sql.Tx, error)
	Fetch(ctx context.Context, params *Request) (res []MasterDataService, total int64, err error)
	Delete(ctx context.Context, ID int64) (err error)
	GetByID(ctx context.Context, ID int64) (res MasterDataService, err error)
	Update(context.Context, *StoreMasterDataService, *MasterDataServiceEntityID, *sql.Tx) (err error)
	TabStatus(ctx context.Context, params *Request) (res []TabStatusResponse, err error)
	Archive(ctx context.Context, params *Request) (res []MasterDataService, err error)
}
