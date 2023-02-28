package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type spbeRalsUsecase struct {
	spbeRalsRepo   domain.SpbeRalsRepository
	cfg            *config.Config
	contextTimeout time.Duration
}

// NewSpbeRalsUsecase creates a new spbe-rals usecase
func NewSpbeRalsUsecase(srals domain.SpbeRalsRepository, cfg *config.Config, timeout time.Duration) domain.SpbeRalsUsecase {
	return &spbeRalsUsecase{
		spbeRalsRepo:   srals,
		cfg:            cfg,
		contextTimeout: timeout,
	}
}

func (u *spbeRalsUsecase) Fetch(c context.Context, params *domain.Request) (res []domain.SpbeRals, total int64, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, total, err = u.spbeRalsRepo.Fetch(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	return
}
