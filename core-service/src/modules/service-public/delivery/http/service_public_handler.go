package http

import (
	"net/http"
	"regexp"

	"github.com/go-playground/validator"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

// ServicePublicHandler ...
type ServicePublicHandler struct {
	SPUsecase domain.ServicePublicUsecase
}

// NewServicePublicHandler will create a new ServicePublicHandler
func NewServicePublicHandler(e *echo.Group, p *echo.Group, r *echo.Group, sp domain.ServicePublicUsecase) {
	handler := &ServicePublicHandler{
		SPUsecase: sp,
	}
	p.GET("/service-public", handler.Fetch)
	p.GET("/service-public/slug/:slug", handler.GetBySlug)
	r.POST("/service-public", handler.Store)
}

// Fetch will fetch the service-public
func (h *ServicePublicHandler) Fetch(c echo.Context) error {

	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"type":     helpers.RegexReplaceString(c, c.QueryParam("type"), ""),
		"category": helpers.RegexReplaceString(c, c.QueryParam("cat"), ""),
	}

	res, err := h.SPUsecase.Fetch(ctx, &params)
	if err != nil {
		return err
	}

	total, lastUpdated, staticCount, err := h.SPUsecase.MetaFetch(ctx, &params)

	if err != nil {
		return err
	}

	listServicePublic := []domain.ListServicePublicResponse{}
	for _, row := range res {
		// get struct response
		tmp := domain.ListServicePublicResponse{}
		copier.Copy(&tmp, &row.GeneralInformation)
		listServicePublic = append(listServicePublic, tmp)
	}

	data := domain.ServicePublicResult{
		Data: listServicePublic,
		Meta: &domain.CustomMetaData{
			TotalCount:  total,
			LastUpdated: lastUpdated,
			StaticCount: staticCount,
		},
	}

	return c.JSON(http.StatusOK, data)
}

// GetBySlug will get service public by given slug
func (h *ServicePublicHandler) GetBySlug(c echo.Context) error {
	slug := c.Param("slug")
	ctx := c.Request().Context()

	res, err := h.SPUsecase.GetBySlug(ctx, slug)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	// get struct response
	detailRes := domain.DetailServicePublicResponse{
		ID: res.ID,
		GeneralInformation: domain.DetailGeneralInformation{
			ID:          res.GeneralInformation.ID,
			Name:        res.GeneralInformation.Name,
			Alias:       res.GeneralInformation.Alias,
			Description: res.GeneralInformation.Description,
			Slug:        res.GeneralInformation.Slug,
			Category:    res.GeneralInformation.Category,
			Email:       res.GeneralInformation.Email,
			Unit:        res.GeneralInformation.Unit,
			Logo:        res.GeneralInformation.Logo,
			Type:        res.GeneralInformation.Type,
		},
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}

	// un-marshalling json string to object
	helpers.GetObjectFromString(res.Purpose.String, &detailRes.Purpose)
	helpers.GetObjectFromString(res.Facility.String, &detailRes.Facility)
	helpers.GetObjectFromString(res.Requirement.String, &detailRes.Requirement)
	helpers.GetObjectFromString(res.ToS.String, &detailRes.ToS)
	helpers.GetObjectFromString(res.InfoGraphic.String, &detailRes.InfoGraphic)
	helpers.GetObjectFromString(res.FAQ.String, &detailRes.FAQ)
	helpers.GetObjectFromString(res.GeneralInformation.Phone, &detailRes.GeneralInformation.Phone)
	helpers.GetObjectFromString(res.GeneralInformation.Link, &detailRes.GeneralInformation.Link)
	helpers.GetObjectFromString(res.GeneralInformation.Addresses, &detailRes.GeneralInformation.Addresses)
	helpers.GetObjectFromString(res.GeneralInformation.OperationalHours, &detailRes.GeneralInformation.OperationalHours)
	helpers.GetObjectFromString(res.GeneralInformation.Media, &detailRes.GeneralInformation.Media)
	helpers.GetObjectFromString(res.GeneralInformation.SocialMedia, &detailRes.GeneralInformation.SocialMedia)

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &detailRes})
}

func (h *ServicePublicHandler) Store(c echo.Context) (err error) {
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
	err = h.SPUsecase.Store(ctx, *ps)
	if err != nil {
		return err
	}

	result := map[string]interface{}{
		"message": "CREATED",
	}

	return c.JSON(http.StatusCreated, result)
}

func IsISO8601Date(fl validator.FieldLevel) bool {
	ISO8601DateRegexString := "^(?:[1-9]\\d{3}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1\\d|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[1-9]\\d(?:0[48]|[2468][048]|[13579][26])|(?:[2468][048]|[13579][26])00)-02-29)T(?:[01]\\d|2[0-3]):[0-5]\\d:[0-5]\\d(?:\\.\\d{1,9})?(?:Z|[+-][01]\\d:[0-5]\\d)$"
	ISO8601DateRegex := regexp.MustCompile(ISO8601DateRegexString)
	return ISO8601DateRegex.MatchString(fl.Field().String())
}

func isRequestValid(ps *domain.StorePublicService) (bool, error) {
	validate := validator.New()
	_ = validate.RegisterValidation("ISO8601date", IsISO8601Date)
	err := validate.Struct(ps)
	if err != nil {
		return false, err
	}
	return true, nil
}
