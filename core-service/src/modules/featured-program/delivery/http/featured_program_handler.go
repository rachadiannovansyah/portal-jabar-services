package http

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

// FeaturedProgramHandler ...
type FeaturedProgramHandler struct {
	FPUsecase domain.FeaturedProgramUsecase
}

// NewFeaturedProgramHandler will initialize the featured-program resources endpoint
func NewFeaturedProgramHandler(e *echo.Group, r *echo.Group, us domain.FeaturedProgramUsecase) {
	handler := &FeaturedProgramHandler{FPUsecase: us}
	e.GET("/featured-programs", handler.FetchFeaturedPrograms)
}

// FetchFeaturedPrograms will fetch the featured-programs
func (h *FeaturedProgramHandler) FetchFeaturedPrograms(c echo.Context) error {

	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"categories": c.Request().URL.Query()["cat[]"],
	}

	featuredProgramsList, total, lastUpdated, err := h.FPUsecase.Fetch(ctx, &params)

	if err != nil {
		return err
	}

	data := domain.ResultsData{
		Data: featuredProgramsList,
		Meta: &domain.CustomMetaData{
			TotalCount:  total,
			LastUpdated: lastUpdated,
		},
	}

	return c.JSON(http.StatusOK, data)
}
