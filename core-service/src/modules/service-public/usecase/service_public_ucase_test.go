package usecase_test

import (
	"context"
	"fmt"
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

func TestFetch(t *testing.T) {
	spRepoMock := mocks.ServicePublicRepository{}
	var mockStruct domain.ServicePublic
	err := faker.FakeData(&mockStruct)
	assert.NoError(t, err)

	mockList := make([]domain.ServicePublic, 0)
	mockList = append(mockList, mockStruct)

	cfg := &config.Config{}
	userRepoMock := mocks.UserRepository{}
	genInfoRepoMock := mocks.GeneralInformationRepository{}
	searchRepoMock := mocks.SearchRepository{}
	ctxTimeout := time.Second * 60
	u := ucase.NewServicePublicUsecase(&spRepoMock, &genInfoRepoMock, &userRepoMock, &searchRepoMock, cfg, ctxTimeout)

	params := &domain.Request{
		Keyword:   "",
		PerPage:   10,
		Offset:    0,
		SortBy:    "",
		SortOrder: "",
	}

	t.Run("success", func(t *testing.T) {
		spRepoMock.On("Fetch", mock.Anything, mock.Anything).Return(mockList, nil).Once()
		list, err := u.Fetch(context.TODO(), params)

		assert.NoError(t, err)
		assert.Len(t, list, len(mockList))
		spRepoMock.AssertExpectations(t)
		spRepoMock.AssertCalled(t, "Fetch", mock.Anything, mock.Anything)
	})

	t.Run("error-occurred", func(t *testing.T) {
		spRepoMock.On("Fetch", mock.Anything, mock.Anything).Return(nil, domain.ErrInternalServerError).Once()

		list, err := u.Fetch(context.TODO(), params)

		assert.Error(t, err)
		assert.Len(t, list, 0)
		spRepoMock.AssertExpectations(t)
		spRepoMock.AssertCalled(t, "Fetch", mock.Anything, mock.Anything)
	})
}

func TestGetBySlug(t *testing.T) {
	spRepoMock := mocks.ServicePublicRepository{}
	var mockStruct domain.ServicePublic
	err := faker.FakeData(&mockStruct)
	assert.NoError(t, err)

	cfg := &config.Config{}
	userRepoMock := mocks.UserRepository{}
	genInfoRepoMock := mocks.GeneralInformationRepository{}
	searchRepoMock := mocks.SearchRepository{}
	ctxTimeout := time.Second * 60
	u := ucase.NewServicePublicUsecase(&spRepoMock, &genInfoRepoMock, &userRepoMock, &searchRepoMock, cfg, ctxTimeout)

	t.Run("success", func(t *testing.T) {
		spRepoMock.On("GetBySlug", mock.Anything, mock.AnythingOfType("string")).Return(mockStruct, nil).Once()
		a, err := u.GetBySlug(context.TODO(), mockStruct.GeneralInformation.Slug)

		assert.NoError(t, err)
		assert.NotNil(t, a)
		spRepoMock.AssertExpectations(t)
		spRepoMock.AssertCalled(t, "GetBySlug", mock.Anything, mock.AnythingOfType("string"))
	})

	t.Run("error-occurred", func(t *testing.T) {
		spRepoMock.On("GetBySlug", mock.Anything, mock.AnythingOfType("string")).Return(domain.ServicePublic{}, domain.ErrInternalServerError).Once()
		a, err := u.GetBySlug(context.TODO(), mockStruct.GeneralInformation.Slug)

		assert.Error(t, err)
		assert.Equal(t, domain.ServicePublic{}, a)
		spRepoMock.AssertExpectations(t)
		spRepoMock.AssertCalled(t, "GetBySlug", mock.Anything, mock.AnythingOfType("string"))
	})
}

func TestDelete(t *testing.T) {
	spRepoMock := mocks.ServicePublicRepository{}
	var mockStruct domain.ServicePublic
	err := faker.FakeData(&mockStruct)
	assert.NoError(t, err)

	cfg := &config.Config{}
	userRepoMock := mocks.UserRepository{}
	genInfoRepoMock := mocks.GeneralInformationRepository{}
	searchRepoMock := mocks.SearchRepository{}
	ctxTimeout := time.Second * 60
	u := ucase.NewServicePublicUsecase(&spRepoMock, &genInfoRepoMock, &userRepoMock, &searchRepoMock, cfg, ctxTimeout)

	t.Run("success", func(t *testing.T) {
		spRepoMock.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockStruct, nil)
		spRepoMock.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()
		err := u.Delete(context.TODO(), mockStruct.GeneralInformation.ID)

		assert.NoError(t, err)
		spRepoMock.AssertExpectations(t)
		spRepoMock.AssertCalled(t, "GetByID", mock.Anything, mock.AnythingOfType("int64"))
		spRepoMock.AssertCalled(t, "Delete", mock.Anything, mock.AnythingOfType("int64"))
	})

	t.Run("error-occurred", func(t *testing.T) {
		spRepoMock.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(nil, domain.ErrInternalServerError)
		spRepoMock.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(domain.ErrInternalServerError).Once()
		err := u.Delete(context.TODO(), mockStruct.GeneralInformation.ID)

		assert.Error(t, err)
		spRepoMock.AssertExpectations(t)
		spRepoMock.AssertCalled(t, "Delete", mock.Anything, mock.AnythingOfType("int64"))
		spRepoMock.AssertCalled(t, "Delete", mock.Anything, mock.AnythingOfType("int64"))
	})
}

func TestStore(t *testing.T) {
	spRepoMock := mocks.ServicePublicRepository{}
	var mockStruct domain.StorePublicService
	err := faker.FakeData(&mockStruct)
	assert.NoError(t, err)

	cfg := &config.Config{}
	userRepoMock := mocks.UserRepository{}
	genInfoRepoMock := mocks.GeneralInformationRepository{}
	searchRepoMock := mocks.SearchRepository{}
	ctxTimeout := time.Second * 60
	u := ucase.NewServicePublicUsecase(&spRepoMock, &genInfoRepoMock, &userRepoMock, &searchRepoMock, cfg, ctxTimeout)

	t.Run("success", func(t *testing.T) {
		genInfoRepoMock.On("GetTx", mock.Anything).Return(nil, nil)
		genInfoRepoMock.On("Store", mock.Anything, mock.Anything, mock.Anything).Return(mockStruct.GeneralInformation.ID, nil)
		spRepoMock.On("Store", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		fmt.Println("1")
		err := u.Store(context.TODO(), mockStruct)
		fmt.Println("2")
		assert.NoError(t, err)
		genInfoRepoMock.AssertExpectations(t)
		genInfoRepoMock.AssertCalled(t, "GetTx", mock.Anything)
		genInfoRepoMock.AssertCalled(t, "Store", mock.Anything, mock.Anything, mock.Anything)
		spRepoMock.AssertExpectations(t)
		spRepoMock.AssertCalled(t, "Store", mock.Anything, mock.Anything, mock.Anything)
	})

	// t.Run("error-occurred", func(t *testing.T) {
	// 	genInfoRepoMock.On("Store", mock.Anything, mock.Anything, mock.Anything).Return(mockStruct.GeneralInformation.ID, domain.ErrInternalServerError).Once()

	// 	err := u.Store(context.TODO(), mockStruct)

	// 	assert.Error(t, err)
	// 	genInfoRepoMock.AssertExpectations(t)
	// 	genInfoRepoMock.AssertCalled(t, "Store", mock.Anything, mock.Anything, mock.Anything)
	// })
}
