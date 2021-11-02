package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

// FeedbackHandler ...
type FeedbackHandler struct {
	FUsecase domain.FeedbackUsecase
}

// NewFeedbackHandler will create a new FeedbackHandler
func NewFeedbackHandler(e *echo.Group, r *echo.Group, fu domain.FeedbackUsecase) {
	handler := &FeedbackHandler{
		FUsecase: fu,
	}
	e.POST("/feedback", handler.Store)
}

func isRequestValid(f *domain.Feedback) (bool, error) {
	validate := validator.New()
	err := validate.Struct(f)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Store will store the feedback by given request body
func (h *FeedbackHandler) Store(c echo.Context) (err error) {
	// FIXME: Check and verify the recaptcha response token.

	f := new(domain.Feedback)
	if err = c.Bind(f); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(f); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = h.FUsecase.Store(ctx, f)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, f)
}
