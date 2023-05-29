package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type infographicBannerUsecase struct {
	infographicBannerRepo domain.InfographicBannerRepository
	cfg                   *config.Config
	contextTimeout        time.Duration
}

func NewInfographicBannerUsecase(infographicBannerRepo domain.InfographicBannerRepository, cfg *config.Config, timeout time.Duration) domain.InfographicBannerUsecase {
	return &infographicBannerUsecase{
		infographicBannerRepo: infographicBannerRepo,
		cfg:                   cfg,
		contextTimeout:        timeout,
	}
}

func (i *infographicBannerUsecase) Store(c context.Context, body *domain.StoreInfographicBanner) (err error) {
	ctx, cancel := context.WithTimeout(c, i.contextTimeout)
	defer cancel()

	tx, _ := i.infographicBannerRepo.GetTx(ctx)

	if err = i.infographicBannerRepo.SyncSequence(ctx, 2, tx); err != nil {
		return
	}

	if err = i.infographicBannerRepo.Store(ctx, body, tx); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}
