package domain

import (
	"bytes"
	"context"
	"mime/multipart"
)

type MediaResponse struct {
	FileName        string `json:"file_name"`
	FileDownloadUri string `json:"file_download_uri"`
	Size            int64  `json:"size"`
}

type DeleteMediaRequest struct {
	Domain string `json:"domain"`
	Key    string `json:"key"`
}

// MediaUsecase is an interface for media use cases
type MediaUsecase interface {
	Store(context.Context, *multipart.FileHeader, bytes.Buffer, string) (*MediaResponse, error)
	Delete(ctx context.Context, body *DeleteMediaRequest) (err error)
}
