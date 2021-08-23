package usecases

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type informationUcase struct {
	informationRepo domain.InformationsRepo
	contextTimeout  time.Duration
}

func NewInformationUcase(repo domain.InformationsRepo, timeout time.Duration) domain.InformationsUcase {
	return &informationUcase{
		informationRepo: repo,
		contextTimeout:  timeout,
	}
}

func (usecase *informationUcase) FetchAll(c context.Context, params *domain.FetchInformationsRequest) (res []domain.Informations, total int64, err error) {

	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, total, err = usecase.informationRepo.FetchAll(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	return
}

func (usecase *informationUcase) FetchOne(c context.Context, id int64) (res domain.Informations, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err = usecase.informationRepo.FetchOne(ctx, id)
	if err != nil {
		return
	}

	return
}
