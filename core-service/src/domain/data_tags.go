package domain

import "context"

// DataTags ..
type DataTags struct {
	DataID   int64  `json:"data_id"`
	TagsName string `json:"tags_name"`
	Type     string `json:"type"`
}

// DataTagsRepository ..
type DataTagsRepository interface {
	FetchDataTags(ctx context.Context, id int64) ([]DataTags, error)
}
