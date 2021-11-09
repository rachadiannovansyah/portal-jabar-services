package http

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

// AuthHandler ...
type AuthHandler struct {
	AUsecase domain.AuthUsecase
}

// NewAuthHandler will initialize the contents/ resources endpoint
func NewAuthHandler(e *echo.Group, r *echo.Group, us domain.AuthUsecase) {
	handler := &AuthHandler{
		AUsecase: us,
	}
	e.POST("/login", handler.Login)
}

// Login ...
func (h *AuthHandler) Login(c echo.Context) error {
	cred := new(domain.LoginRequest)
	if err := c.Bind(cred); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()

	res, err := h.AUsecase.Login(ctx, cred)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
