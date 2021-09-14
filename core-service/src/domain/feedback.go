package domain

import (
	"context"
	"time"
)

// Feedback is a domain model for feedback
type Feedback struct {
	ID          int64     `json:"id"`
	Rating      int8      `json:"rating"`
	Compliments string    `json:"compliments"`
	Criticism   string    `json:"criticism"`
	Suggestions string    `json:"suggestions"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// FeedbackUsecase is an interface for feedback use cases
type FeedbackUsecase interface {
	Store(context.Context, *Feedback) error
}

// FeedbackRepository represents a feedback repository
type FeedbackRepository interface {
	Store(ctx context.Context, a *Feedback) error
}
