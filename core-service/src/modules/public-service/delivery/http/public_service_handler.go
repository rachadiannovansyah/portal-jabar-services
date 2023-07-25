package http

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

// PublicServiceHandler ...
type PublicServiceHandler struct {
	PSUsecase domain.PublicServiceUsecase
}

// NewPublicServiceHandler will create a new PublicServiceHandler
func NewPublicServiceHandler(_ *echo.Group, p *echo.Group, ps domain.PublicServiceUsecase) {
	handler := &PublicServiceHandler{
		PSUsecase: ps,
	}
	p.GET("/public-service", handler.Fetch)
	p.GET("/public-service/slug/:slug", handler.GetBySlug)
}

// Fetch will fetch the public-service
func (h *PublicServiceHandler) Fetch(c echo.Context) error {

	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"service_type": c.QueryParam("type"),
		"category":     c.QueryParam("cat"),
	}

	publicServiceList, err := h.PSUsecase.Fetch(ctx, &params)

	// Copy slice to slice
	listPSRes := []domain.ListPublicServiceResponse{}
	copier.Copy(&listPSRes, &publicServiceList)

	if err != nil {
		return err
	}

	total, lastUpdated, err := h.PSUsecase.MetaFetch(ctx, &params)

	if err != nil {
		return err
	}

	// handled if no rows in result then res will provide empty arr
	if total == 0 {
		data := domain.ResultsData{
			Data: []string{},
		}

		return c.JSON(http.StatusOK, data)
	}

	data := domain.ResultsData{
		Data: listPSRes,
		Meta: &domain.CustomMetaData{
			TotalCount:  total,
			LastUpdated: lastUpdated,
		},
	}

	return c.JSON(http.StatusOK, data)
}

// GetBySlug will get public service by given slug
func (h *PublicServiceHandler) GetBySlug(c echo.Context) error {
	slug := c.Param("slug")
	ctx := c.Request().Context()

	news, err := h.PSUsecase.GetBySlug(ctx, slug)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	// get struct response
	detailRes := domain.DetailPublicServiceResponse{}

	// get arr of facilites then unmarshall it
	facilities := make([]domain.Facility, len(news.Facilities))
	json.Unmarshal([]byte(news.Facilities), &facilities)

	// Copy slice to slice
	copier.Copy(&detailRes, &news)
	detailRes.Facilities = facilities

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &detailRes})
}
