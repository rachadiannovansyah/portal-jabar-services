package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type publicServiceUsecase struct {
	publicServiceRepo domain.PublicServiceRepository
	userRepo          domain.UserRepository
	searchRepo        domain.SearchRepository
	cfg               *config.Config
	contextTimeout    time.Duration
}

// NewPublicServiceUsecase creates a new public-service usecase
func NewPublicServiceUsecase(ps domain.PublicServiceRepository, u domain.UserRepository, sr domain.SearchRepository, cfg *config.Config, timeout time.Duration) domain.PublicServiceUsecase {
	return &publicServiceUsecase{
		publicServiceRepo: ps,
		userRepo:          u,
		searchRepo:        sr,
		cfg:               cfg,
		contextTimeout:    timeout,
	}
}

func (u *publicServiceUsecase) Fetch(c context.Context, params *domain.Request) (res []domain.PublicService, err error) {

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.publicServiceRepo.Fetch(ctx, params)

	if err != nil {
		return nil, err
	}

	return
}

func (u *publicServiceUsecase) MetaFetch(c context.Context, params *domain.Request) (total int64, lastUpdated string, err error) {

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	total, lastUpdated, err = u.publicServiceRepo.MetaFetch(ctx, params)

	if err != nil {
		return 0, "", err
	}

	return
}

func (n *publicServiceUsecase) GetBySlug(c context.Context, slug string) (res domain.PublicService, err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	res, err = n.publicServiceRepo.GetBySlug(ctx, slug)
	if err != nil {
		return
	}

	return
}

func (n *publicServiceUsecase) Store(ctx context.Context, ps domain.StorePublicService) (err error) {
	err = n.publicServiceRepo.Store(ctx, ps)

	return
}
