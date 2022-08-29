package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type documentArchiveUsecase struct {
	documentArchiveRepo domain.DocumentArchiveRepository
	userRepo            domain.UserRepository
	cfg                 *config.Config
	contextTimeout      time.Duration
}

// NewDocumentArchiveUsecase will create new an documentArchiveUsecase object representation of domain.documentArchive interface
func NewDocumentArchiveUsecase(d domain.DocumentArchiveRepository, u domain.UserRepository, cfg *config.Config, timeout time.Duration) domain.DocumentArchiveUsecase {
	return &documentArchiveUsecase{
		documentArchiveRepo: d,
		userRepo:            u,
		cfg:                 cfg,
		contextTimeout:      timeout,
	}
}

func (n *documentArchiveUsecase) fillUserDetails(c context.Context, data []domain.DocumentArchive) ([]domain.DocumentArchive, error) {
	g, ctx := errgroup.WithContext(c)

	// Get the user's id
	mapUsers := map[uuid.UUID]domain.User{}

	for _, news := range data {
		mapUsers[news.CreatedBy.ID] = domain.User{}
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
		if a, ok := mapUsers[item.CreatedBy.ID]; ok {
			data[index].CreatedBy = a
		}
	}

	return data, nil
}

func (n *documentArchiveUsecase) get(c context.Context, params *domain.Request) (res []domain.DocumentArchive, total int64, err error) {
	ctx, cancel := context.WithTimeout(c, n.contextTimeout)
	defer cancel()

	res, total, err = n.documentArchiveRepo.Fetch(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	res, err = n.fillUserDetails(ctx, res)

	if err != nil {
		return nil, 0, err
	}

	if err != nil {
		return nil, 0, err
	}

	return
}

func (n *documentArchiveUsecase) Fetch(c context.Context, params *domain.Request) (res []domain.DocumentArchive, total int64, err error) {
	return n.get(c, params)
}
