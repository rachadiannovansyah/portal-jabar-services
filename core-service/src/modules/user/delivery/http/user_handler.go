package http

import (
	"net/http"

	"github.com/jinzhu/copier"
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
	r.GET("/users/me", handler.UserProfile)
	r.PUT("/users/me", handler.UpdateProfile)
	r.PUT("/users/me/change-password", handler.ChangePassword)
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

func (h *UserHandler) UserProfile(c echo.Context) error {
	ctx := c.Request().Context()

	au := helpers.GetAuthenticatedUser(c)

	res, err := h.UUsecase.GetByID(ctx, au.ID)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	userinfo := domain.UserInfo{}
	copier.Copy(&userinfo, &res)

	return c.JSON(http.StatusOK, &domain.ResultsData{Data: &userinfo})
}

func (h *UserHandler) ChangePassword(c echo.Context) error {
	req := new(domain.ChangePasswordRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if err := validator.New().Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	au := helpers.GetAuthenticatedUser(c)

	err := h.UUsecase.ChangePassword(ctx, au.ID, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "password changed"})
}
