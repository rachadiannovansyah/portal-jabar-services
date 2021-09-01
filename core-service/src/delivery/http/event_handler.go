package http

import (
	"math"
	"net/http"
	"strconv"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/labstack/echo/v4"
)

type EventHandler struct {
	EventUcase domain.EventUsecase
}

func NewEventHandler(e *echo.Group, r *echo.Group, us domain.EventUsecase) {
	handler := &EventHandler{
		EventUcase: us,
	}

	e.GET("/events", handler.Fetch)
	e.GET("/events/:id", handler.GetByID)
}

func (handler *EventHandler) Fetch(c echo.Context) error {

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
		Keyword:   c.QueryParam("q"),
		PerPage:   perPage,
		Offset:    offset,
		OrderBy:   c.QueryParam("order_by"),
		SortBy:    c.QueryParam("sort_by"),
		StartDate: c.QueryParam("start_date"),
		EndDate:   c.QueryParam("end_date"),
	}

	listEvent, total, err := handler.EventUcase.Fetch(ctx, &params)

	if err != nil {
		return c.JSON(getStatusCode(err), &ResponseError{Message: err.Error()})
	}

	res := &domain.ResultsData{
		Data: listEvent,
		Meta: &domain.MetaData{
			TotalCount:  total,
			TotalPage:   math.Ceil(float64(total) / float64(perPage)),
			CurrentPage: page,
			PerPage:     perPage,
		},
	}

	return c.JSON(http.StatusOK, res)
}

func (handler *EventHandler) GetByID(c echo.Context) error {
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(reqID)
	ctx := c.Request().Context()

	event, err := handler.EventUcase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &event})
}
