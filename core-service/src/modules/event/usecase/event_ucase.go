package usecase

import (
	"context"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type eventUcase struct {
	eventRepo      domain.EventRepository
	categories     domain.CategoryRepository
	tagsRepo       domain.DataTagsRepository
	contextTimeout time.Duration
}

// NewEventUsecase will create new an eventUsecase object representation of domain.eventUsecase interface
func NewEventUsecase(repo domain.EventRepository, ctg domain.CategoryRepository, tr domain.DataTagsRepository, timeout time.Duration) domain.EventUsecase {
	return &eventUcase{
		eventRepo:      repo,
		categories:     ctg,
		contextTimeout: timeout,
		tagsRepo:       tr,
	}
}

// Fetch ...
func (u *eventUcase) Fetch(c context.Context, params *domain.Request) (res []domain.Event, total int64, err error) {

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, total, err = u.eventRepo.Fetch(ctx, params)
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
func (u *eventUcase) Store(c context.Context, m *domain.StoreRequestEvent, dt *domain.DataTags) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	err = u.eventRepo.Store(ctx, m)
	err = u.tagsRepo.StoreDataTags(ctx, dt)

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

	existedEvent, err := u.eventRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	if existedEvent == (domain.Event{}) {
		return domain.ErrNotFound
	}

	return u.eventRepo.Delete(ctx, id)
}

func (u *eventUcase) Update(c context.Context, id int64, body *domain.UpdateRequestEvent) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	body.UpdatedAt = time.Now()
	return u.eventRepo.Update(ctx, id, body)
}
