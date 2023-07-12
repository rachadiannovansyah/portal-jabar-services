package http

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
	"github.com/labstack/echo/v4"
)

type logoHandler struct {
	IUsecase domain.LogoUsecase
	apm      *utils.Apm
}

func NewLogoHandler(r *echo.Group, ucase domain.LogoUsecase, apm *utils.Apm) {
	handler := &logoHandler{
		IUsecase: ucase,
		apm:      apm,
	}

	r.GET("/logos", handler.Fetch)
	r.POST("/logos", handler.Store)
}

func (h *logoHandler) Fetch(c echo.Context) (err error) {
	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)

	rows, total, err := h.IUsecase.Fetch(ctx, params)
	if err != nil {
		return
	}

	res := helpers.Paginate(c, rows, total, params)

	return c.JSON(http.StatusOK, res)
}

func (h *logoHandler) Store(c echo.Context) (err error) {
	req := new(domain.StoreLogoRequest)
	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(req); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = h.IUsecase.Store(ctx, req)
	if err != nil {
		return err
	}

	res := domain.MessageResponse{
		Message: "successfully stored.",
	}

	return c.JSON(http.StatusCreated, res)
}

func isRequestValid(ps interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(ps)
	if err != nil {
		return false, err
	}
	return true, nil
}
