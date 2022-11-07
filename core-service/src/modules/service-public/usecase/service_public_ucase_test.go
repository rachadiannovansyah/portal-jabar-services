package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mocks "github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain/mocks/repositories"
	"github.com/stretchr/testify/assert"
)

func TestServicePublicFetch(t *testing.T) {
	cfg := &config.Config{}
	spMocksRepo := &mocks.ServicePublicRepository{}
	spMocksRepo.On("Fetch").
		Return(nil, nil).
		Once()
	uMocksRepo := &mocks.UserRepository{}
	giMocksRepo := &mocks.GeneralInformationRepository{}
	sMocksRepo := &mocks.SearchRepository{}
	ctxTimeout := time.Second * 60

	spUcase := NewServicePublicUsecase(spMocksRepo, giMocksRepo, uMocksRepo, sMocksRepo, cfg, ctxTimeout)

	c := context.Background()
	ctx, _ := context.WithTimeout(c, ctxTimeout)
	params := domain.Request{}
	res, err := spUcase.Fetch(ctx, &params)

	assert.NoError(t, err)
	assert.Len(t, res, 1)
}
