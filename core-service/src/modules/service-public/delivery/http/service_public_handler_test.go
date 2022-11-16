package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-faker/faker/v4"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mocksUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain/mocks/usecases"
	servicePublicHttp "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/service-public/delivery/http"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var redisServer *miniredis.Miniredis

func setup() {
	redisServer = mockRedis()
}

func mockRedis() *miniredis.Miniredis {
	s, err := miniredis.Run()

	if err != nil {
		panic(err)
	}

	return s
}

func teardown() {
	redisServer.Close()
}

func TestFetch(t *testing.T) {
	setup()
	var mockStruct domain.ServicePublic
	err := faker.FakeData(&mockStruct)
	assert.NoError(t, err)

	mockUcase := mocksUcase.ServicePublicUsecase{}
	mockList := make([]domain.ServicePublic, 0)
	mockList = append(mockList, mockStruct)

	mockUcase.On("Fetch", mock.Anything, mock.Anything).Return(mockList, nil).Once()
	mockUcase.On("MetaFetch", mock.Anything, mock.Anything).Return(int64(10), "2022-10-01", int64(5), nil).Once()

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/service-public", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := &servicePublicHttp.ServicePublicHandler{
		SPUsecase: &mockUcase,
	}

	err = handler.Fetch(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUcase.AssertExpectations(t)
	teardown()
}
