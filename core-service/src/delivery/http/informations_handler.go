package http

import (
	"math"
	"net/http"
	"strconv"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/labstack/echo/v4"
)

// InformationHandler ...
type InformationHandler struct {
	InformationsUcase domain.InformationUsecase
}

// NewInformationHandler ...
func NewInformationHandler(e *echo.Group, r *echo.Group, us domain.InformationUsecase) {
	handler := &InformationHandler{
		InformationsUcase: us,
	}

	e.GET("/informations", handler.Fetch)
	e.GET("/informations/:id", handler.GetByID)
}

// Fetch ...
func (handler *InformationHandler) Fetch(c echo.Context) error {

	ctx := c.Request().Context()

	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	perPage, _ := strconv.ParseInt(c.QueryParam("per_page"), 10, 64)

	if page == 0 {
		page = 1
	}
	if perPage == 0 {
		perPage = 10
	}

	offset := (page - 1) * perPage

	params := domain.Request{
		Keyword: c.QueryParam("keyword"),
		PerPage: perPage,
		Offset:  offset,
		OrderBy: c.QueryParam("order_by"),
		SortBy:  c.QueryParam("sort_by"),
	}

	listInformations, total, err := handler.InformationsUcase.Fetch(ctx, &params)

	if err != nil {
		return c.JSON(getStatusCode(err), &ResponseError{Message: err.Error()})
	}

	res := &domain.ResultsData{
		Data: listInformations,
		Meta: &domain.MetaData{
			TotalCount:  total,
			TotalPage:   math.Ceil(float64(total) / float64(perPage)),
			CurrentPage: page,
			PerPage:     perPage,
		},
	}

	return c.JSON(http.StatusOK, res)
}

// GetByID ...
func (handler *InformationHandler) GetByID(c echo.Context) error {
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(reqID)
	ctx := c.Request().Context()

	informations, err := handler.InformationsUcase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &informations})
}
