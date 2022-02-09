package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type areaUsecase struct {
	areaRepo       domain.AreaRepository
	contextTimeout time.Duration
}

// NewAreaUsecase will create new an areaUsecase object representation of domain.areaUsecase interface
func NewAreaUsecase(a domain.AreaRepository, timeout time.Duration) domain.AreaUsecase {
	return &areaUsecase{
		areaRepo:       a,
		contextTimeout: timeout,
	}
}

func (n *areaUsecase) Fetch(c context.Context, params *domain.Request) (res []domain.Area, total int64, err error) {

	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	res, total, err = n.areaRepo.Fetch(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	return
}
