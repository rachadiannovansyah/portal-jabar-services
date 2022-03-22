package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type mailUsecase struct {
	mailRepo       domain.MailRepository
	contextTimeout time.Duration
}

func NewMailUsecase(m domain.MailRepository, timeout time.Duration) domain.MailUsecase {
	return &mailUsecase{
		mailRepo:       m,
		contextTimeout: timeout,
	}
}

func (m *mailUsecase) GetByTemplate(ctx context.Context, key string) (res domain.Mail, err error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	res, err = m.mailRepo.GetByTemplate(ctx, key)
	if err != nil {
		return
	}

	return
}
