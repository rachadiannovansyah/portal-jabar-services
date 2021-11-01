package domain

import "context"

// TagsData ..
type TagsData struct {
	ID       int64      `json:"id"`
	Data     NullInt64  `json:"data_id"`
	Tags     Tags       `json:"tags_id"`
	TagsName NullString `json:"tags_name"`
	Type     NullString `json:"type"`
}

// TagsDataRepository ..
type TagsDataRepository interface {
	GetByName(ctx context.Context, name string) ([]TagsData, error)
}
