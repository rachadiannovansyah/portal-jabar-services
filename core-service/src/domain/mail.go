package domain

import "context"

type Mail struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type MailRepository interface {
	Enqueue(ctx context.Context, mail Mail) error
}
