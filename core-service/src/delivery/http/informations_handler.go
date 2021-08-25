package http

import (
	"math"
	"net/http"
	"strconv"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/labstack/echo/v4"
)

type InformationHandler struct {
	InformationsUcase domain.InformationsUcase
}

func NewInformationHandler(e *echo.Group, r *echo.Group, us domain.InformationsUcase) {
	handler := &InformationHandler{
		InformationsUcase: us,
	}

	e.GET("/informations", handler.FetchAll)
	e.GET("/informations/:id", handler.GetById)
}

func (handler *InformationHandler) FetchAll(c echo.Context) error {

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

	params := domain.FetchInformationsRequest{
		Keyword: c.QueryParam("keyword"),
		PerPage: perPage,
		Offset:  offset,
		OrderBy: c.QueryParam("order_by"),
		SortBy:  c.QueryParam("sort_by"),
	}

	listInformations, total, err := handler.InformationsUcase.FetchAll(ctx, &params)

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

func (handler *InformationHandler) GetById(c echo.Context) error {
	reqId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(reqId)
	ctx := c.Request().Context()

	informations, err := handler.InformationsUcase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &informations})
}
