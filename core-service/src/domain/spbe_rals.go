package domain

import "context"

type SpbeRals struct {
	ID       int64  `json:"id"`
	RalCode2 string `json:"ral_code_2"`
	Code     string `json:"code"`
	Item     string `json:"item"`
}

type SpbeRalsUsecase interface {
	Fetch(ctx context.Context) (res []SpbeRals, err error)
}

type SpbeRalsRepository interface {
	Fetch(ctx context.Context) (res []SpbeRals, err error)
}
