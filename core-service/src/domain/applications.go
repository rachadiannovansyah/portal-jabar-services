package domain

import "context"

type Application struct {
	ID       int64      `json:"id"`
	Name     string     `json:"name"`
	Status   string     `json:"status"`
	Features NullString `json:"features"`
}

type ApplicationRepository interface {
	Store(ctx context.Context, body *StoreMasterDataService) (ID int64, err error)
	Update(ctx context.Context, apID int64, body *StoreMasterDataService) (err error)
}
