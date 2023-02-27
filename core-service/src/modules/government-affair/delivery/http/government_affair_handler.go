package http

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
)

// GovernmentAffairHandler ...
type GovernmentAffairHandler struct {
	GAUsecase domain.GovernmentAffairUsecase
	apm       *utils.Apm
}

// NewGovernmentAffairHandler will create a new GovernmentAffairHandler
func NewGovernmentAffairHandler(r *echo.Group, ucase domain.GovernmentAffairUsecase, apm *utils.Apm) {
	handler := &GovernmentAffairHandler{
		GAUsecase: ucase,
		apm:       apm,
	}

	r.GET("/government_affairs", handler.Fetch)
}

// Fetch will fetch the government affair
func (h *GovernmentAffairHandler) Fetch(c echo.Context) error {
	// define requirements of request
	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)

	// usecase needed
	govAffairRes, total, err := h.GAUsecase.Fetch(ctx, &params)
	if err != nil {
		return err
	}

	// represent to clients
	res := helpers.Paginate(c, govAffairRes, total, params)

	return c.JSON(http.StatusOK, res)
}
