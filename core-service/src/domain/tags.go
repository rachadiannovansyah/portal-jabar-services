package domain

import "context"

// Tags ..
type Tags struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// TagsRepository ..
type TagsRepository interface {
	StoreTags(ctx context.Context, t *Tags) error
}
