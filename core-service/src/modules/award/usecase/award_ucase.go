package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type awardUsecase struct {
	awardRepo      domain.AwardRepository
	contextTimeout time.Duration
}

// NewAwardUsecase will create new an awardUsecase object representation of domain.awardUsecase interface
func NewAwardUsecase(n domain.AwardRepository, timeout time.Duration) domain.AwardUsecase {
	return &awardUsecase{
		awardRepo:      n,
		contextTimeout: timeout,
	}
}

func (n *awardUsecase) Fetch(c context.Context, params *domain.Request) (res []domain.Award, total int64, err error) {

	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	res, total, err = n.awardRepo.Fetch(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	return
}

func (n *awardUsecase) GetByID(c context.Context, id int64) (res domain.Award, err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	res, err = n.awardRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}
