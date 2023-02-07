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

func TestGetByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// mock expectation being called
		ts.popUpBannerRepo.On("GetByID", mock.Anything, mock.Anything).Return(mockStruct, nil).Once()
		obj, err := usecase.GetByID(context.TODO(), int64(mockStruct.ID))

		// assertions
		assert.Equal(t, mockStruct, obj)
		assert.NoError(t, err)

		ts.popUpBannerRepo.AssertExpectations(t)
		ts.popUpBannerRepo.AssertCalled(t, "GetByID", mock.Anything, mock.AnythingOfType("int64"))
	})

	t.Run("error-occurred", func(t *testing.T) {
		// mock expectation being called
		ts.popUpBannerRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.PopUpBanner{}, domain.ErrInternalServerError).Once()
		_, err := usecase.GetByID(context.TODO(), mockStruct.ID)

		// assertions
		assert.Error(t, err)
	})
}

func TestStore(t *testing.T) {
	var mockRequest domain.StorePopUpBannerRequest
	err = faker.FakeData(&mockRequest)
	t.Run("success", func(t *testing.T) {
		// mock expectation being called
		ts.popUpBannerRepo.On("Store", mock.Anything, mock.Anything).Return(nil).Once()

		err := usecase.Store(context.TODO(), &mockJwtStruct, mockRequest)

		// assertions
		assert.NoError(t, err)
		ts.popUpBannerRepo.AssertExpectations(t)
		ts.popUpBannerRepo.AssertCalled(t, "Store", mock.Anything, mock.Anything, mock.Anything)
	})

	t.Run("error-occurred", func(t *testing.T) {
		// mock expectation being called
		ts.popUpBannerRepo.On("Store", mock.Anything, mock.Anything).Return(domain.ErrInternalServerError).Once()

		err := usecase.Store(context.TODO(), &mockJwtStruct, mockRequest)

		// assertions
		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// mock expectation being called
		ts.popUpBannerRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()
		err := usecase.Delete(context.TODO(), mockStruct.ID)

		// assertions
		assert.NoError(t, err)
		ts.popUpBannerRepo.AssertExpectations(t)
		ts.popUpBannerRepo.AssertCalled(t, "Delete", mock.Anything, mock.AnythingOfType("int64"))
	})

	t.Run("error-occurred", func(t *testing.T) {
		// mock expectation being called
		ts.popUpBannerRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(domain.ErrInternalServerError).Once()
		err := usecase.Delete(context.TODO(), mockStruct.ID)

		// assertions
		assert.Error(t, err)
	})
}

func TestUpdateStatus(t *testing.T) {
	var mockUpdateStruct domain.UpdateStatusPopUpBannerRequest
	err = faker.FakeData(&mockUpdateStruct)

	t.Run("success", func(t *testing.T) {
		// mock expectation being called
		ts.popUpBannerRepo.On("DeactiveStatus", mock.Anything).Return(nil).Once()
		ts.popUpBannerRepo.On("GetByID", mock.Anything, mock.Anything).Return(mockStruct, nil).Once()
		ts.popUpBannerRepo.On("UpdateStatus", mock.Anything, mock.AnythingOfType("int64"), mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("int64")).Return(nil)

		err := usecase.UpdateStatus(context.TODO(), mockStruct.ID, mockUpdateStruct.Status)

		// assertions
		assert.NoError(t, err)
		ts.popUpBannerRepo.AssertCalled(t, "DeactiveStatus", mock.Anything)
		ts.popUpBannerRepo.AssertCalled(t, "GetByID", mock.Anything, mock.Anything)
		ts.popUpBannerRepo.AssertCalled(t, "UpdateStatus", mock.Anything, mock.AnythingOfType("int64"), mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("int64"))
	})

	t.Run("error-occurred", func(t *testing.T) {
		// mock expectation being called
		ts.popUpBannerRepo.On("DeactiveStatus", mock.Anything).Return(domain.ErrInternalServerError).Once()
		ts.popUpBannerRepo.On("GetByID", mock.Anything, mock.Anything).Return(mockStruct, domain.ErrNotFound).Once()
		ts.popUpBannerRepo.On("UpdateStatus", mock.Anything, mock.AnythingOfType("int64"), mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("int64")).Return(domain.ErrNotFound)

		err := usecase.UpdateStatus(context.TODO(), mockStruct.ID, "NON-ACTIVE")

		// assertions
		assert.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	var mockUpdateStruct domain.StorePopUpBannerRequest
	err = faker.FakeData(&mockUpdateStruct)

	t.Run("success", func(t *testing.T) {
		// mock expectation being called
		ts.popUpBannerRepo.On("Update", mock.Anything, mock.AnythingOfType("int64"), mock.Anything).Return(nil).Once()

		err := usecase.Update(context.TODO(), &mockJwtStruct, mockStruct.ID, &mockUpdateStruct)

		// assertions
		assert.NoError(t, err)
		ts.popUpBannerRepo.AssertCalled(t, "Update", mock.Anything, mock.AnythingOfType("int64"), mock.Anything)
	})

	t.Run("error-occurred", func(t *testing.T) {
		// mock expectation being called
		ts.popUpBannerRepo.On("Update", mock.Anything, mock.AnythingOfType("int64"), mock.Anything).Return(domain.ErrNotFound).Once()

		err := usecase.Update(context.TODO(), &mockJwtStruct, mockStruct.ID+1, &mockUpdateStruct)

		// assertions
		assert.Error(t, err)
	})
}
