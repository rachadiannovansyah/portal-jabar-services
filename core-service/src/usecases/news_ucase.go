package usecases

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type contentUsecase struct {
	contentRepo    domain.NewsRepository
	contextTimeout time.Duration
}

// NewContentUsecase will create new an contentUsecase object representation of domain.contentUsecase interface
func NewContentUsecase(a domain.NewsRepository, timeout time.Duration) domain.NewsUsecase {
	return &contentUsecase{
		contentRepo:    a,
		contextTimeout: timeout,
	}
}

func (a *contentUsecase) Fetch(c context.Context, params *domain.FetchNewsRequest) (res []domain.News, total int64, err error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, total, err = a.contentRepo.Fetch(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	return
}

func (a *contentUsecase) GetByID(c context.Context, id int64) (res domain.News, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.contentRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}
