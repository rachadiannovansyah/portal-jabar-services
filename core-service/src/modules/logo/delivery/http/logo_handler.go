package http

import (
	"net/http"

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
