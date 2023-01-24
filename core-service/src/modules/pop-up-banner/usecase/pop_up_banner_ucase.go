package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type popUpBannerUsecase struct {
	popUpBannerRepo domain.PopUpBannerRepository
	cfg             *config.Config
	contextTimeout  time.Duration
}

// NewPopUpBannerUsecase creates a new service-public usecase
func NewPopUpBannerUsecase(pb domain.PopUpBannerRepository, cfg *config.Config, timeout time.Duration) domain.PopUpBannerUsecase {
	return &popUpBannerUsecase{
		popUpBannerRepo: pb,
		cfg:             cfg,
		contextTimeout:  timeout,
	}
}

func (u *popUpBannerUsecase) Fetch(c context.Context, auth *domain.JwtCustomClaims, params *domain.Request) (res []domain.PopUpBanner, total int64, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, total, err = u.popUpBannerRepo.Fetch(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	return
}
