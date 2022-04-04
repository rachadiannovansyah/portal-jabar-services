package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

type templateUsecase struct {
	mailRepo       domain.TemplateRepository
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewTemplateUsecase(m domain.TemplateRepository, u domain.UserRepository, timeout time.Duration) domain.TemplateUsecase {
	return &templateUsecase{
		mailRepo:       m,
		userRepo:       u,
		contextTimeout: timeout,
	}
}

func (m *templateUsecase) GetByTemplate(ctx context.Context, id uuid.UUID, key string) (res domain.Template, err error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()

	res, err = m.mailRepo.GetByTemplate(ctx, key)
	if err != nil {
		return
	}

	user, err := m.userRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	// append for response get account submission
	bodyTmp := helpers.ReplaceBodyParams(res.Body, []string{user.Name, user.UnitName})

	res.Body = bodyTmp

	return
}
