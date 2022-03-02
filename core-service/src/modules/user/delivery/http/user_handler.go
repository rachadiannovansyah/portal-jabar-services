package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

// userHandler ...
type UserHandler struct {
	UUsecase domain.UserUsecase
}

// NewUserHandler will create a new UserHandler
func NewUserHandler(e *echo.Group, r *echo.Group, uu domain.UserUsecase) {
	handler := &UserHandler{
		UUsecase: uu,
	}
	r.POST("/users", handler.Store)
	r.PUT("/users/profile", handler.UpdateProfile)
}

func isRequestValid(u *domain.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Store will store the user by given request body
func (h *UserHandler) Store(c echo.Context) (err error) {
	u := new(domain.User)
	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	// FIXME: validate request based on AC Rules (waiting)

	ctx := c.Request().Context()
	err = h.UUsecase.Store(ctx, u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}

func (h *UserHandler) UpdateProfile(c echo.Context) (err error) {
	u := new(domain.User)
	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(u); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	au := helpers.GetAuthenticatedUser(c)

	u.ID = au.ID
	err = h.UUsecase.UpdateProfile(ctx, u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}
