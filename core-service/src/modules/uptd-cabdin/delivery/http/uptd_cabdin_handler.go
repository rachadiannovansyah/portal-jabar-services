package http

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
)

// UptdCabdinHandler ...
type UptdCabdinHandler struct {
	CabdinRepo domain.UptdCabdinUsecase
	apm        *utils.Apm
}

// NewUptdCabdinHandler will create a new UptdCabdinHandler
func NewUptdCabdinHandler(r *echo.Group, ucase domain.UptdCabdinUsecase, apm *utils.Apm) {
	handler := &UptdCabdinHandler{
		CabdinRepo: ucase,
		apm:        apm,
	}

	r.GET("/uptd_cabdins", handler.Fetch)
}

// Fetch will fetch the spbe rals
func (h *UptdCabdinHandler) Fetch(c echo.Context) error {
	// define requirements of request
	ctx := c.Request().Context()

	// usecase needed
	data, err := h.CabdinRepo.Fetch(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}
