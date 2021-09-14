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

func isRequestValid(m *domain.Feedback) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Store will store the feedback by given request body
func (a *FeedbackHandler) Store(c echo.Context) (err error) {
	var feedback domain.Feedback
	err = c.Bind(&feedback)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&feedback); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = a.FUsecase.Store(ctx, &feedback)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, feedback)
}
