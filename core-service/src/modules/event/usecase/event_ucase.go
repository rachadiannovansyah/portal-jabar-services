package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type eventUcase struct {
	eventRepo      domain.EventRepository
	categories     domain.CategoryRepository
	tagsRepo       domain.TagRepository
	dataTagRepo    domain.DataTagRepository
	contextTimeout time.Duration
}

// NewEventUsecase will create new an eventUsecase object representation of domain.eventUsecase interface
func NewEventUsecase(repo domain.EventRepository, ctg domain.CategoryRepository, tags domain.TagRepository, dtags domain.DataTagRepository, timeout time.Duration) domain.EventUsecase {
	return &eventUcase{
		eventRepo:      repo,
		categories:     ctg,
		contextTimeout: timeout,
		tagsRepo:       tags,
		dataTagRepo:    dtags,
	}
}

func (u *eventUcase) fillDataTags(c context.Context, data []domain.Event) ([]domain.Event, error) {
	g, ctx := errgroup.WithContext(c)

	// Get the tags from the tags domain
	mapTags := map[int64][]domain.DataTag{}

	for _, eventTag := range data {
		mapTags[eventTag.ID] = []domain.DataTag{}
	}

	// Using goroutine to fetch the list tags
	chanTags := make(chan []domain.DataTag)
	for idx := range mapTags {
		eventID := idx
		g.Go(func() (err error) {
			res, err := u.dataTagRepo.FetchDataTags(ctx, eventID)
			chanTags <- res
			return
		})
	}

	go func() {
		err := g.Wait()
		if err != nil {
			logrus.Error(err)
			return
		}
		close(chanTags)
	}()

	for listTags := range chanTags {
		eventTags := []domain.DataTag{}
		copier.Copy(&eventTags, &listTags)
		if len(listTags) < 1 {
			continue
		}
		mapTags[listTags[0].DataID] = eventTags
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	for index, element := range data {
		if tags, ok := mapTags[element.ID]; ok {
			data[index].Tags = tags
		}
	}

	return data, nil
}

// Fetch ...
func (u *eventUcase) Fetch(c context.Context, params *domain.Request) (res []domain.Event, total int64, err error) {

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, total, err = u.eventRepo.Fetch(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	res, err = u.fillDataTags(ctx, res)

	if err != nil {
		return nil, 0, err
	}

	return
}

// GetByID ...
func (u *eventUcase) GetByID(c context.Context, id int64) (res domain.Event, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.eventRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

// ListCalendar will get data without paginate
func (u *eventUcase) ListCalendar(c context.Context, params *domain.Request) (res []domain.Event, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.eventRepo.ListCalendar(ctx, params)
	if err != nil {
		return nil, err
	}

	return
}

// Store an events
func (u *eventUcase) Store(c context.Context, m *domain.StoreRequestEvent) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	err = u.eventRepo.Store(ctx, m)
	if err != nil {
		return
	}

	for _, tagName := range m.Tags {
		tag := &domain.Tag{
			Name: tagName,
		}
		err = u.tagsRepo.StoreTag(ctx, tag)
		if err != nil {
			return
		}

		dataTag := &domain.DataTag{
			DataID:  m.ID,
			TagID:   tag.ID,
			TagName: tagName,
			Type:    "events",
		}
		err = u.dataTagRepo.StoreDataTag(ctx, dataTag)
		if err != nil {
			return
		}
	}

	return
}

func (u *eventUcase) GetByTitle(c context.Context, title string) (res domain.Event, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.eventRepo.GetByTitle(ctx, title)
	if err != nil {
		return
	}

	return
}

func (u *eventUcase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	return u.eventRepo.Delete(ctx, id)
}

func (u *eventUcase) Update(c context.Context, id int64, body *domain.UpdateRequestEvent) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	body.UpdatedAt = time.Now()
	return u.eventRepo.Update(ctx, id, body)
}
