package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type logoUsecase struct {
	logoRepo       domain.LogoRepository
	cfg            *config.Config
	contextTimeout time.Duration
}

func NewLogoUsecase(logoRepo domain.LogoRepository, cfg *config.Config, timeout time.Duration) domain.LogoUsecase {
	return &logoUsecase{
		logoRepo:       logoRepo,
		cfg:            cfg,
		contextTimeout: timeout,
	}
}

func (i *logoUsecase) Fetch(c context.Context, params domain.Request) (res []domain.Logo, total int64, err error) {
	ctx, cancel := context.WithTimeout(c, i.contextTimeout)
	defer cancel()

	res, total, err = i.logoRepo.Fetch(ctx, params)
	return
}

func (i *logoUsecase) Store(c context.Context, body *domain.StoreLogoRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, i.contextTimeout)
	defer cancel()

	if err = i.logoRepo.Store(ctx, body); err != nil {
		return
	}

	return
}
