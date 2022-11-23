package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mocksUcase "github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain/mocks/usecases"
	httpDelivery "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/feedback/delivery/http"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStore(t *testing.T) {
	var mockStruct domain.Feedback
	err := faker.FakeData(&mockStruct)
	assert.NoError(t, err)

	body, err := json.Marshal(mockStruct)
	assert.NoError(t, err)

	mockUcase := mocksUcase.FeedbackUsecase{}
	mockUcase.On("Store", mock.Anything, mock.Anything).Return(mockStruct).Once()

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/feedback", strings.NewReader(string(body)))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := &httpDelivery.FeedbackHandler{
		FUsecase: &mockUcase,
	}

	_ = handler.Store(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}
