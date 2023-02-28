package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type updateCabdinUsecase struct {
	uptdCabdinRepo domain.UptdCabdinRepository
	cfg            *config.Config
	contextTimeout time.Duration
}

// NewUptdCabdinUsecase creates a new uptd-cabdin usecase
func NewUptdCabdinUsecase(uc domain.UptdCabdinRepository, cfg *config.Config, timeout time.Duration) domain.UptdCabdinUsecase {
	return &updateCabdinUsecase{
		uptdCabdinRepo: uc,
		cfg:            cfg,
		contextTimeout: timeout,
	}
}

func (u *updateCabdinUsecase) Fetch(c context.Context) (res []domain.UptdCabdin, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.uptdCabdinRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return
}
