package usecase

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type mailTemplateUsecase struct {
	mailRepo       domain.TemplateRepository
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewMailTemplateUsecase(m domain.TemplateRepository, u domain.UserRepository, timeout time.Duration) domain.TemplateUsecase {
	return &mailTemplateUsecase{
		mailRepo:       m,
		userRepo:       u,
		contextTimeout: timeout,
	}
}

func (m *mailTemplateUsecase) GetByTemplate(ctx context.Context, id uuid.UUID, key string) (res domain.Template, err error) {
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
	bodyTmp := res.Body
	bodyTmp = strings.ReplaceAll(bodyTmp, "{name}", user.Name)
	bodyTmp = strings.ReplaceAll(bodyTmp, "{unitName}", user.UnitName)

	res.Body = bodyTmp

	return
}
