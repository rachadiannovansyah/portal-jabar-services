package usecases

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type agendaUcase struct {
	agendaRepo     domain.AgendaRepository
	categories     domain.CategoryRepository
	contextTimeout time.Duration
}

func NewAgendaUsecase(repo domain.AgendaRepository, ctg domain.CategoryRepository, timeout time.Duration) domain.AgendaUsecase {
	return &agendaUcase{
		agendaRepo:     repo,
		categories:     ctg,
		contextTimeout: timeout,
	}
}

func (i *agendaUcase) fillCategoryDetails(c context.Context, data []domain.Agenda) ([]domain.Agenda, error) {
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
			res, err := i.categories.GetByID(ctx, categoryID)
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

// Fetch ...
func (i *agendaUcase) Fetch(c context.Context, params *domain.Request) (res []domain.Agenda, total int64, err error) {

	ctx, cancel := context.WithTimeout(c, i.contextTimeout)
	defer cancel()

	res, total, err = i.agendaRepo.Fetch(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	res, err = i.fillCategoryDetails(ctx, res)
	if err != nil {
		return nil, 0, err
	}

	return
}

// GetByID ...
func (i *agendaUcase) GetByID(c context.Context, id int64) (res domain.Agenda, err error) {
	ctx, cancel := context.WithTimeout(c, i.contextTimeout)
	defer cancel()

	res, err = i.agendaRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	resCategory, err := i.categories.GetByID(ctx, res.Category.ID)
	if err != nil {
		return domain.Agenda{}, err
	}
	res.Category = resCategory

	return
}
