package usecase

import (
	"context"
	"time"

	"github.com/forPelevin/gomoji"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/mrz1836/go-sanitize"
)

type feedbackUsecase struct {
	feedbackRepo   domain.FeedbackRepository
	contextTimeout time.Duration
}

// NewFeedbackUsecase creates a new feedback usecase
func NewFeedbackUsecase(f domain.FeedbackRepository, timeout time.Duration) domain.FeedbackUsecase {
	return &feedbackUsecase{
		feedbackRepo:   f,
		contextTimeout: timeout,
	}
}

func (u *feedbackUsecase) Store(c context.Context, f *domain.Feedback) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	// detects and remove if emojis contained on the request
	f.Compliments = gomoji.RemoveEmojis(f.Compliments)
	f.Criticism = gomoji.RemoveEmojis(f.Criticism)
	f.Suggestions = gomoji.RemoveEmojis(f.Suggestions)
	f.Sector = gomoji.RemoveEmojis(f.Sector)

	// sanitizing the request from script attack
	f.Compliments = sanitize.Scripts(f.Compliments)
	f.Criticism = sanitize.Scripts(f.Criticism)
	f.Suggestions = sanitize.Scripts(f.Suggestions)
	f.Sector = sanitize.Scripts(f.Sector)
	f.CreatedAt = time.Now()

	err = u.feedbackRepo.Store(ctx, f)
	return
}
