package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mocks "github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain/mocks/repositories"
	ucase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/pop-up-banner/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type popUpBannerUsecaseTestSuite struct {
	popUpBannerRepo    mocks.PopUpBannerRepository
	cfg                *config.Config
	ctxTimeout         time.Duration
	popUpBannerUsecase domain.PopUpBannerUsecase
}

func testSuite() *popUpBannerUsecaseTestSuite {
	pbRepoMock := mocks.PopUpBannerRepository{}
	cfg := &config.Config{}
	ctxTimeout := time.Second * 60

	return &popUpBannerUsecaseTestSuite{
		popUpBannerRepo:    pbRepoMock,
		cfg:                cfg,
		ctxTimeout:         ctxTimeout,
		popUpBannerUsecase: ucase.NewPopUpBannerUsecase(&pbRepoMock, cfg, ctxTimeout),
	}
}

var ts *popUpBannerUsecaseTestSuite
var err error
var mockStruct domain.PopUpBanner
var mockJwtStruct domain.JwtCustomClaims
var usecase domain.PopUpBannerUsecase
var params *domain.Request

func TestMain(m *testing.M) {
	// prepare test
	ts = testSuite()
	err = faker.FakeData(&mockStruct)
	faker.FakeData(&mockJwtStruct)
	usecase = ucase.NewPopUpBannerUsecase(&ts.popUpBannerRepo, ts.cfg, ts.ctxTimeout)
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
	mockList := make([]domain.PopUpBanner, 0)
	mockList = append(mockList, mockStruct)

	t.Run("success", func(t *testing.T) {
		// mock expectation being called
		ts.popUpBannerRepo.On("Fetch", mock.Anything, mock.Anything).Return(mockList, int64(1), nil).Once()
		list, total, err := usecase.Fetch(context.TODO(), &mockJwtStruct, params)

		// assertions
		assert.Equal(t, mockList, list)
		assert.Len(t, list, int(total))
		assert.NoError(t, err)

		ts.popUpBannerRepo.AssertExpectations(t)
		ts.popUpBannerRepo.AssertCalled(t, "Fetch", mock.Anything, mock.Anything)
	})

	t.Run("error-occurred", func(t *testing.T) {
		// mock expectation being called
		ts.popUpBannerRepo.On("Fetch", mock.Anything, mock.Anything).Return(nil, int64(0), domain.ErrInternalServerError).Once()
		list, total, err := usecase.Fetch(context.TODO(), &mockJwtStruct, params)

		// assertions
		assert.Error(t, err)
		assert.Len(t, list, int(total))
	})
}
