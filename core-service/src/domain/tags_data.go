package domain

import "context"

// TagsData ..
type TagsData struct {
	DataID   int64  `json:"data_id"`
	TagsName string `json:"tags_name"`
	Type     string `json:"type"`
}

// TagsDataRepository ..
type TagsDataRepository interface {
	FetchTagsData(ctx context.Context, id int64) ([]TagsData, error)
}
