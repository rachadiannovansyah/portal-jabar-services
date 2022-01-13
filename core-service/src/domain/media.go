package domain

import (
	"bytes"
	"context"
	"mime/multipart"
)

// MediaUsecase is an interface for media use cases
type MediaUsecase interface {
	Store(context.Context, *multipart.FileHeader, bytes.Buffer) error
}
