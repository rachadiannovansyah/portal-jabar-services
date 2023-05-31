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

func (i *infographicBannerUsecase) Fetch(c context.Context, params domain.Request) (res []domain.InfographicBanner, total int64, err error) {
	ctx, cancel := context.WithTimeout(c, i.contextTimeout)
	defer cancel()

	res, total, err = i.infographicBannerRepo.Fetch(ctx, params)
	return
}

func (i *infographicBannerUsecase) Delete(c context.Context, ID int64) (err error) {
	ctx, cancel := context.WithTimeout(c, i.contextTimeout)
	defer cancel()

	tx, _ := i.infographicBannerRepo.GetTx(ctx)

	_, err = i.infographicBannerRepo.GetByID(ctx, ID, tx)
	if err != nil {
		return
	}

	if err = i.infographicBannerRepo.Delete(ctx, ID, tx); err != nil {
		return
	}

	if err = i.infographicBannerRepo.SyncSequence(ctx, 1, tx); err != nil {
		return
	}

	err = tx.Commit()

	return
}

func (i *infographicBannerUsecase) GetByID(c context.Context, ID int64) (res domain.InfographicBanner, err error) {
	ctx, cancel := context.WithTimeout(c, i.contextTimeout)
	defer cancel()

	tx, _ := i.infographicBannerRepo.GetTx(ctx)

	res, err = i.infographicBannerRepo.GetByID(ctx, ID, tx)
	if err != nil {
		return
	}

	err = tx.Commit()

	return
}

func (i *infographicBannerUsecase) UpdateStatus(c context.Context, ID int64, body *domain.UpdateStatusInfographicBanner) (err error) {
	ctx, cancel := context.WithTimeout(c, i.contextTimeout)
	defer cancel()

	tx, _ := i.infographicBannerRepo.GetTx(ctx)

	if _, err = i.infographicBannerRepo.GetByID(ctx, ID, tx); err != nil {
		return
	}

	if body.IsActive == 1 {
		i.infographicBannerRepo.SyncSequence(ctx, 2, tx)
		i.infographicBannerRepo.UpdateStatus(ctx, ID, body, tx)
	} else {
		i.infographicBannerRepo.UpdateStatus(ctx, ID, body, tx)
		i.infographicBannerRepo.SyncSequence(ctx, 1, tx)
	}

	err = tx.Commit()
	return
}
