package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type tagUsecase struct {
	tagRepo        domain.TagRepository
	contextTimeout time.Duration
}

// NewTagUsecase will create new an tagUsecase object representation of domain.tagUsecase interface
func NewTagUsecase(n domain.TagRepository, timeout time.Duration) domain.TagUsecase {
	return &tagUsecase{
		tagRepo:        n,
		contextTimeout: timeout,
	}
}

func (n *tagUsecase) FetchTag(c context.Context, params *domain.Request) (res []domain.Tag, total int64, err error) {

	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	res, total, err = n.tagRepo.FetchTag(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	return
}
