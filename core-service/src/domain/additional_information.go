package domain

import "context"

type AdditionalInformation struct {
	ID              int64      `json:"id"`
	ResponsibleName string     `json:"responsible_name"`
	PhoneNumber     string     `json:"phone_number"`
	Email           string     `json:"email"`
	SocialMedia     NullString `json:"social_media"`
}

type AdditionalInformationRepository interface {
	Store(ctx context.Context, body *StoreMasterDataService) (ID int64, err error)
}
