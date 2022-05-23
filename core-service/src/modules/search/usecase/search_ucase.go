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

func (n *searchUsecase) Store(c context.Context, indices string, body *domain.Search) (err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)

	defer cancel()

	err = n.searchRepo.Store(ctx, indices, body)
	if err != nil {
		return err
	}

	return
}

func (n *searchUsecase) Update(c context.Context, indices string, id int, body *domain.Search) (err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)

	defer cancel()

	err = n.searchRepo.Update(ctx, indices, id, body)
	if err != nil {
		return err
	}

	return
}

func (n *searchUsecase) Delete(c context.Context, indices string, id int, domain string) (err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)

	defer cancel()

	err = n.searchRepo.Delete(ctx, indices, id, domain)

	if err != nil {
		return err
	}

	return

}
