package http

import (
	"math"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

// UnitHandler ...
type UnitHandler struct {
	UUsecase domain.UnitUsecase
}

// NewUnitHandler will initialize the contents/ resources endpoint
func NewUnitHandler(e *echo.Group, r *echo.Group, us domain.UnitUsecase) {
	handler := &UnitHandler{
		UUsecase: us,
	}
	e.GET("/units", handler.FetchUnits)
	e.GET("/units/:id", handler.GetByID)
}

// FetchUnits will fetch the content based on given params
func (a *UnitHandler) FetchUnits(c echo.Context) error {

	ctx := c.Request().Context()

	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	perPage, _ := strconv.ParseInt(c.QueryParam("per_page"), 10, 64)

	if page == 0 {
		page = 1
	}
	if perPage == 0 {
		perPage = 10
	}

	offset := (page - 1) * perPage

	params := domain.Request{
		Keyword: c.QueryParam("q"),
		PerPage: perPage,
		Offset:  offset,
		OrderBy: c.QueryParam("order_by"),
		SortBy:  c.QueryParam("sort_by"),
	}

	listUnit, total, err := a.UUsecase.Fetch(ctx, &params)

	if err != nil {
		return c.JSON(getStatusCode(err), &ResponseError{Message: err.Error()})
	}

	// Copy slice to slice
	listUnitRes := []domain.UnitListResponse{}
	copier.Copy(&listUnitRes, &listUnit)

	res := &domain.ResultsData{
		Data: listUnitRes,
		Meta: &domain.MetaData{
			TotalCount:  total,
			TotalPage:   math.Ceil(float64(total) / float64(perPage)),
			CurrentPage: page,
			PerPage:     perPage,
		},
	}

	return c.JSON(http.StatusOK, res)
}

// GetByID will get units by given id
func (a *UnitHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	unit, err := a.UUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &unit})
}
