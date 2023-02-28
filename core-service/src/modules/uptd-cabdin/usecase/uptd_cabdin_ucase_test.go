package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mocks "github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain/mocks/repositories"
	ucase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/uptd-cabdin/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type uptdCabdinUsecaseTestSuite struct {
	uptdCabdinRepo    mocks.UptdCabdinRepository
	cfg               *config.Config
	ctxTimeout        time.Duration
	uptdCabdinUsecase domain.UptdCabdinUsecase
}

func testSuite() *uptdCabdinUsecaseTestSuite {
	repoMock := mocks.UptdCabdinRepository{}
	cfg := &config.Config{}
	ctxTimeout := time.Second * 60

	return &uptdCabdinUsecaseTestSuite{
		uptdCabdinRepo:    repoMock,
		cfg:               cfg,
		ctxTimeout:        ctxTimeout,
		uptdCabdinUsecase: ucase.NewUptdCabdinUsecase(&repoMock, cfg, ctxTimeout),
	}
}

var ts *uptdCabdinUsecaseTestSuite
var err error
var mockStruct domain.UptdCabdin
var usecase domain.UptdCabdinUsecase

func TestMain(m *testing.M) {
	// prepare test
	ts = testSuite()
	err = faker.FakeData(&mockStruct)
	usecase = ucase.NewUptdCabdinUsecase(&ts.uptdCabdinRepo, ts.cfg, ts.ctxTimeout)

	// execute test
	m.Run()
}

func TestFetch(t *testing.T) {
	mockList := make([]domain.UptdCabdin, 0)
	mockList = append(mockList, mockStruct)

	t.Run("success", func(t *testing.T) {
		// mock expectation being called
		ts.uptdCabdinRepo.On("Fetch", mock.Anything).Return(mockList, nil).Once()
		list, err := usecase.Fetch(context.TODO())

		// assertions
		assert.Equal(t, mockList, list)
		assert.NoError(t, err)

		ts.uptdCabdinRepo.AssertExpectations(t)
		ts.uptdCabdinRepo.AssertCalled(t, "Fetch", mock.Anything)
	})

	t.Run("error-occurred", func(t *testing.T) {
		// mock expectation being called
		ts.uptdCabdinRepo.On("Fetch", mock.Anything).Return(nil, domain.ErrInternalServerError).Once()
		_, err := usecase.Fetch(context.TODO())

		// assertions
		assert.Error(t, err)
	})
}
