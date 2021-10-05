package http

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
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
	featuredProgramsList, err := h.FPUsecase.Fetch(ctx)

	if err != nil {
		return c.JSON(getStatusCode(err), &ResponseError{Message: err.Error()})
	}

	data := map[string]interface{}{"data": featuredProgramsList}

	return c.JSON(http.StatusOK, data)
}
