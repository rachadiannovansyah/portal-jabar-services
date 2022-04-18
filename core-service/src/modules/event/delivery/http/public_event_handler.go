package http

import (
	"net/http"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

// PublicEventHandler is represented by domain.EventUsecase
type PublicEventHandler struct {
	EventUcase domain.EventUsecase
}

// NewPublicEventHandler will initialize the event endpoint
func NewPublicEventHandler(p *echo.Group, us domain.EventUsecase) {
	handler := &PublicEventHandler{
		EventUcase: us,
	}
	p.GET("/events", handler.AgendaPortal)
	p.GET("/events/calendar", handler.AgendaCalendar)
}

// AgendaPortal for fetching data on portal agenda
func (h *PublicEventHandler) AgendaPortal(c echo.Context) error {
	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)

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
func (h *PublicEventHandler) AgendaCalendar(c echo.Context) error {
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
