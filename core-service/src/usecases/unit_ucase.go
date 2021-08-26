package usecases

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type unitUsecase struct {
	unitRepo       domain.UnitRepository
	contextTimeout time.Duration
}

// NewUnitUsecase will create new an unitUsecase object representation of domain.unitUsecase interface
func NewUnitUsecase(n domain.UnitRepository, timeout time.Duration) domain.UnitUsecase {
	return &unitUsecase{
		unitRepo:       n,
		contextTimeout: timeout,
	}
}

func (n *unitUsecase) Fetch(c context.Context, params *domain.Request) (res []domain.Unit, total int64, err error) {

	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	res, total, err = n.unitRepo.Fetch(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	return
}

func (n *unitUsecase) GetByID(c context.Context, id int64) (res domain.Unit, err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	res, err = n.unitRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}
