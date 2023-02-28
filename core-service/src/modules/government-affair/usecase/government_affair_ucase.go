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

func (u *governmentAffairUsecase) Fetch(c context.Context) (res []domain.GovernmentAffair, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.governmentAffairRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return
}
