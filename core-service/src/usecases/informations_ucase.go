package usecases

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type informationsUcase struct {
	informationRepo domain.InformationsRepo
	categories      domain.CategoriesRepository
	contextTimeout  time.Duration
}

func NewInformationUcase(repo domain.InformationsRepo, ctg domain.CategoriesRepository, timeout time.Duration) domain.InformationsUcase {
	return &informationsUcase{
		informationRepo: repo,
		categories:      ctg,
		contextTimeout:  timeout,
	}
}

func (n *informationsUcase) fillCategoryDetails(c context.Context, data []domain.Informations) ([]domain.Informations, error) {
	g, ctx := errgroup.WithContext(c)

	// Get the category's id
	mapCategories := map[int64]domain.Category{}

	for _, infos := range data {
		mapCategories[infos.Category.ID] = domain.Category{}
	}

	// Using goroutine to fetch the category's detail
	chanCategory := make(chan domain.Category)
	for categoryID := range mapCategories {
		categoryID := categoryID
		g.Go(func() error {
			res, err := n.categories.GetByID(ctx, categoryID)
			if err != nil {
				return err
			}
			chanCategory <- res
			return nil
		})
	}

	go func() {
		err := g.Wait()
		if err != nil {
			logrus.Error(err)
			return
		}
		close(chanCategory)
	}()

	for category := range chanCategory {
		if category != (domain.Category{}) {
			mapCategories[category.ID] = category
		}
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	// merge the category's data
	for index, item := range data {
		if a, ok := mapCategories[item.Category.ID]; ok {
			data[index].Category = a
		}
	}

	return data, nil
}

func (usecase *informationsUcase) FetchAll(c context.Context, params *domain.FetchInformationsRequest) (res []domain.Informations, total int64, err error) {

	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, total, err = usecase.informationRepo.FetchAll(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	res, err = usecase.fillCategoryDetails(ctx, res)
	if err != nil {
		return nil, 0, err
	}

	return
}

func (usecase *informationsUcase) FetchOne(c context.Context, id int64) (res domain.Informations, err error) {
	ctx, cancel := context.WithTimeout(c, usecase.contextTimeout)
	defer cancel()

	res, err = usecase.informationRepo.FetchOne(ctx, id)
	if err != nil {
		return
	}

	resCategory, err := usecase.categories.GetByID(ctx, res.Category.ID)
	if err != nil {
		return domain.Informations{}, err
	}
	res.Category = resCategory

	return
}
