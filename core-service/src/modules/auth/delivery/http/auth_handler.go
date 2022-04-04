package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"

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

	e.POST("/auth/login", handler.Login)
	e.POST("/auth/refresh", handler.RefreshToken)
	r.GET("/auth/permissions", handler.GetPermissions)
}

func isRequestValid(f *domain.LoginRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(f)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Login ...
func (h *AuthHandler) Login(c echo.Context) error {
	cred := new(domain.LoginRequest)
	if err := c.Bind(cred); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(cred); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()

	res, err := h.AUsecase.Login(ctx, cred)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &res)
}

func (h *AuthHandler) RefreshToken(c echo.Context) error {
	refreshRequest := new(domain.RefreshRequest)
	if err := c.Bind(refreshRequest); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()

	res, err := h.AUsecase.RefreshToken(ctx, refreshRequest)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *AuthHandler) GetPermissions(c echo.Context) error {
	ctx := c.Request().Context()
	au := helpers.GetAuthenticatedUser(c)

	res, err := h.AUsecase.GetPermissionsByRoleID(ctx, au.Role.ID)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"permissions": res})
}
