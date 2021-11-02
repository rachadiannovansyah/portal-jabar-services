package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type featuredProgramUsecase struct {
	featuredProgramRepo domain.FeaturedProgramRepository
	contextTimeout      time.Duration
}

// NewFeaturedProgramUsecase will create new an unitUsecase object representation of domain.unitUsecase interface
func NewFeaturedProgramUsecase(p domain.FeaturedProgramRepository, timeout time.Duration) domain.FeaturedProgramUsecase {
	return &featuredProgramUsecase{
		featuredProgramRepo: p,
		contextTimeout:      timeout,
	}
}

func (u *featuredProgramUsecase) Fetch(c context.Context, params *domain.Request) (res []domain.FeaturedProgram, err error) {

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.featuredProgramRepo.Fetch(ctx, params)
	if err != nil {
		return nil, err
	}

	return
}
