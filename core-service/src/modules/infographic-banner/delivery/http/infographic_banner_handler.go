package http

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
	"github.com/labstack/echo/v4"
)

type infographicBannerHandler struct {
	IUsecase domain.InfographicBannerUsecase
	apm      *utils.Apm
}

func NewInfographicBannerHandler(r *echo.Group, ucase domain.InfographicBannerUsecase, apm *utils.Apm) {
	handler := &infographicBannerHandler{
		IUsecase: ucase,
		apm:      apm,
	}

	r.POST("/infographic-banners", handler.store)
}

func (h *infographicBannerHandler) store(c echo.Context) (err error) {
	req := new(domain.StoreInfographicBanner)
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
