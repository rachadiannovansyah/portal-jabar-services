package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mocks "github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain/mocks/repositories"
	ucase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/feedback/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type feedbackUsecaseTestSuite struct {
	feedbackRepoMock mocks.FeedbackRepository
	contextTimeout   time.Duration
}

func testSuite() *feedbackUsecaseTestSuite {
	return &feedbackUsecaseTestSuite{
		feedbackRepoMock: mocks.FeedbackRepository{},
		contextTimeout:   60 * time.Second,
	}
}

func TestStore(t *testing.T) {
	suite := testSuite()

	var mockReqBody domain.Feedback
	err := faker.FakeData(&mockReqBody)
	assert.NoError(t, err)

	u := ucase.NewFeedbackUsecase(&suite.feedbackRepoMock, suite.contextTimeout)

	t.Run("success", func(t *testing.T) {
		suite.feedbackRepoMock.On("Store", mock.Anything, mock.Anything).Return(nil).Once()

		err := u.Store(context.TODO(), &mockReqBody)
		assert.NoError(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		suite.feedbackRepoMock.On("Store", mock.Anything, mock.Anything).Return(domain.ErrInternalServerError).Once()

		err := u.Store(context.TODO(), &mockReqBody)
		assert.NotNil(t, err)
	})
}
