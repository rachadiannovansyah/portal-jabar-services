package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mocks "github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain/mocks/repositories"
	ucase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/service-public/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type servicePublicUsecaseTestSuite struct {
	servicePublicRepo    mocks.ServicePublicRepository
	userRepoMock         mocks.UserRepository
	genInfoRepoMock      mocks.GeneralInformationRepository
	searchRepoMock       mocks.SearchRepository
	cfg                  *config.Config
	ctxTimeout           time.Duration
	servicePublicUsecase domain.ServicePublicUsecase
}

func testSuite() *servicePublicUsecaseTestSuite {
	spRepoMock := mocks.ServicePublicRepository{}
	userRepoMock := mocks.UserRepository{}
	genInfoRepoMock := mocks.GeneralInformationRepository{}
	searchRepoMock := mocks.SearchRepository{}
	cfg := &config.Config{}
	ctxTimeout := time.Second * 60

	return &servicePublicUsecaseTestSuite{
		servicePublicRepo:    spRepoMock,
		userRepoMock:         userRepoMock,
		genInfoRepoMock:      genInfoRepoMock,
		searchRepoMock:       searchRepoMock,
		cfg:                  cfg,
		ctxTimeout:           ctxTimeout,
		servicePublicUsecase: ucase.NewServicePublicUsecase(&spRepoMock, &genInfoRepoMock, &userRepoMock, &searchRepoMock, cfg, ctxTimeout),
	}
}

var suite *servicePublicUsecaseTestSuite
var err error
var mockStruct domain.ServicePublic
var usecase domain.ServicePublicUsecase
var params *domain.Request

func TestMain(m *testing.M) {
	// prepare test
	suite = testSuite()
	err = faker.FakeData(&mockStruct)
	usecase = ucase.NewServicePublicUsecase(&suite.servicePublicRepo, &suite.genInfoRepoMock, &suite.userRepoMock, &suite.searchRepoMock, suite.cfg, suite.ctxTimeout)
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
	mockList := make([]domain.ServicePublic, 0)
	mockList = append(mockList, mockStruct)

	t.Run("success", func(t *testing.T) {
		// mock expectation being called
		suite.servicePublicRepo.On("Fetch", mock.Anything, mock.Anything).Return(mockList, nil).Once()
		list, err := usecase.Fetch(context.TODO(), params)

		// assertions
		assert.NoError(t, err)
		assert.Len(t, list, len(mockList))
		assert.Equal(t, mockList, list)
		suite.servicePublicRepo.AssertExpectations(t)
		suite.servicePublicRepo.AssertCalled(t, "Fetch", mock.Anything, mock.Anything)
	})

	t.Run("error-occurred", func(t *testing.T) {
		// mock expectation being called
		suite.servicePublicRepo.On("Fetch", mock.Anything, mock.Anything).Return(nil, domain.ErrInternalServerError).Once()
		list, err := usecase.Fetch(context.TODO(), params)

		// assertions
		assert.Error(t, err)
		assert.Len(t, list, 0)
	})
}

func TestMetaFetch(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// mock expectation being called
		suite.servicePublicRepo.On("MetaFetch", mock.Anything, mock.Anything).Return(int64(1), "2022-11-15", int64(2), nil).Once()
		total, lastUpdate, staticCount, err := usecase.MetaFetch(context.TODO(), params)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), total)
		assert.Equal(t, "2022-11-15", lastUpdate)
		assert.Equal(t, int64(2), staticCount)
	})

	t.Run("error-occurred", func(t *testing.T) {
		// mock expectation being called
		suite.servicePublicRepo.On("MetaFetch", mock.Anything, mock.Anything).Return(int64(0), "", int64(0), domain.ErrInternalServerError).Once()
		total, lastUpdate, staticCount, err := usecase.MetaFetch(context.TODO(), params)
		assert.Error(t, err)
		assert.Equal(t, int64(0), total)
		assert.Equal(t, "", lastUpdate)
		assert.Equal(t, int64(0), staticCount)
	})
}

func TestGetBySlug(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// mock expectation being called
		suite.servicePublicRepo.On("GetBySlug", mock.Anything, mock.AnythingOfType("string")).Return(mockStruct, nil).Once()
		obj, err := usecase.GetBySlug(context.TODO(), mockStruct.GeneralInformation.Slug)

		// assertions
		assert.NoError(t, err)
		assert.NotNil(t, obj)
		assert.Equal(t, mockStruct, obj)
		suite.servicePublicRepo.AssertExpectations(t)
		suite.servicePublicRepo.AssertCalled(t, "GetBySlug", mock.Anything, mock.AnythingOfType("string"))
	})

	t.Run("error-occurred", func(t *testing.T) {
		// mock expectation being called
		suite.servicePublicRepo.On("GetBySlug", mock.Anything, mock.AnythingOfType("string")).Return(domain.ServicePublic{}, domain.ErrInternalServerError).Once()
		obj, err := usecase.GetBySlug(context.TODO(), mockStruct.GeneralInformation.Slug)

		// assertions
		assert.Error(t, err)
		assert.Equal(t, domain.ServicePublic{}, obj)
		suite.servicePublicRepo.AssertExpectations(t)
		suite.servicePublicRepo.AssertCalled(t, "GetBySlug", mock.Anything, mock.AnythingOfType("string"))
	})
}

func TestDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// mock expectation being called
		suite.servicePublicRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockStruct, nil)
		suite.servicePublicRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()
		err := usecase.Delete(context.TODO(), mockStruct.GeneralInformation.ID)

		// assertions
		assert.NoError(t, err)
		suite.servicePublicRepo.AssertExpectations(t)
		suite.servicePublicRepo.AssertCalled(t, "GetByID", mock.Anything, mock.AnythingOfType("int64"))
		suite.servicePublicRepo.AssertCalled(t, "Delete", mock.Anything, mock.AnythingOfType("int64"))
	})

	t.Run("error-occurred", func(t *testing.T) {
		// mock expectation being called
		suite.servicePublicRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, domain.ErrInternalServerError)
		suite.servicePublicRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(domain.ErrInternalServerError).Once()
		err := usecase.Delete(context.TODO(), mockStruct.GeneralInformation.ID)

		// assertions
		assert.Error(t, err)
	})
}

func TestStore(t *testing.T) {
	var mockStorePublicService domain.StorePublicService
	err = faker.FakeData(&mockStorePublicService)
	db, sqlMock, _ := sqlmock.New()
	sqlMock.ExpectBegin()
	sqlMock.ExpectCommit()
	txMock, _ := db.Begin()
	t.Run("success", func(t *testing.T) {
		// mock expectation being called
		suite.genInfoRepoMock.On("GetTx", mock.Anything).Return(txMock, nil).Once()
		suite.genInfoRepoMock.On("Store", mock.Anything, mock.Anything, txMock).Return(mockStruct.GeneralInformation.ID, nil).Once()
		suite.servicePublicRepo.On("Store", mock.Anything, mock.Anything, txMock).Return(nil).Once()

		err := usecase.Store(context.TODO(), mockStorePublicService)

		// assertions
		assert.NoError(t, err)
		suite.servicePublicRepo.AssertExpectations(t)
		suite.genInfoRepoMock.AssertExpectations(t)
		suite.genInfoRepoMock.AssertCalled(t, "GetTx", mock.Anything)
		suite.genInfoRepoMock.AssertCalled(t, "Store", mock.Anything, mock.Anything, mock.Anything)
		suite.servicePublicRepo.AssertCalled(t, "Store", mock.Anything, mock.Anything, mock.Anything)
	})

	t.Run("error-occurred", func(t *testing.T) {
		// mock expectation being called
		suite.genInfoRepoMock.On("GetTx", mock.Anything).Return(nil, domain.ErrInternalServerError).Once()
		suite.genInfoRepoMock.On("Store", mock.Anything, mock.Anything, txMock).Return(nil, domain.ErrInternalServerError).Once()
		suite.servicePublicRepo.On("Store", mock.Anything, mock.Anything, txMock).Return(domain.ErrInternalServerError).Once()

		err := usecase.Store(context.TODO(), mockStorePublicService)

		// assertions
		assert.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	var mockUpdatePublicService domain.UpdatePublicService
	err = faker.FakeData(&mockUpdatePublicService)
	db, sqlMock, _ := sqlmock.New()
	sqlMock.ExpectBegin()
	sqlMock.ExpectCommit()
	txMock, _ := db.Begin()
	t.Run("success", func(t *testing.T) {
		// mock expectation being called
		suite.genInfoRepoMock.On("GetTx", mock.Anything).Return(txMock, nil)
		suite.servicePublicRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockStruct, nil)
		suite.genInfoRepoMock.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("int64"), txMock).Return(nil)
		suite.servicePublicRepo.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("int64"), txMock).Return(nil)

		err := usecase.Update(context.TODO(), mockUpdatePublicService, mockStruct.ID)

		// assertions
		assert.NoError(t, err)
		suite.genInfoRepoMock.AssertCalled(t, "GetTx", mock.Anything)
		suite.servicePublicRepo.AssertCalled(t, "GetByID", mock.Anything, mock.AnythingOfType("int64"))
		suite.genInfoRepoMock.AssertCalled(t, "Update", mock.Anything, mock.Anything, mock.AnythingOfType("int64"), mock.Anything)
		suite.servicePublicRepo.AssertCalled(t, "Update", mock.Anything, mock.Anything, mock.AnythingOfType("int64"), mock.Anything)
	})

	t.Run("error-occurred", func(t *testing.T) {
		// mock expectation being called
		suite.genInfoRepoMock.On("GetTx", mock.Anything).Return(nil, domain.ErrInternalServerError)
		suite.servicePublicRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, domain.ErrInternalServerError)
		suite.genInfoRepoMock.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("int64"), txMock).Return(domain.ErrInternalServerError)
		suite.servicePublicRepo.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("int64"), txMock).Return(domain.ErrInternalServerError)

		err := usecase.Update(context.TODO(), mockUpdatePublicService, mockStruct.ID)

		// assertions
		assert.Error(t, err)
	})
}
