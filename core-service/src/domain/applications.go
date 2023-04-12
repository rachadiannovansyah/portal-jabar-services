package domain

import (
	"context"
	"database/sql"
)

type Application struct {
	ID       int64      `json:"id"`
	Name     string     `json:"name"`
	Status   string     `json:"status"`
	Title    NullString `json:"title"`
	Features NullString `json:"features"`
}

type ApplicationRepository interface {
	Store(ctx context.Context, body *StoreMasterDataService, tx *sql.Tx) (ID int64, err error)
	Update(ctx context.Context, apID int64, body *StoreMasterDataService, tx *sql.Tx) (err error)
}
