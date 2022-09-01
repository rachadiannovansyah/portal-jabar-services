package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type publicServiceUsecase struct {
	publicServiceRepo domain.PublicServiceRepository
	userRepo          domain.UserRepository
	searchRepo        domain.SearchRepository
	cfg               *config.Config
	contextTimeout    time.Duration
}

// NewPublicServiceUsecase creates a new feedback usecase
func NewPublicServiceUsecase(ps domain.PublicServiceRepository, u domain.UserRepository, sr domain.SearchRepository, cfg *config.Config, timeout time.Duration) domain.PublicServiceUsecase {
	return &publicServiceUsecase{
		publicServiceRepo: ps,
		userRepo:          u,
		searchRepo:        sr,
		cfg:               cfg,
		contextTimeout:    timeout,
	}
}

func (p *publicServiceUsecase) Store(c context.Context, ps *domain.StorePserviceRequest) (err error) {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	ps.CreatedAt = time.Now()
	ps.UpdatedAt = time.Now()

	if err = p.publicServiceRepo.Store(ctx, ps); err != nil {
		return
	}

	// FIXME: make a function to prepare data for search index
	err = p.searchRepo.Store(ctx, p.cfg.ELastic.IndexContent, &domain.Search{
		ID:        int(ps.ID),
		Domain:    "public_service",
		Title:     ps.Name,
		Content:   ps.Description,
		Category:  ps.Category,
		Thumbnail: ps.Image,
		CreatedAt: ps.CreatedAt,
		UpdatedAt: ps.UpdatedAt,
		IsActive:  ps.IsActive == 1,
	})

	return
}
