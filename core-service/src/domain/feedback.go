package domain

import (
	"context"
	"time"
)

// Feedback is a domain model for feedback
type Feedback struct {
	ID          int64     `json:"id"`
	Rating      int8      `json:"rating" validate:"required"`
	Compliments string    `json:"compliments" validate:"required,max=1500"`
	Criticism   string    `json:"criticism" validate:"required,max=1500"`
	Suggestions string    `json:"suggestions" validate:"required,max=1500"`
	Sector      string    `json:"sector" validate:"required,max=1510"`
	CreatedAt   time.Time `json:"created_at"`
}

// FeedbackUsecase is an interface for feedback use cases
type FeedbackUsecase interface {
	Store(context.Context, *Feedback) error
}

// FeedbackRepository represents a feedback repository
type FeedbackRepository interface {
	Store(ctx context.Context, a *Feedback) error
}
