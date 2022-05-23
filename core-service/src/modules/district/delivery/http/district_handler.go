package http

import (
	"net/http"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/labstack/echo/v4"
)

// District Handler ...
type DistrictHandler struct {
	DistrictUcase domain.DistrictUsecase
}

// NewDistrictHandler creates a new DistrictHandler
func NewDistrictHandler(e *echo.Group, du domain.DistrictUsecase) {
	handler := &DistrictHandler{
		DistrictUcase: du,
	}

	e.GET("/districts", handler.FetchDistrict)
}

// FetchDistrict returns all districts
func (dh *DistrictHandler) FetchDistrict(c echo.Context) error {
	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)

	list, total, err := dh.DistrictUcase.Fetch(ctx, &params)
	if err != nil {
		return err
	}

	res := helpers.Paginate(c, list, total, params)

	return c.JSON(http.StatusOK, res)
}
