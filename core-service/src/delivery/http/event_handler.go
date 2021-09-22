package http

import (
	"net/http"
	"strconv"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

// EventHandler is represented by domain.EventUsecase
type EventHandler struct {
	EventUcase domain.EventUsecase
}

// NewEventHandler will initialize the event endpoint
func NewEventHandler(e *echo.Group, r *echo.Group, us domain.EventUsecase) {
	handler := &EventHandler{
		EventUcase: us,
	}

	e.GET("/events", handler.Fetch)
	e.GET("/events/:id", handler.GetByID)
}

// Fetch will get events data
func (h *EventHandler) Fetch(c echo.Context) error {

	ctx := c.Request().Context()

	params := GetRequestParams(c)

	listEvent, total, err := h.EventUcase.Fetch(ctx, &params)

	if err != nil {
		return c.JSON(getStatusCode(err), &ResponseError{Message: err.Error()})
	}

	listEventRes := []domain.ListEventResponse{}
	copier.Copy(&listEventRes, &listEvent)

	res := Paginate(c, listEventRes, total, params)

	return c.JSON(http.StatusOK, res)
}

// GetByID will get event by given id
func (h *EventHandler) GetByID(c echo.Context) error {
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(reqID)
	ctx := c.Request().Context()

	event, err := h.EventUcase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &event})
}
