package usecases

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
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
	err = u.feedbackRepo.Store(ctx, f)
	return
}
