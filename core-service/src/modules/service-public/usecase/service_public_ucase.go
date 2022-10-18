package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
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

	// get data from general information repository
	res, err = u.fillGeneralInformation(ctx, res)

	if err != nil {
		return
	}

	return
}

func (u *servicePublicUsecase) fillGeneralInformation(c context.Context, data []domain.ServicePublic) ([]domain.ServicePublic, error) {
	g, ctx := errgroup.WithContext(c)

	// Get the general information's id
	mapGenInfos := map[int64]domain.GeneralInformation{}

	for _, servPub := range data {
		mapGenInfos[servPub.GeneralInformation.ID] = domain.GeneralInformation{}
	}

	// Using goroutine to fetch the general information's detail
	chanGenInfo := make(chan domain.GeneralInformation)
	for genInfoID := range mapGenInfos {
		genInfoID := genInfoID
		g.Go(func() error {
			res, err := u.generalInformationRepo.GetByID(ctx, genInfoID)
			if err != nil {
				return err
			}
			chanGenInfo <- res
			return nil
		})
	}

	go func() {
		err := g.Wait()
		if err != nil {
			logrus.Error(err)
			return
		}
		close(chanGenInfo)
	}()

	for genInfo := range chanGenInfo {
		if genInfo != (domain.GeneralInformation{}) {
			mapGenInfos[genInfo.ID] = genInfo
		}
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	// merge the gen info's data
	for index, item := range data {
		if a, ok := mapGenInfos[item.GeneralInformation.ID]; ok {
			data[index].GeneralInformation = a
		}
	}

	return data, nil
}

func (u *servicePublicUsecase) MetaFetch(c context.Context, params *domain.Request) (total int64, lastUpdated string, err error) {

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	total, lastUpdated, err = u.servicePublicRepo.MetaFetch(ctx, params)

	if err != nil {
		return 0, "", err
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
