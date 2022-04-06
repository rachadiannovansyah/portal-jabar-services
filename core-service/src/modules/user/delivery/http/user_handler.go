package http

import (
	"net/http"

	uuid "github.com/google/uuid"
	middl "github.com/jabardigitalservice/portal-jabar-services/core-service/src/middleware"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
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
	r.PUT("/users/me/account-submission", handler.AccountSubmission, middl.CheckPermission(domain.PermissionRequestToBeAdmin))
	e.POST("/users/register", handler.Register)
	r.GET("/users", handler.UserList, middl.CheckPermission(domain.PermissionManageUser))
	r.GET("/users/:id", handler.GetByID, middl.CheckPermission(domain.PermissionManageUser))
	r.PUT("/users/:id/set-as-admin", handler.SetAsAdmin, middl.CheckPermission(domain.PermissionSetAsAdmin))
	r.PUT("/users/:id/change-email", handler.ChangeEmail, middl.CheckPermission(domain.PermissionChangeEmail))
	r.PUT("/users/:id/activate-account", handler.ActivateAccount, middl.CheckPermission(domain.PermissionActivateAccount))
	e.POST("/users/check-nip-exists", handler.CheckNipExists)
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
	res, err := h.UUsecase.UpdateProfile(ctx, u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.MapUserInfo(res))
}

func (h *UserHandler) UserProfile(c echo.Context) error {
	ctx := c.Request().Context()

	au := helpers.GetAuthenticatedUser(c)

	res, err := h.UUsecase.GetByID(ctx, au.ID)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &domain.ResultsData{Data: helpers.MapUserInfo(res)})
}

func (h *UserHandler) GetByID(c echo.Context) error {
	ctx := c.Request().Context()
	reqId := uuid.MustParse(c.Param("id"))

	member, err := h.UUsecase.GetUserByID(ctx, reqId)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	// Copy slice to slice
	UserDetailRes := domain.UserDetailResponse{}
	copier.Copy(&UserDetailRes, &member)

	return c.JSON(http.StatusOK, &domain.ResultsData{Data: &UserDetailRes})
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

func (h *UserHandler) AccountSubmission(c echo.Context) error {
	ctx := c.Request().Context()
	au := helpers.GetAuthenticatedUser(c)

	_, err := h.UUsecase.AccountSubmission(ctx, au.ID, "administrator")
	if err != nil {
		logrus.Error(err)
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "mail sent.",
	})
}

func (h *UserHandler) Register(c echo.Context) error {
	req := new(domain.User)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(req); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err := h.UUsecase.RegisterByInvitation(ctx, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "register success"})
}

func (h *UserHandler) UserList(c echo.Context) error {
	ctx := c.Request().Context()

	params := helpers.GetRequestParams(c)

	list, total, err := h.UUsecase.UserList(ctx, &params)
	if err != nil {
		return err
	}

	// Copy slice to slice
	UserListRes := []domain.UserListResponse{}
	copier.Copy(&UserListRes, &list)

	res := helpers.Paginate(c, UserListRes, total, params)

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) CheckNipExists(c echo.Context) error {
	req := new(domain.CheckNipExistRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if err := validator.New().Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	res, err := h.UUsecase.CheckIfNipExists(ctx, req.Nip)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"exist": res,
	})
}

func (h *UserHandler) SetAsAdmin(c echo.Context) error {
	req := new(domain.CheckPasswordRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if err := validator.New().Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	roleID := int8(3)
	userID := uuid.MustParse(c.Param("id"))
	ctx := c.Request().Context()
	au := helpers.GetAuthenticatedUser(c)

	err := h.UUsecase.SetAsAdmin(ctx, au.ID, req, userID, roleID)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "role changed"})
}

func (h *UserHandler) ChangeEmail(c echo.Context) error {
	req := new(domain.CheckPasswordRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if err := validator.New().Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	au := helpers.GetAuthenticatedUser(c)
	userID := uuid.MustParse(c.Param("id"))

	err := h.UUsecase.ChangeEmail(ctx, au.ID, req, userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "email changed"})
}

func (h *UserHandler) ActivateAccount(c echo.Context) error {
	req := new(domain.CheckPasswordRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if err := validator.New().Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	au := helpers.GetAuthenticatedUser(c)
	userID := uuid.MustParse(c.Param("id"))

	err := h.UUsecase.ActivateAccount(ctx, au.ID, req, userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "account status changed"})
}
