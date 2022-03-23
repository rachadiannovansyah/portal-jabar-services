package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type mailTemplateUsecase struct {
	mailRepo       domain.TemplateRepository
	contextTimeout time.Duration
}

func NewMailTemplateUsecase(m domain.TemplateRepository, timeout time.Duration) domain.TemplateUsecase {
	return &mailTemplateUsecase{
		mailRepo:       m,
		contextTimeout: timeout,
	}
}

func (m *mailTemplateUsecase) GetByTemplate(ctx context.Context, key string) (res domain.Template, err error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	res, err = m.mailRepo.GetByTemplate(ctx, key)
	if err != nil {
		return
	}

	return
}
