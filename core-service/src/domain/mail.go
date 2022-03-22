package domain

import "context"

// Mail ..
type Mail struct {
	ID       int8   `json:"id"`
	From     string `json:"from"`
	To       string `json:"to"`
	Subject  string `json:"subject"`
	CC       string `json:"cc"`
	Body     string `json:"body"`
	Template string `json:"template"`
}

type MailUsecase interface {
	GetByTemplate(context.Context, string) (Mail, error)
}

type MailRepository interface {
	GetByTemplate(ctx context.Context, key string) (res Mail, err error)
}
