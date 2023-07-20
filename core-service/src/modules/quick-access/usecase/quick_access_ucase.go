package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type quickAccessUsecase struct {
	quickAccessRepo domain.QuickAccessRepository
	cfg             *config.Config
	contextTimeout  time.Duration
}

func NewQuickAccessUsecase(quickAccessRepo domain.QuickAccessRepository, cfg *config.Config, timeout time.Duration) domain.QuickAccessUsecase {
	return &quickAccessUsecase{
		quickAccessRepo: quickAccessRepo,
		cfg:             cfg,
		contextTimeout:  timeout,
	}
}

func (i *quickAccessUsecase) Store(c context.Context, body *domain.StoreQuickAccess) (err error) {
	ctx, cancel := context.WithTimeout(c, i.contextTimeout)
	defer cancel()

	tx, _ := i.quickAccessRepo.GetTx(ctx)

	if err = i.quickAccessRepo.Store(ctx, body, tx); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}

func (i *quickAccessUsecase) Fetch(c context.Context, params domain.Request) (res []domain.QuickAccess, total int64, err error) {
	ctx, cancel := context.WithTimeout(c, i.contextTimeout)
	defer cancel()

	res, total, err = i.quickAccessRepo.Fetch(ctx, params)
	return
}

func (i *quickAccessUsecase) Delete(c context.Context, ID int64) (err error) {
	ctx, cancel := context.WithTimeout(c, i.contextTimeout)
	defer cancel()

	tx, _ := i.quickAccessRepo.GetTx(ctx)

	if _, err = i.quickAccessRepo.GetByID(ctx, ID); err != nil {
		return
	}

	if err = i.quickAccessRepo.Delete(ctx, ID, tx); err != nil {
		return
	}

	err = tx.Commit()

	return
}

func (i *quickAccessUsecase) GetByID(c context.Context, ID int64) (res domain.QuickAccess, err error) {
	ctx, cancel := context.WithTimeout(c, i.contextTimeout)
	defer cancel()

	res, err = i.quickAccessRepo.GetByID(ctx, ID)

	return
}

func (i *quickAccessUsecase) UpdateStatus(c context.Context, ID int64, body *domain.UpdateStatusQuickAccess) (err error) {
	ctx, cancel := context.WithTimeout(c, i.contextTimeout)
	defer cancel()

	tx, _ := i.quickAccessRepo.GetTx(ctx)

	if _, err = i.quickAccessRepo.GetByID(ctx, ID); err != nil {
		return
	}

	// for case request is active true then checking max active
	if *body.IsActive == int8(1) {
		if total := i.quickAccessRepo.CountByActived(ctx); total >= int64(domain.QuickAccessMaxActived) {
			err = domain.ErrBadRequest
			return
		}
	}

	if err = i.quickAccessRepo.UpdateStatus(ctx, ID, body, tx); err != nil {
		return
	}

	err = tx.Commit()
	return
}

func (i *quickAccessUsecase) Update(c context.Context, ID int64, body *domain.StoreQuickAccess) (err error) {
	ctx, cancel := context.WithTimeout(c, i.contextTimeout)
	defer cancel()

	tx, _ := i.quickAccessRepo.GetTx(ctx)

	if _, err = i.quickAccessRepo.GetByID(ctx, ID); err != nil {
		return
	}

	if err = i.quickAccessRepo.Update(ctx, ID, body, tx); err != nil {
		return
	}

	err = tx.Commit()
	return
}
