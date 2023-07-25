package http

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

// UnitHandler ...
type UnitHandler struct {
	UUsecase domain.UnitUsecase
}

// NewUnitHandler will initialize the contents/ resources endpoint
func NewUnitHandler(e, _ *echo.Group, us domain.UnitUsecase) {
	handler := &UnitHandler{
		UUsecase: us,
	}
	e.GET("/units", handler.FetchUnits)
	e.GET("/units/:id", handler.GetByID)
}

// FetchUnits will fetch the content based on given params
func (h *UnitHandler) FetchUnits(c echo.Context) error {

	ctx := c.Request().Context()

	params := helpers.GetRequestParams(c)

	listUnit, total, err := h.UUsecase.Fetch(ctx, &params)

	if err != nil {
		return err
	}

	// Copy slice to slice
	listUnitRes := []domain.UnitInfo{}
	copier.Copy(&listUnitRes, &listUnit)

	res := helpers.Paginate(c, listUnitRes, total, params)

	return c.JSON(http.StatusOK, res)
}

// GetByID will get units by given id
func (h *UnitHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	unit, err := h.UUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &unit})
}
