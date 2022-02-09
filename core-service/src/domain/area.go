package domain

import "context"

// Area model ..
type Area struct {
	ID                   int64      `json:"id"`
	Depth                NullString `json:"depth" validate:"required"`
	Name                 NullString `json:"name" validate:"required"`
	ParentCodeKemendagri NullString `json:"parent_code_kemendagri" validate:"required"`
	CodeKemendagri       NullString `json:"code_kemendagri" validate:"required"`
	CodeBps              NullString `json:"code_bps" validate:"required"`
	Latitude             NullString `json:"latitude" validate:"required"`
	Longtitude           NullString `json:"longtitude" validate:"required"`
	Meta                 NullString `json:"meta" validate:"required"`
}

// AreaListResponse ...
type AreaListResponse struct {
	ID             int64      `json:"id"`
	Name           NullString `json:"name"`
	CodeKemendagri NullString `json:"code_kemendagri"`
}

// AreaUsecase ..
type AreaUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]Area, int64, error)
}

// AreaRepository ..
type AreaRepository interface {
	Fetch(ctx context.Context, params *Request) (new []Area, total int64, err error)
	GetByID(ctx context.Context, id int64) (Area, error)
}
