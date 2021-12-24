package http

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/go-playground/validator.v9"
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

	e.POST("/auth/login", handler.Login)
	e.POST("/auth/refresh", handler.RefreshToken)
	r.GET("/auth/me", handler.UserProfile)
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

func (h *AuthHandler) UserProfile(c echo.Context) error {
	ctx := c.Request().Context()

	// FIXME: authenticated variables must be global variables to be accessible everywhere
	auth := domain.JwtCustomClaims{}
	mapstructure.Decode(c.Get("auth:user"), &auth)

	res, err := h.AUsecase.UserProfile(ctx, auth.ID)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	userinfo := domain.UserInfo{}
	copier.Copy(&userinfo, &res)

	return c.JSON(http.StatusOK, &domain.ResultsData{Data: &userinfo})
}
