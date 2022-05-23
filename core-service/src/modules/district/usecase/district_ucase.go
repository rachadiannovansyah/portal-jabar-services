package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

// District Usecase ...
type districtUsecase struct {
	districtRepo   domain.DistrictRepository
	contextTimeout time.Duration
}

// NewDistrictUsecase creates a new district Usecase
func NewDistrictUsecase(dr domain.DistrictRepository, timeout time.Duration) domain.DistrictUsecase {
	return &districtUsecase{
		districtRepo:   dr,
		contextTimeout: timeout,
	}
}

// Fetch returns all districts
func (du *districtUsecase) Fetch(ctx context.Context, params *domain.Request) ([]domain.District, int64, error) {
	ctx, cancel := context.WithTimeout(ctx, du.contextTimeout)
	defer cancel()

	return du.districtRepo.Fetch(ctx, params)
}
