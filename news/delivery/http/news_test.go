package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	newsHttp "github.com/jabardigitalservice/portal-jabar-api/news/delivery/http"

	"github.com/jabardigitalservice/portal-jabar-api/domain"
	"github.com/jabardigitalservice/portal-jabar-api/domain/mocks"
)

func TestFetch(t *testing.T) {
	var mockNews domain.News
	err := faker.FakeData(&mockNews)
	assert.NoError(t, err)
	mockUCase := new(mocks.NewsUsecase)
	mockListNews := make([]domain.News, 0)
	mockListNews = append(mockListNews, mockNews)
	num := 1
	cursor := "2"
	mockUCase.On("Fetch", mock.Anything, cursor, int64(num)).Return(mockListNews, "10", nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/news?num=1&cursor="+cursor, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := newsHttp.NewsHandler{
		AUsecase: mockUCase,
	}
	err = handler.FetchNews(c)
	require.NoError(t, err)

	responseCursor := rec.Header().Get("X-Cursor")
	assert.Equal(t, "10", responseCursor)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestFetchError(t *testing.T) {
	mockUCase := new(mocks.NewsUsecase)
	num := 1
	cursor := "2"
	mockUCase.On("Fetch", mock.Anything, cursor, int64(num)).Return(nil, "", domain.ErrInternalServerError)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/news?num=1&cursor="+cursor, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := newsHttp.NewsHandler{
		AUsecase: mockUCase,
	}
	err = handler.FetchNews(c)
	require.NoError(t, err)

	responseCursor := rec.Header().Get("X-Cursor")
	assert.Equal(t, "", responseCursor)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	var mockNews domain.News
	err := faker.FakeData(&mockNews)
	assert.NoError(t, err)

	mockUCase := new(mocks.NewsUsecase)

	num := int(mockNews.ID)

	mockUCase.On("GetByID", mock.Anything, int64(num)).Return(mockNews, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/news/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("news/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := newsHttp.NewsHandler{
		AUsecase: mockUCase,
	}
	err = handler.GetByID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestStore(t *testing.T) {
	mockNews := domain.News{
		Title:     "Title",
		Content:   "News",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tempmockNews := mockNews
	tempmockNews.ID = 0
	mockUCase := new(mocks.NewsUsecase)

	j, err := json.Marshal(tempmockNews)
	assert.NoError(t, err)

	mockUCase.On("Store", mock.Anything, mock.AnythingOfType("*domain.News")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/contents", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/contents")

	handler := newsHttp.NewsHandler{
		AUsecase: mockUCase,
	}
	err = handler.Store(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	var mockNews domain.News
	err := faker.FakeData(&mockNews)
	assert.NoError(t, err)

	mockUCase := new(mocks.NewsUsecase)

	num := int(mockNews.ID)

	mockUCase.On("Delete", mock.Anything, int64(num)).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/news/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("news/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := newsHttp.NewsHandler{
		AUsecase: mockUCase,
	}
	err = handler.Delete(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockUCase.AssertExpectations(t)

}
