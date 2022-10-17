package http

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/go-playground/validator.v9"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

// PublicServiceHandler ...
type PublicServiceHandler struct {
	PSUsecase domain.PublicServiceUsecase
}

// NewPublicServiceHandler will create a new PublicServiceHandler
func NewPublicServiceHandler(e *echo.Group, p *echo.Group, ps domain.PublicServiceUsecase) {
	handler := &PublicServiceHandler{
		PSUsecase: ps,
	}
	p.GET("/public-service", handler.Fetch)
	p.GET("/public-service/slug/:slug", handler.GetBySlug)
	p.POST("/public-service", handler.Store)
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

func isRequestValid(ps *domain.StorePublicService) (bool, error) {
	validate := validator.New()
	err := validate.Struct(ps)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (h *PublicServiceHandler) Store(c echo.Context) (err error) {
	ps := new(domain.StorePublicService)
	if err = c.Bind(ps); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(ps); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// FIXME: authenticated variables must be global variables to be accessible everywhere
	auth := domain.JwtCustomClaims{}
	mapstructure.Decode(c.Get("auth:user"), &auth)

	ctx := c.Request().Context()
	err = h.PSUsecase.Store(ctx, *ps)
	if err != nil {
		return err
	}

	result := map[string]interface{}{
		"message": "CREATED",
	}

	return c.JSON(http.StatusCreated, result)
}
