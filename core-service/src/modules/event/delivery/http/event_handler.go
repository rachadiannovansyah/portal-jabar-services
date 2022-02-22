package http

import (
	"net/http"
	"strconv"

	"github.com/mitchellh/mapstructure"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	validator "gopkg.in/go-playground/validator.v9"
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
	r.POST("/events", handler.Store)
	e.DELETE("/events/:id", handler.Delete)
	e.PUT("/events/:id", handler.Update)
	e.GET("/events/calendar", handler.AgendaCalendar)
	e.GET("/events/portal", handler.AgendaPortal)
}

// Validate domain
func isRequestValid(m *domain.StoreRequestEvent) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Fetch will get events data
func (h *EventHandler) Fetch(c echo.Context) error {

	ctx := c.Request().Context()

	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"type":     c.QueryParam("type"),
		"category": c.QueryParam("cat"),
	}

	listEvent, total, err := h.EventUcase.Fetch(ctx, &params)

	if err != nil {
		return err
	}

	listEventRes := []domain.ListEventResponse{}
	copier.Copy(&listEventRes, &listEvent)

	res := helpers.Paginate(c, listEventRes, total, params)

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
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	detailEventRes := domain.DetailEventResponse{}
	copier.Copy(&detailEventRes, &event)

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &detailEventRes})
}

// Store a new event ..
func (h *EventHandler) Store(c echo.Context) (err error) {
	var events domain.StoreRequestEvent
	err = c.Bind(&events)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&events); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	auth := domain.JwtCustomClaims{}
	mapstructure.Decode(c.Get("auth:user"), &auth)

	events.CreatedBy.ID = auth.ID
	events.CreatedBy.Name = auth.Name

	ctx := c.Request().Context()
	err = h.EventUcase.Store(ctx, &events)

	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, events)
}

// Update an event ..
func (h *EventHandler) Update(c echo.Context) (err error) {
	var events domain.StoreRequestEvent
	err = c.Bind(&events)

	if events.EndHour < events.StartHour {
		return c.JSON(http.StatusUnprocessableEntity, "end hour cannot be earlier than start hour.")
	}

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(reqID)
	ctx := c.Request().Context()

	err = h.EventUcase.Update(ctx, id, &events)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrNotFound.Error())
	}

	return c.JSON(http.StatusOK, events)
}

// Delete an event ..
func (h *EventHandler) Delete(c echo.Context) (err error) {
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(reqID)
	ctx := c.Request().Context()

	err = h.EventUcase.Delete(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

// AgendaPortal for fetching data on portal agenda
func (h *EventHandler) AgendaPortal(c echo.Context) error {
	ctx := c.Request().Context()

	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"isPortal": true, // flagging is portal
	}

	listEvent, total, err := h.EventUcase.AgendaPortal(ctx, &params)

	if err != nil {
		return err
	}

	listEventRes := []domain.ListEventResponse{}
	copier.Copy(&listEventRes, &listEvent)

	res := helpers.Paginate(c, listEventRes, total, params)

	return c.JSON(http.StatusOK, res)
}

// AgendaCalendar for fetching data for calendar
func (h *EventHandler) AgendaCalendar(c echo.Context) error {
	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)

	listEvents, err := h.EventUcase.ListCalendar(ctx, &params)

	if err != nil {
		return nil
	}

	listEventCalendar := []domain.ListEventCalendarReponse{}
	copier.Copy(&listEventCalendar, &listEvents)

	return c.JSON(http.StatusOK, listEventCalendar)
}
