package domain

import "context"

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

type AreaUseCase interface {
	GetByID(ctx context.Context, id int64) (Area, error)
}

type AreaRepository interface {
	GetByID(ctx context.Context, id int64) (Area, error)
}
