package domain

import "context"

// DataTags ..
type DataTags struct {
	ID       int64  `json:"id"`
	DataID   int64  `json:"data_id"`
	TagsName string `json:"tags_name"`
	Type     string `json:"type"`
}

// DataTagsRepository ..
type DataTagsRepository interface {
	FetchDataTags(ctx context.Context, id int64) ([]DataTags, error)
	StoreDataTags(ctx context.Context, dt *DataTags) error
}
