package http

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
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
}

// FetchNews will fetch the content based on given params
func (h *NewsHandler) FetchNews(c echo.Context) error {

	ctx := c.Request().Context()

	params := GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"highlight":   c.QueryParam("highlight"),
		"category_id": c.QueryParam("category_id"),
		"type":        c.QueryParam("type"),
	}

	listNews, total, err := h.CUsecase.Fetch(ctx, &params)

	if err != nil {
		return err
	}

	// Copy slice to slice
	listNewsRes := []domain.NewsListResponse{}
	copier.Copy(&listNewsRes, &listNews)

	res := Paginate(c, listNewsRes, total, params)

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
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &news})
}

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// helper response
func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
