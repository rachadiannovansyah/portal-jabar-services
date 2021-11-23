package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type searchUsecase struct {
	searchRepo     domain.SearchRepository
	contextTimeout time.Duration
}

// NewSearchUsecase will create new an searchUsecase object representation of domain.searchUsecase interface
func NewSearchUsecase(s domain.SearchRepository, timeout time.Duration) domain.SearchUsecase {
	return &searchUsecase{
		searchRepo:     s,
		contextTimeout: timeout,
	}
}

func (n *searchUsecase) Fetch(c context.Context, params *domain.Request) (res []domain.SearchListResponse, tot int64, err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	res, tot, err = n.searchRepo.Fetch(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	return
}
