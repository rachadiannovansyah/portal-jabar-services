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

func (u *popUpBannerUsecase) GetByID(c context.Context, id int64) (res domain.PopUpBanner, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.popUpBannerRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (u *popUpBannerUsecase) Store(c context.Context, au *domain.JwtCustomClaims, body domain.StorePopUpBannerRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if err = u.popUpBannerRepo.Store(ctx, body); err != nil {
		return
	}

	return
}

func (u *popUpBannerUsecase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	if err = u.popUpBannerRepo.Delete(ctx, id); err != nil {
		return
	}

	return
}

func (n *popUpBannerUsecase) UpdateStatus(ctx context.Context, ID int64, status string) (err error) {
	// deactive existing active pop up banner
	if err = n.popUpBannerRepo.DeactiveStatus(ctx); err != nil {
		return
	}

	// update within selected banner to live publish
	is_live := int64(0)
	if status == "ACTIVE" {
		is_live = int64(1)
	}
	if err = n.popUpBannerRepo.UpdateStatus(ctx, ID, status, is_live); err != nil {
		return
	}

	return
}

func (n *popUpBannerUsecase) Update(ctx context.Context, au *domain.JwtCustomClaims, ID int64, body *domain.StorePopUpBannerRequest) (err error) {
	if err = n.popUpBannerRepo.Update(ctx, ID, body); err != nil {
		return
	}

	return
}
