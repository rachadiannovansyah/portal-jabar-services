package usecase

import (
	"context"
	"github.com/jinzhu/copier"
	"time"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type newsUsecase struct {
	newsRepo       domain.NewsRepository
	categories     domain.CategoryRepository
	userRepo       domain.UserRepository
	contextTimeout time.Duration
}

// NewNewsUsecase will create new an newsUsecase object representation of domain.newsUsecase interface
func NewNewsUsecase(n domain.NewsRepository, nc domain.CategoryRepository, u domain.UserRepository, timeout time.Duration) domain.NewsUsecase {
	return &newsUsecase{
		newsRepo:       n,
		categories:     nc,
		userRepo:       u,
		contextTimeout: timeout,
	}
}

func (n *newsUsecase) fillAuthorDetails(c context.Context, data []domain.News) ([]domain.News, error) {
	g, ctx := errgroup.WithContext(c)

	// Get the user's id
	mapUsers := map[uuid.UUID]domain.User{}

	for _, news := range data {
		mapUsers[news.Author.ID] = domain.User{}
	}

	// Using goroutine to fetch the user's detail
	chanUser := make(chan domain.User)
	for authorID := range mapUsers {
		authorID := authorID
		g.Go(func() error {
			res, err := n.userRepo.GetByID(ctx, authorID)
			if err != nil {
				return err
			}
			chanUser <- res
			return nil
		})
	}

	go func() {
		err := g.Wait()
		if err != nil {
			logrus.Error(err)
			return
		}
		close(chanUser)
	}()

	for user := range chanUser {
		if user != (domain.User{}) {
			mapUsers[user.ID] = user
		}
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	// merge the user's data
	for index, item := range data {
		if a, ok := mapUsers[item.Author.ID]; ok {
			data[index].Author = a
		}
	}

	return data, nil
}

func (n *newsUsecase) fillRelatedNews(c context.Context, data []domain.NewsBanner) ([]domain.NewsBanner, error) {
	g, ctx := errgroup.WithContext(c)

	// Get the category
	mapCategories := map[string][]domain.NewsBanner{}

	for _, news := range data {
		mapCategories[news.Category] = []domain.NewsBanner{}
	}

	// Using goroutine to fetch the user's detail
	chanNews := make(chan []domain.News)
	for category := range mapCategories {
		g.Go(func() error {
			params := domain.Request{PerPage: 4}
			params.Filters = map[string]interface{}{
				"highlight": "0",
				"category":  category,
			}
			res, _, err := n.newsRepo.Fetch(ctx, &params)
			if err != nil {
				return err
			}

			chanNews <- res
			return nil
		})
	}

	go func() {
		err := g.Wait()
		if err != nil {
			logrus.Error(err)
			return
		}
		close(chanNews)
	}()

	for relatedNews := range chanNews {
		if len(relatedNews) < 1 {
			continue
		}
		relatedNewsBanner := []domain.NewsBanner{}
		copier.Copy(&relatedNewsBanner, &relatedNews)
		mapCategories[relatedNews[0].Category] = relatedNewsBanner
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	// merge the user's data
	for index, item := range data {
		if a, ok := mapCategories[item.Category]; ok {
			data[index].RelatedNews = a
		}
	}

	return data, nil
}

func (n *newsUsecase) Fetch(c context.Context, params *domain.Request) (res []domain.News, total int64, err error) {

	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	res, total, err = n.newsRepo.Fetch(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	res, err = n.fillAuthorDetails(ctx, res)
	if err != nil {
		return nil, 0, err
	}

	return
}

func (n *newsUsecase) GetByID(c context.Context, id int64) (res domain.News, err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	res, err = n.newsRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	resAuthor, err := n.userRepo.GetByID(ctx, res.Author.ID)
	if err != nil {
		return
	}
	res.Author = resAuthor

	// FIXME: prevent abuse page views counter by using cache (redis)
	err = n.newsRepo.AddView(ctx, id)
	if err != nil {
		logrus.Error(err)
	}

	return
}

func (n *newsUsecase) FetchNewsBanner(c context.Context) (res []domain.NewsBanner, err error) {

	news := []domain.News{}
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	news, err = n.newsRepo.FetchNewsBanner(ctx)
	if err != nil {
		return nil, err
	}

	news, err = n.fillAuthorDetails(ctx, news)
	if err != nil {
		return nil, err
	}

	copier.Copy(&res, &news)

	res, err = n.fillRelatedNews(ctx, res)
	if err != nil {
		return nil, err
	}

	return
}
