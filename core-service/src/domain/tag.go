package domain

import "context"

// Tag struct ..
type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// TagUsecase represent the tag usecases
type TagUsecase interface {
	FetchTag(ctx context.Context, params *Request) ([]Tag, int64, error)
}

// TagRepository interface ..
type TagRepository interface {
	StoreTag(ctx context.Context, t *Tag) error
	FetchTag(ctx context.Context, param *Request) ([]Tag, int64, error)
	GetTagByName(ctx context.Context, name string) (Tag, error)
}
