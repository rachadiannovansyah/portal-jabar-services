package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-api/domain"
)

type newsUsecase struct {
	newsRepo       domain.NewsRepository
	contextTimeout time.Duration
}

// NewNewsUsecase will create new an newsUsecase object representation of domain.newsUsecase interface
func NewNewsUsecase(a domain.NewsRepository, timeout time.Duration) domain.NewsUsecase {
	return &newsUsecase{
		newsRepo:       a,
		contextTimeout: timeout,
	}
}

func (a *newsUsecase) Fetch(c context.Context, cursor string, num int64) (res []domain.News, nextCursor string, err error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, nextCursor, err = a.newsRepo.Fetch(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}

	return
}

func (a *newsUsecase) GetByID(c context.Context, id int64) (res domain.News, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.newsRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	if err != nil {
		return domain.News{}, err
	}
	return
}

func (a *newsUsecase) Update(c context.Context, ar *domain.News) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	ar.UpdatedAt = time.Now()
	return a.newsRepo.Update(ctx, ar)
}

func (a *newsUsecase) GetBySlug(c context.Context, slug string) (res domain.News, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	res, err = a.newsRepo.GetBySlug(ctx, slug)
	if err != nil {
		return
	}

	if err != nil {
		return domain.News{}, err
	}

	return
}

func (a *newsUsecase) Store(c context.Context, m *domain.News) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedNews, _ := a.GetBySlug(ctx, m.Slug)
	if existedNews != (domain.News{}) {
		return domain.ErrConflict
	}

	err = a.newsRepo.Store(ctx, m)
	return
}

func (a *newsUsecase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedNews, err := a.newsRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedNews == (domain.News{}) {
		return domain.ErrNotFound
	}
	return a.newsRepo.Delete(ctx, id)
}
