package domain

import (
	"context"
	"database/sql"
)

type MainService struct {
	ID                  int64      `json:"id"`
	OpdName             string     `json:"opd_name"`
	GovernmentAffair    string     `json:"government_affair"`
	SubGovernmentAffair string     `json:"sub_government_affair"`
	ServiceForm         string     `json:"form"`
	ServiceType         string     `json:"type"`
	SubServiceType      string     `json:"sub_service_type"`
	ServiceName         string     `json:"name"`
	ProgramName         string     `json:"program_name"`
	Description         string     `json:"description"`
	ServiceUser         string     `json:"user"`
	SubServiceSpbe      string     `json:"sub_service_spbe"`
	OperationalStatus   string     `json:"operational_status"`
	Technical           string     `json:"technical"`
	Benefits            string     `json:"benefits"`
	Facilities          string     `json:"facilities"`
	Website             string     `json:"website"`
	Links               NullString `json:"links"`
	TermsAndConditions  string     `json:"terms_and_conditions"`
	ServiceProcedures   string     `json:"service_procedures"`
	ServiceFee          string     `json:"service_fee"`
	OperationalTimes    NullString `json:"operational_times"`
	HotlineNumber       string     `json:"hotline_number"`
	HotlineMail         string     `json:"hotline_mail"`
	Locations           NullString `json:"locations"`
}

type MainServiceRepository interface {
	Store(ctx context.Context, body *StoreMasterDataService, tx *sql.Tx) (ID int64, err error)
	Update(ctx context.Context, msID int64, body *StoreMasterDataService, tx *sql.Tx) (err error)
	UpdateFromPublication(ctx context.Context, msID int64, body *StoreMasterDataService, tx *sql.Tx) (err error)
}
