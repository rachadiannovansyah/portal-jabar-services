package domain

import (
	"context"

	"github.com/google/uuid"
)

// Template ..
type Template struct {
	ID      int8   `json:"id"`
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	CC      string `json:"cc"`
	Body    string `json:"body"`
	Key     string `json:"key"`
}

type TemplateUsecase interface {
	GetByTemplate(context.Context, uuid.UUID, string) (Template, error)
}

type TemplateRepository interface {
	GetByTemplate(ctx context.Context, key string) (res Template, err error)
}
