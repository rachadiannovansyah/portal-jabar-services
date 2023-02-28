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

func (u *spbeRalsUsecase) Fetch(c context.Context) (res []domain.SpbeRals, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.spbeRalsRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return
}
