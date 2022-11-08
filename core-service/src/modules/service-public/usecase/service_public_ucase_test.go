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
