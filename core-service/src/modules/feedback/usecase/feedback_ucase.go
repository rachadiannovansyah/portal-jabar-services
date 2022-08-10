package usecase

import (
	"context"
	"time"

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

	// sanitizing the request
	f.Compliments = sanitize.Alpha(f.Compliments, true)
	f.Criticism = sanitize.Alpha(f.Criticism, true)
	f.Suggestions = sanitize.Alpha(f.Suggestions, true)
	f.Sector = sanitize.Alpha(f.Sector, true)
	f.CreatedAt = time.Now()

	err = u.feedbackRepo.Store(ctx, f)
	return
}
