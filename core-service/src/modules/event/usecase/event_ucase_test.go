package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mocks "github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain/mocks/repositories"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/event/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetByTitle(t *testing.T) {
	mockEventRepo := mocks.EventRepository{}
	mockCategoryRepo := mocks.CategoryRepository{}
	mockTagRepo := mocks.TagRepository{}
	mockDataTagRepo := mocks.DataTagRepository{}
	mockUserRepo := mocks.UserRepository{}
	ctxTimeout := time.Second * 60

	var mockDataEvent domain.Event
	err := faker.FakeData(&mockDataEvent)

	assert.NoError(t, err)

	uc := usecase.NewEventUsecase(&mockEventRepo, &mockCategoryRepo, &mockTagRepo, &mockDataTagRepo, &mockUserRepo, ctxTimeout)
	t.Run("success", func(t *testing.T) {
		mockEventRepo.On("GetByTitle", mock.Anything, mock.Anything).Return(mockDataEvent, nil).Once()
		event, err := uc.GetByTitle(context.TODO(), mockDataEvent.Title)
		assert.EqualValues(t, mockDataEvent, event)
		assert.NoError(t, err)
	})
}
