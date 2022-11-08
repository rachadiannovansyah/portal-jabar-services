package http

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/middleware"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
)

// ServicePublicHandler ...
type ServicePublicHandler struct {
	SPUsecase domain.ServicePublicUsecase
	apm       *utils.Apm
}

// NewServicePublicHandler will create a new ServicePublicHandler
func NewServicePublicHandler(e *echo.Group, p *echo.Group, r *echo.Group, sp domain.ServicePublicUsecase, apm *utils.Apm) {
	handler := &ServicePublicHandler{
		SPUsecase: sp,
		apm:       apm,
	}
	p.GET("/service-public", handler.Fetch, middleware.VerifyCache())
	p.GET("/service-public/slug/:slug", handler.GetBySlug, middleware.VerifyCache())
	r.POST("/service-public", handler.Store)
	r.PUT("/service-public/:id", handler.Update)
	r.DELETE("/service-public/:id", handler.Delete)
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

	data := domain.ResultsData{
		Data: listServicePublic,
		Meta: &domain.CustomMetaData{
			TotalCount:  total,
			LastUpdated: lastUpdated,
			StaticCount: staticCount,
		},
	}

	// set cache from dependency injection redis
	helpers.Cache(c.Request().URL.RequestURI(), data.Data, data.Meta)

	return c.JSON(http.StatusOK, data)
}

// GetBySlug will get service public by given slug
func (h *ServicePublicHandler) GetBySlug(c echo.Context) error {
	ctx := c.Request().Context()
	slug := c.Param("slug")

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

	// set cache from dependency injection redis
	helpers.Cache(c.Request().URL.RequestURI(), &detailRes, nil)

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &detailRes})
}

func (h *ServicePublicHandler) Store(c echo.Context) (err error) {
	ctx := c.Request().Context()
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

	err = h.SPUsecase.Store(ctx, *ps)
	if err != nil {
		return err
	}

	result := map[string]interface{}{
		"message": "CREATED",
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *ServicePublicHandler) Update(c echo.Context) (err error) {
	ctx := c.Request().Context()
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	ID := int64(idP)
	ps := new(domain.UpdatePublicService)
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

	err = h.SPUsecase.Update(ctx, *ps, ID)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	result := map[string]interface{}{
		"message": "UPDATED",
	}

	return c.JSON(http.StatusOK, result)
}

func (h *ServicePublicHandler) Delete(c echo.Context) (err error) {
	ctx := c.Request().Context()
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	ID := int64(idP)

	err = h.SPUsecase.Delete(ctx, ID)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	result := map[string]interface{}{
		"message": "DELETED",
	}

	return c.JSON(http.StatusOK, result)
}

func isRequestValid(ps interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(ps)
	if err != nil {
		return false, err
	}
	return true, nil
}
