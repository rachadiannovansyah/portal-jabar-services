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
}

func (h *MasterDataServiceHandler) Store(c echo.Context) (err error) {
	// get a req context
	ctx := c.Request().Context()

	// bind a request body
	mds := new(domain.StoreMasterDataService)
	if err = c.Bind(mds); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(mds); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// get claims info
	au := domain.JwtCustomClaims{}
	mapstructure.Decode(c.Get("auth:user"), &au)

	err = h.MdsUcase.Store(ctx, &au, mds)
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

func (h *MasterDataServiceHandler) Fetch(c echo.Context) error {

	ctx := c.Request().Context()
	au := helpers.GetAuthenticatedUser(c)
	params := helpers.GetRequestParams(c)

	data, total, err := h.MdsUcase.Fetch(ctx, au, &params)
	if err != nil {
		return err
	}

	// represent response to the client
	mdsRes := []domain.ListMasterDataResponse{}
	for _, row := range data {
		res := domain.ListMasterDataResponse{
			ID:                row.ID,
			ServiceName:       row.MainService.ServiceName,
			OpdName:           row.MainService.OpdName,
			ServiceUser:       row.MainService.ServiceUser,
			OperationalStatus: row.MainService.OperationalStatus,
			UpdatedAt:         row.UpdatedAt,
			Status:            row.Status,
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
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	res, err := h.MdsUcase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	// represent response to the client
	detailRes := domain.DetailMasterDataServiceResponse{}
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
