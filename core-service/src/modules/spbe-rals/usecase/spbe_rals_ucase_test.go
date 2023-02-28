package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mocks "github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain/mocks/repositories"
	ucase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/spbe-rals/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type spbeRalsUsecaseTestSuite struct {
	spbeRalsRepo    mocks.SpbeRalsRepository
	cfg             *config.Config
	ctxTimeout      time.Duration
	spbeRalsUsecase domain.SpbeRalsUsecase
}

func testSuite() *spbeRalsUsecaseTestSuite {
	repoMock := mocks.SpbeRalsRepository{}
	cfg := &config.Config{}
	ctxTimeout := time.Second * 60

	return &spbeRalsUsecaseTestSuite{
		spbeRalsRepo:    repoMock,
		cfg:             cfg,
		ctxTimeout:      ctxTimeout,
		spbeRalsUsecase: ucase.NewSpbeRalsUsecase(&repoMock, cfg, ctxTimeout),
	}
}

var ts *spbeRalsUsecaseTestSuite
var err error
var mockStruct domain.SpbeRals
var usecase domain.SpbeRalsUsecase
var params *domain.Request

func TestMain(m *testing.M) {
	// prepare test
	ts = testSuite()
	err = faker.FakeData(&mockStruct)
	usecase = ucase.NewSpbeRalsUsecase(&ts.spbeRalsRepo, ts.cfg, ts.ctxTimeout)
	params = &domain.Request{
		Keyword:   "",
		PerPage:   10,
		Offset:    0,
		SortBy:    "",
		SortOrder: "",
	}

	// execute test
	m.Run()
}

func TestFetch(t *testing.T) {
	mockList := make([]domain.SpbeRals, 0)
	mockList = append(mockList, mockStruct)

	t.Run("success", func(t *testing.T) {
		// mock expectation being called
		ts.spbeRalsRepo.On("Fetch", mock.Anything, mock.Anything).Return(mockList, int64(1), nil).Once()
		list, total, err := usecase.Fetch(context.TODO(), params)

		// assertions
		assert.Equal(t, mockList, list)
		assert.Len(t, list, int(total))
		assert.NoError(t, err)

		ts.spbeRalsRepo.AssertExpectations(t)
		ts.spbeRalsRepo.AssertCalled(t, "Fetch", mock.Anything, mock.Anything)
	})

	t.Run("error-occurred", func(t *testing.T) {
		// mock expectation being called
		ts.spbeRalsRepo.On("Fetch", mock.Anything, mock.Anything).Return(nil, int64(0), domain.ErrInternalServerError).Once()
		list, total, err := usecase.Fetch(context.TODO(), params)

		// assertions
		assert.Error(t, err)
		assert.Len(t, list, int(total))
	})
}
