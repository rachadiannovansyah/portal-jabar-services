package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type governmentAffairUsecase struct {
	governmentAffairRepo domain.GovernmentAffairRepository
	cfg                  *config.Config
	contextTimeout       time.Duration
}

// NewGovernmentAffairUsecase creates a new service-public usecase
func NewGovernmentAffairUsecase(ga domain.GovernmentAffairRepository, cfg *config.Config, timeout time.Duration) domain.GovernmentAffairUsecase {
	return &governmentAffairUsecase{
		governmentAffairRepo: ga,
		cfg:                  cfg,
		contextTimeout:       timeout,
	}
}

func (u *governmentAffairUsecase) Fetch(c context.Context, params *domain.Request) (res []domain.GovernmentAffair, total int64, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, total, err = u.governmentAffairRepo.Fetch(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	return
}
