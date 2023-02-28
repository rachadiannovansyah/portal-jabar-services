package http

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
)

// SpbeRalsHandler ...
type SpbeRalsHandler struct {
	SRalsUcase domain.SpbeRalsUsecase
	apm        *utils.Apm
}

// NewSpbeRalsHandler will create a new SpbeRalsHandler
func NewSpbeRalsHandler(r *echo.Group, ucase domain.SpbeRalsUsecase, apm *utils.Apm) {
	handler := &SpbeRalsHandler{
		SRalsUcase: ucase,
		apm:        apm,
	}

	r.GET("/spbe_rals", handler.Fetch)
}

// Fetch will fetch the spbe rals
func (h *SpbeRalsHandler) Fetch(c echo.Context) error {
	// define requirements of request
	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)

	// usecase needed
	data, total, err := h.SRalsUcase.Fetch(ctx, &params)
	if err != nil {
		return err
	}

	// represent to clients
	res := helpers.Paginate(c, data, total, params)

	return c.JSON(http.StatusOK, res)
}
