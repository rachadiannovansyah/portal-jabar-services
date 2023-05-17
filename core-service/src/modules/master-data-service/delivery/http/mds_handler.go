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
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/policies"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
)

// MasterDataServiceHandler ...
type MasterDataServiceHandler struct {
	MdsUcase domain.MasterDataServiceUsecase
	apm      *utils.Apm
}

// NewMasterDataServiceHandler will create a new MasterDataServiceHandler
func NewMasterDataServiceHandler(r *echo.Group, sp domain.MasterDataServiceUsecase, apm *utils.Apm) {
	handler := &MasterDataServiceHandler{
		MdsUcase: sp,
		apm:      apm,
	}
	r.POST("/master-data-services", handler.Store)
	r.GET("/master-data-services", handler.Fetch)
	r.DELETE("/master-data-services/:id", handler.Delete)
	r.GET("/master-data-services/:id", handler.GetByID)
	r.PUT("/master-data-services/:id", handler.Update)
	r.GET("/master-data-services/tabs", handler.TabStatus)
	r.GET("/master-data-services/archives", handler.Archive)
}

func (h *MasterDataServiceHandler) Store(c echo.Context) (err error) {
	// get a req context
	ctx := c.Request().Context()

	// bind a request body
	body, err := h.bindRequest(c)
	if err != nil {
		return
	}

	// get claims info
	au := domain.JwtCustomClaims{}
	mapstructure.Decode(c.Get("auth:user"), &au)

	body.CreatedBy.ID = au.ID

	err = h.MdsUcase.Store(ctx, &au, body)
	if err != nil {
		return err
	}

	result := map[string]interface{}{
		"message": "CREATED.",
	}

	return c.JSON(http.StatusCreated, result)
}

func isRequestValid(st interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(st)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (h *MasterDataServiceHandler) bindRequest(c echo.Context) (body *domain.StoreMasterDataService, err error) {
	body = new(domain.StoreMasterDataService)
	if err = c.Bind(body); err != nil {
		return &domain.StoreMasterDataService{}, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(body); !ok {
		return &domain.StoreMasterDataService{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return
}

func (h *MasterDataServiceHandler) Fetch(c echo.Context) error {

	ctx := c.Request().Context()
	au := helpers.GetAuthenticatedUser(c)
	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"status": c.QueryParam("status"),
	}

	data, total, err := h.MdsUcase.Fetch(ctx, au, &params)
	if err != nil {
		return err
	}

	// represent response to the client
	mdsRes := []domain.ListMasterDataResponse{}
	for _, row := range data {
		res := domain.ListMasterDataResponse{
			ID:          row.ID,
			ServiceName: row.MainService.ServiceName,
			OpdName:     row.MainService.OpdName,
			ServiceUser: row.MainService.ServiceUser,
			Technical:   row.MainService.Technical,
			UpdatedAt:   row.UpdatedAt,
			Status:      row.Status,
		}

		mdsRes = append(mdsRes, res)
	}

	res := helpers.Paginate(c, mdsRes, total, params)

	return c.JSON(http.StatusOK, res)
}

// Delete will delete the master-data-service by given id
func (h *MasterDataServiceHandler) Delete(c echo.Context) (err error) {
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(reqID)
	ctx := c.Request().Context()

	err = h.MdsUcase.Delete(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

// GetByID will get master data serviceby given id
func (h *MasterDataServiceHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	au := helpers.GetAuthenticatedUser(c)

	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	res, err := h.MdsUcase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	if !policies.AllowMdsAccess(au, res) {
		return c.JSON(http.StatusForbidden, helpers.ResponseError{Message: domain.ErrForbidden.Error()})
	}

	// represent response to the client
	detailRes := domain.DetailMasterDataServiceResponse{}
	detailRes.Application.Title = res.Application.Title.String // rendered by nullstring value

	copier.Copy(&detailRes, &res)
	// un-marshalling json string to object
	helpers.GetObjectFromString(res.MainService.Links.String, &detailRes.MainService.Links)
	helpers.GetObjectFromString(res.MainService.OperationalTimes.String, &detailRes.MainService.OperationalTimes)
	helpers.GetObjectFromString(res.MainService.Locations.String, &detailRes.MainService.Locations)
	helpers.GetObjectFromString(res.Application.Features.String, &detailRes.Application.Features)
	helpers.GetObjectFromString(res.AdditionalInformation.SocialMedia.String, &detailRes.AdditionalInformation.SocialMedia)
	helpers.GetObjectFromString(res.MainService.Benefits, &detailRes.MainService.Benefits)
	helpers.GetObjectFromString(res.MainService.Facilities, &detailRes.MainService.Facilities)
	helpers.GetObjectFromString(res.MainService.TermsAndConditions, &detailRes.MainService.TermsAndConditions)
	helpers.GetObjectFromString(res.MainService.ServiceProcedures, &detailRes.MainService.ServiceProcedures)

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &detailRes})
}

func (h *MasterDataServiceHandler) Update(c echo.Context) (err error) {
	ctx := c.Request().Context()
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	ID := int64(reqID)
	body, err := h.bindRequest(c)
	if err != nil {
		return
	}

	auth := domain.JwtCustomClaims{}
	mapstructure.Decode(c.Get("auth:user"), &auth)

	err = h.MdsUcase.Update(ctx, body, ID)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	result := map[string]interface{}{
		"message": "UPDATED",
	}

	return c.JSON(http.StatusOK, result)
}

func (h *MasterDataServiceHandler) TabStatus(c echo.Context) (err error) {
	ctx := c.Request().Context()

	tabs, err := h.MdsUcase.TabStatus(ctx)
	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &tabs})
}

func (h *MasterDataServiceHandler) Archive(c echo.Context) error {

	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"status": c.QueryParam("status"),
	}

	data, err := h.MdsUcase.Archive(ctx, &params)
	if err != nil {
		return err
	}

	// represent response to the client
	mdsRes := []domain.ListMasterDataResponse{}
	for _, row := range data {
		res := domain.ListMasterDataResponse{
			ID:          row.ID,
			ServiceName: row.MainService.ServiceName,
			OpdName:     row.MainService.OpdName,
			ServiceUser: row.MainService.ServiceUser,
			Technical:   row.MainService.Technical,
			UpdatedAt:   row.UpdatedAt,
			Status:      row.Status,
		}

		mdsRes = append(mdsRes, res)
	}

	return c.JSON(http.StatusOK, domain.ResultData{Data: mdsRes})
}
