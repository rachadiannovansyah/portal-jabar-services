package http

import (
	"net/http"
	"strconv"

	"github.com/mitchellh/mapstructure"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	middl "github.com/jabardigitalservice/portal-jabar-services/core-service/src/middleware"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/policies"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	validator "gopkg.in/go-playground/validator.v9"
)

// EventHandler is represented by domain.EventUsecase
type EventHandler struct {
	EventUcase domain.EventUsecase
}

// NewEventHandler will initialize the event endpoint
func NewEventHandler(_, r *echo.Group, us domain.EventUsecase) {
	handler := &EventHandler{
		EventUcase: us,
	}
	permManageEvent := domain.PermissionManageEvent
	r.GET("/events", handler.Fetch, middl.CheckPermission(permManageEvent))
	r.GET("/events/:id", handler.GetByID, middl.CheckPermission(permManageEvent))
	r.POST("/events", handler.Store, middl.CheckPermission(permManageEvent))
	r.DELETE("/events/:id", handler.Delete, middl.CheckPermission(permManageEvent))
	r.PUT("/events/:id", handler.Update, middl.CheckPermission(permManageEvent))
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
	au := helpers.GetAuthenticatedUser(c)

	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"type":       c.Request().URL.Query()["type[]"],
		"categories": c.Request().URL.Query()["cat[]"],
	}

	listEvent, total, err := h.EventUcase.Fetch(ctx, au, &params)

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
	au := helpers.GetAuthenticatedUser(c)

	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(reqID)
	ctx := c.Request().Context()

	event, err := h.EventUcase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	if !policies.AllowEventAccess(au, event) {
		return c.JSON(http.StatusForbidden, helpers.ResponseError{Message: domain.ErrForbidden.Error()})
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
