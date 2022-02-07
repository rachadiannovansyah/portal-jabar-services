package http

import (
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

// AreaHandler ...
type AreaHandler struct {
	AUsecase domain.AreaUsecase
}

// NewAreaHandler will initialize the contents/ resources endpoint
func NewAreaHandler(e *echo.Group, r *echo.Group, us domain.AreaUsecase) {
	handler := &AreaHandler{
		AUsecase: us,
	}
	e.GET("/areas", handler.FetchAreas)
}

// FetchAreas will fetch the content based on given params
func (h *AreaHandler) FetchAreas(c echo.Context) error {

	ctx := c.Request().Context()

	params := helpers.GetRequestParams(c)

	listArea, total, err := h.AUsecase.Fetch(ctx, &params)

	if err != nil {
		return err
	}

	// Copy slice to slice
	listAreaRes := []domain.AreaListResponse{}
	copier.Copy(&listAreaRes, &listArea)

	res := helpers.Paginate(c, listAreaRes, total, params)

	return c.JSON(http.StatusOK, res)
}
