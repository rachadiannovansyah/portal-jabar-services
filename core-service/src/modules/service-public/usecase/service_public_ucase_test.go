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
)

type servicePublicUsecaseTestSuite struct {
	spRepoMock           mocks.ServicePublicRepository
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
		spRepoMock:           spRepoMock,
		userRepoMock:         userRepoMock,
		genInfoRepoMock:      genInfoRepoMock,
		searchRepoMock:       searchRepoMock,
		cfg:                  cfg,
		ctxTimeout:           ctxTimeout,
		servicePublicUsecase: ucase.NewServicePublicUsecase(&spRepoMock, &genInfoRepoMock, &userRepoMock, &searchRepoMock, cfg, ctxTimeout),
	}
}

func TestFetch(t *testing.T) {
	suite := testSuite()

	var mockStruct domain.ServicePublic
	err := faker.FakeData(&mockStruct)
	assert.NoError(t, err)
	mockList := make([]domain.ServicePublic, 0)
	mockList = append(mockList, mockStruct)
	u := ucase.NewServicePublicUsecase(&suite.spRepoMock, &suite.genInfoRepoMock, &suite.userRepoMock, &suite.searchRepoMock, suite.cfg, suite.ctxTimeout)

	params := &domain.Request{
		Keyword:   "",
		PerPage:   10,
		Offset:    0,
		SortBy:    "",
		SortOrder: "",
	}

	t.Run("success", func(t *testing.T) {
		suite.spRepoMock.On("Fetch", mock.Anything, mock.Anything).Return(mockList, nil).Once()
		list, err := u.Fetch(context.TODO(), params)

		assert.NoError(t, err)
		assert.Len(t, list, len(mockList))
		suite.spRepoMock.AssertExpectations(t)
		suite.spRepoMock.AssertCalled(t, "Fetch", mock.Anything, mock.Anything)
	})

	t.Run("error-occurred", func(t *testing.T) {
		suite.spRepoMock.On("Fetch", mock.Anything, mock.Anything).Return(nil, domain.ErrInternalServerError).Once()

		list, err := u.Fetch(context.TODO(), params)

		assert.Error(t, err)
		assert.Len(t, list, 0)
		suite.spRepoMock.AssertExpectations(t)
		suite.spRepoMock.AssertCalled(t, "Fetch", mock.Anything, mock.Anything)
	})
}

func TestGetBySlug(t *testing.T) {
	suite := testSuite()

	var mockStruct domain.ServicePublic
	err := faker.FakeData(&mockStruct)
	assert.NoError(t, err)
	u := ucase.NewServicePublicUsecase(&suite.spRepoMock, &suite.genInfoRepoMock, &suite.userRepoMock, &suite.searchRepoMock, suite.cfg, suite.ctxTimeout)

	t.Run("success", func(t *testing.T) {
		suite.spRepoMock.On("GetBySlug", mock.Anything, mock.AnythingOfType("string")).Return(mockStruct, nil).Once()
		a, err := u.GetBySlug(context.TODO(), mockStruct.GeneralInformation.Slug)

		assert.NoError(t, err)
		assert.NotNil(t, a)
		suite.spRepoMock.AssertExpectations(t)
		suite.spRepoMock.AssertCalled(t, "GetBySlug", mock.Anything, mock.AnythingOfType("string"))
	})

	t.Run("error-occurred", func(t *testing.T) {
		suite.spRepoMock.On("GetBySlug", mock.Anything, mock.AnythingOfType("string")).Return(domain.ServicePublic{}, domain.ErrInternalServerError).Once()
		a, err := u.GetBySlug(context.TODO(), mockStruct.GeneralInformation.Slug)

		assert.Error(t, err)
		assert.Equal(t, domain.ServicePublic{}, a)
		suite.spRepoMock.AssertExpectations(t)
		suite.spRepoMock.AssertCalled(t, "GetBySlug", mock.Anything, mock.AnythingOfType("string"))
	})
}

func TestDelete(t *testing.T) {
	suite := testSuite()

	var mockStruct domain.ServicePublic
	err := faker.FakeData(&mockStruct)
	assert.NoError(t, err)
	u := ucase.NewServicePublicUsecase(&suite.spRepoMock, &suite.genInfoRepoMock, &suite.userRepoMock, &suite.searchRepoMock, suite.cfg, suite.ctxTimeout)

	t.Run("success", func(t *testing.T) {
		suite.spRepoMock.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockStruct, nil)
		suite.spRepoMock.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()
		err := u.Delete(context.TODO(), mockStruct.GeneralInformation.ID)

		assert.NoError(t, err)
		suite.spRepoMock.AssertExpectations(t)
		suite.spRepoMock.AssertCalled(t, "GetByID", mock.Anything, mock.AnythingOfType("int64"))
		suite.spRepoMock.AssertCalled(t, "Delete", mock.Anything, mock.AnythingOfType("int64"))
	})

	t.Run("error-occurred", func(t *testing.T) {
		suite.spRepoMock.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, domain.ErrInternalServerError)
		suite.spRepoMock.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(domain.ErrInternalServerError).Once()
		err := u.Delete(context.TODO(), mockStruct.GeneralInformation.ID)

		assert.Error(t, err)
		suite.spRepoMock.AssertExpectations(t)
		suite.spRepoMock.AssertCalled(t, "Delete", mock.Anything, mock.AnythingOfType("int64"))
		suite.spRepoMock.AssertCalled(t, "Delete", mock.Anything, mock.AnythingOfType("int64"))
	})
}
