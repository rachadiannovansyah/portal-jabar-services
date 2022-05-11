package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type searchUsecase struct {
	searchRepo     domain.SearchRepository
	config         *config.Config
	contextTimeout time.Duration
}

// NewSearchUsecase will create new an searchUsecase object representation of domain.searchUsecase interface
func NewSearchUsecase(s domain.SearchRepository, cfg *config.Config, timeout time.Duration) domain.SearchUsecase {
	return &searchUsecase{
		searchRepo:     s,
		config:         cfg,
		contextTimeout: timeout,
	}
}

func (n *searchUsecase) Fetch(c context.Context, params *domain.Request) (res []domain.SearchListResponse, tot int64, aggs interface{}, err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	res, tot, aggs, err = n.searchRepo.Fetch(ctx, n.config.ELastic.IndexContent, params)
	if err != nil {
		return nil, 0, nil, err
	}

	return
}

func (n *searchUsecase) SearchSuggestion(c context.Context, params *domain.Request) (res []domain.SuggestResponse, err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	res, err = n.searchRepo.SearchSuggestion(ctx, n.config.ELastic.IndexContent, params)
	if err != nil {
		return nil, err
	}

	return
}
