package http

import (
	"math"
	"net/http"
	"strconv"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/labstack/echo/v4"
)

type AgendaHandler struct {
	AgendaUcase domain.AgendaUsecase
}

func NewAgendaHandler(e *echo.Group, r *echo.Group, us domain.AgendaUsecase) {
	handler := &AgendaHandler{
		AgendaUcase: us,
	}

	e.GET("/agenda", handler.Fetch)
	e.GET("/agenda/:id", handler.GetByID)
}

func (handler *AgendaHandler) Fetch(c echo.Context) error {

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
		Keyword:   c.QueryParam("keyword"),
		PerPage:   perPage,
		Offset:    offset,
		OrderBy:   c.QueryParam("order_by"),
		SortBy:    c.QueryParam("sort_by"),
		StartDate: c.QueryParam("start_Date"),
		EndDate:   c.QueryParam("end_date"),
	}

	listAgenda, total, err := handler.AgendaUcase.Fetch(ctx, &params)

	if err != nil {
		return c.JSON(getStatusCode(err), &ResponseError{Message: err.Error()})
	}

	res := &domain.ResultsData{
		Data: listAgenda,
		Meta: &domain.MetaData{
			TotalCount:  total,
			TotalPage:   math.Ceil(float64(total) / float64(perPage)),
			CurrentPage: page,
			PerPage:     perPage,
		},
	}

	return c.JSON(http.StatusOK, res)
}

func (handler *AgendaHandler) GetByID(c echo.Context) error {
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(reqID)
	ctx := c.Request().Context()

	agenda, err := handler.AgendaUcase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &agenda})
}
