package http

import (
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/usecases"
	"github.com/labstack/echo/v4"
)

// NewHandler will create a new handler for the given usecase
func NewHandler(e *echo.Group, r *echo.Group, u *usecases.Usecases) {
	NewNewsHandler(e, r, u.NewsUcase)
	NewInformationHandler(e, r, u.InformationUcase)
	NewUnitHandler(e, r, u.UnitUcase)
	NewAgendaHandler(e, r, u.AgendaUcase)
}
