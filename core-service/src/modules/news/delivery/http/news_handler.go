package http

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

// NewsHandler ...
type NewsHandler struct {
	CUsecase domain.NewsUsecase
}

// NewNewsHandler will initialize the contents/ resources endpoint
func NewNewsHandler(e *echo.Group, r *echo.Group, us domain.NewsUsecase) {
	handler := &NewsHandler{
		CUsecase: us,
	}
	e.GET("/news", handler.FetchNews)
	e.GET("/news/:id", handler.GetByID)
	e.GET("/news/banner", handler.FetchNewsBanner)
	e.PATCH("/news/:id/share", handler.AddShare)
}

// FetchNews will fetch the content based on given params
func (h *NewsHandler) FetchNews(c echo.Context) error {

	ctx := c.Request().Context()

	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"highlight": c.QueryParam("highlight"),
		"category":  c.QueryParam("cat"),
		"type":      c.QueryParam("type"),
		"tags":      c.QueryParam("tags"),
	}

	listNews, total, err := h.CUsecase.Fetch(ctx, &params)

	if err != nil {
		return err
	}

	// Copy slice to slice
	listNewsRes := []domain.NewsListResponse{}
	copier.Copy(&listNewsRes, &listNews)

	res := helpers.Paginate(c, listNewsRes, total, params)

	return c.JSON(http.StatusOK, res)
}

// FetchNews will fetch the content based on given params
func (h *NewsHandler) FetchNewsBanner(c echo.Context) error {

	ctx := c.Request().Context()

	listNews, err := h.CUsecase.FetchNewsBanner(ctx)

	if err != nil {
		return err
	}

	res := map[string]interface{}{
		"data": listNews,
	}

	return c.JSON(http.StatusOK, res)
}

// FetchNewsHeadline ...
func (h *NewsHandler) FetchNewsHeadline(c echo.Context) error {

	ctx := c.Request().Context()

	headlineNews, err := h.CUsecase.FetchNewsHeadline(ctx)

	if err != nil {
		return err
	}

	// Copy slice to slice
	headlineNewsRes := []domain.NewsBanner{}
	copier.Copy(&headlineNewsRes, &headlineNews)

	res := map[string]interface{}{
		"data": headlineNewsRes,
	}

	return c.JSON(http.StatusOK, res)
}

// GetByID will get article by given id
func (h *NewsHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	news, err := h.CUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	// Copy slice to slice
	newsRes := []domain.DetailNewsResponse{}
	copier.Copy(&newsRes, &news)

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &newsRes})
}

// AddShare counter to share
func (h *NewsHandler) AddShare(c echo.Context) error {
	// FIXME: Check and verify the recaptcha response token.

	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	err = h.CUsecase.AddShare(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully add share count",
	})
}
