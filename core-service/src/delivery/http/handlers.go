package http

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/usecases"
	"github.com/labstack/echo/v4"
)

// NewHandler will create a new handler for the given usecase
func NewHandler(e *echo.Group, r *echo.Group, u *usecases.Usecases) {
	NewNewsHandler(e, r, u.NewsUcase)
	NewInformationHandler(e, r, u.InformationUcase)
	NewUnitHandler(e, r, u.UnitUcase)
	NewEventHandler(e, r, u.EventUcase)
	NewFeedbackHandler(e, r, u.FeedbackUcase)
	NewFeaturedProgramHandler(e, r, u.FeaturedProgramUcase)
}

// ErrorHandler ...
func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	sentry.CaptureException(err)
	c.Logger().Error(report)
	c.JSON(report.Code, report)
}
