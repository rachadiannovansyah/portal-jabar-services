package domain

import "context"

// DataTag ..
type DataTag struct {
	ID      int64  `json:"id"`
	DataID  int64  `json:"data_id"`
	TagID   int64  `json:"tag_id"`
	TagName string `json:"tag_name"`
	Type    string `json:"type"`
}

// DataTagRepository ..
type DataTagRepository interface {
	FetchDataTags(ctx context.Context, id int64) ([]DataTag, error)
	StoreDataTag(ctx context.Context, dt *DataTag) error
	DeleteDataTag(ctx context.Context, id int64) error
}
