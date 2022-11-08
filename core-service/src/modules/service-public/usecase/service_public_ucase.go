package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

type servicePublicUsecase struct {
	servicePublicRepo      domain.ServicePublicRepository
	generalInformationRepo domain.GeneralInformationRepository
	userRepo               domain.UserRepository
	searchRepo             domain.SearchRepository
	cfg                    *config.Config
	contextTimeout         time.Duration
}

// NewServicePublicUsecase creates a new service-public usecase
func NewServicePublicUsecase(sp domain.ServicePublicRepository, g domain.GeneralInformationRepository, u domain.UserRepository, sr domain.SearchRepository, cfg *config.Config, timeout time.Duration) domain.ServicePublicUsecase {
	return &servicePublicUsecase{
		servicePublicRepo:      sp,
		generalInformationRepo: g,
		userRepo:               u,
		searchRepo:             sr,
		cfg:                    cfg,
		contextTimeout:         timeout,
	}
}

func (u *servicePublicUsecase) Fetch(c context.Context, params *domain.Request) (res []domain.ServicePublic, err error) {

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.servicePublicRepo.Fetch(ctx, params)

	if err != nil {
		return nil, err
	}

	return
}

func (u *servicePublicUsecase) MetaFetch(c context.Context, params *domain.Request) (total int64, lastUpdated string, staticCount int64, err error) {

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	total, lastUpdated, staticCount, err = u.servicePublicRepo.MetaFetch(ctx, params)

	if err != nil {
		return 0, "", 0, err
	}

	return
}

func (n *servicePublicUsecase) GetBySlug(c context.Context, slug string) (res domain.ServicePublic, err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	res, err = n.servicePublicRepo.GetBySlug(ctx, slug)
	if err != nil {
		return
	}

	if err != nil {
		return
	}

	return
}

func (n *servicePublicUsecase) Store(ctx context.Context, ps domain.StorePublicService) (err error) {
	tx, err := n.generalInformationRepo.GetTx(ctx)
	if err != nil {
		return
	}

	// store gen info first
	genInfoID, err := n.generalInformationRepo.Store(ctx, ps, tx)
	if err != nil {
		return
	}

	// assign id geninfo to service public
	ps.GeneralInformation.ID = genInfoID

	// store it on service public
	err = n.servicePublicRepo.Store(ctx, ps, tx)

	if err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}

func (n *servicePublicUsecase) Delete(ctx context.Context, ID int64) (err error) {
	err = n.servicePublicRepo.Delete(ctx, ID)

	return
}

func (n *servicePublicUsecase) Update(ctx context.Context, UPs domain.UpdatePublicService, ID int64) (err error) {
	tx, err := n.generalInformationRepo.GetTx(ctx)
	if err != nil {
		return
	}

	ps, err := n.servicePublicRepo.GetByID(ctx, ID)
	if err != nil {
		return
	}

	slug := helpers.MakeSlug(UPs.GeneralInformation.Name, ps.GeneralInformation.ID)
	UPs.GeneralInformation.Slug = slug

	if err = n.generalInformationRepo.Update(ctx, UPs, ID, tx); err != nil {
		return
	}

	if err = n.servicePublicRepo.Update(ctx, UPs, ps.ID, tx); err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}
