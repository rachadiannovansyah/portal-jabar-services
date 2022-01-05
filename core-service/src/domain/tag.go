package domain

import "context"

// Tag struct ..
type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// TagRepository interface ..
type TagRepository interface {
	StoreTag(ctx context.Context, t *Tag) error
}
