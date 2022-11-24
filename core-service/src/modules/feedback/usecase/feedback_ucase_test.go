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

var err error
var mockReqBody domain.Feedback

func TestMain(m *testing.M) {
	// prepare test
	err = faker.FakeData(&mockReqBody)

	// exec test
	m.Run()
}

func TestStore(t *testing.T) {
	suite := testSuite()
	u := ucase.NewFeedbackUsecase(&suite.feedbackRepoMock, suite.contextTimeout)

	t.Run("success", func(t *testing.T) {
		// mock expectation being called
		suite.feedbackRepoMock.On("Store", mock.Anything, mock.Anything).Return(nil).Once()
		err := u.Store(context.TODO(), &mockReqBody)

		// assertions
		assert.NoError(t, err)
		assert.Equal(t, mockReqBody, domain.Feedback{
			ID:          mockReqBody.ID,
			Rating:      mockReqBody.Rating,
			Compliments: mockReqBody.Compliments,
			Criticism:   mockReqBody.Criticism,
			Suggestions: mockReqBody.Suggestions,
			Sector:      mockReqBody.Sector,
			CreatedAt:   mockReqBody.CreatedAt,
		}, "expected result.")
	})

	t.Run("failure", func(t *testing.T) {
		// mock expectation being called
		suite.feedbackRepoMock.On("Store", mock.Anything, mock.Anything).Return(domain.ErrInternalServerError).Once()
		err := u.Store(context.TODO(), &mockReqBody)

		// assertions
		assert.Error(t, err)
		assert.NotNil(t, err)
	})
}
