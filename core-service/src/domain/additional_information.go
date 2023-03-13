package domain

import (
	"context"
	"database/sql"
)

type AdditionalInformation struct {
	ID              int64      `json:"id"`
	ResponsibleName string     `json:"responsible_name"`
	PhoneNumber     string     `json:"phone_number"`
	Email           string     `json:"email"`
	SocialMedia     NullString `json:"social_media"`
}

type AdditionalInformationRepository interface {
	Store(ctx context.Context, body *StoreMasterDataService, tx *sql.Tx) (ID int64, err error)
	Update(ctx context.Context, aID int64, body *StoreMasterDataService, tx *sql.Tx) (err error)
}
