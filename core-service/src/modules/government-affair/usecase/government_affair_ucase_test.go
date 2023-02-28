package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mocks "github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain/mocks/repositories"
	ucase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/government-affair/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type governmentAffairUsecaseTestSuite struct {
	governmentAffairRepo    mocks.GovernmentAffairRepository
	cfg                     *config.Config
	ctxTimeout              time.Duration
	governmentAffairUsecase domain.GovernmentAffairUsecase
}

func testSuite() *governmentAffairUsecaseTestSuite {
	repoMock := mocks.GovernmentAffairRepository{}
	cfg := &config.Config{}
	ctxTimeout := time.Second * 60

	return &governmentAffairUsecaseTestSuite{
		governmentAffairRepo:    repoMock,
		cfg:                     cfg,
		ctxTimeout:              ctxTimeout,
		governmentAffairUsecase: ucase.NewGovernmentAffairUsecase(&repoMock, cfg, ctxTimeout),
	}
}

var ts *governmentAffairUsecaseTestSuite
var err error
var mockStruct domain.GovernmentAffair
var usecase domain.GovernmentAffairUsecase

func TestMain(m *testing.M) {
	// prepare test
	ts = testSuite()
	err = faker.FakeData(&mockStruct)
	usecase = ucase.NewGovernmentAffairUsecase(&ts.governmentAffairRepo, ts.cfg, ts.ctxTimeout)

	// execute test
	m.Run()
}

func TestFetch(t *testing.T) {
	mockList := make([]domain.GovernmentAffair, 0)
	mockList = append(mockList, mockStruct)

	t.Run("success", func(t *testing.T) {
		// mock expectation being called
		ts.governmentAffairRepo.On("Fetch", mock.Anything).Return(mockList, nil).Once()
		list, err := usecase.Fetch(context.TODO())

		// assertions
		assert.Equal(t, mockList, list)
		assert.NoError(t, err)

		ts.governmentAffairRepo.AssertExpectations(t)
		ts.governmentAffairRepo.AssertCalled(t, "Fetch", mock.Anything)
	})

	t.Run("error-occurred", func(t *testing.T) {
		// mock expectation being called
		ts.governmentAffairRepo.On("Fetch", mock.Anything).Return(nil, domain.ErrInternalServerError).Once()
		_, err := usecase.Fetch(context.TODO())

		// assertions
		assert.Error(t, err)
	})
}
