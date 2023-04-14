package http

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
)

// MasterDataPublicationHandler ...
type MasterDataPublicationHandler struct {
	MdpUcase domain.MasterDataPublicationUsecase
	apm      *utils.Apm
}

// NewMasterDataPublicationHandler will create a new MasterDataPublicationHandler
func NewMasterDataPublicationHandler(r *echo.Group, sp domain.MasterDataPublicationUsecase, apm *utils.Apm) {
	handler := &MasterDataPublicationHandler{
		MdpUcase: sp,
		apm:      apm,
	}
	r.POST("/master-data-publications", handler.Store)
	r.GET("/master-data-publications", handler.Fetch)
	r.DELETE("/master-data-publications/:id", handler.Delete)
}

func (h *MasterDataPublicationHandler) Store(c echo.Context) (err error) {
	// get a req context
	ctx := c.Request().Context()

	// bind a request body
	body, err := h.bindRequest(c)
	if err != nil {
		return
	}

	err = h.MdpUcase.Store(ctx, body)
	if err != nil {
		return err
	}

	res := map[string]interface{}{
		"message": "CREATED.",
	}

	return c.JSON(http.StatusCreated, res)
}

func isRequestValid(st interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(st)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (h *MasterDataPublicationHandler) bindRequest(c echo.Context) (body *domain.StoreMasterDataPublication, err error) {
	body = new(domain.StoreMasterDataPublication)
	if err = c.Bind(body); err != nil {
		return &domain.StoreMasterDataPublication{}, echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(body); !ok {
		return &domain.StoreMasterDataPublication{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return
}

func (h *MasterDataPublicationHandler) Fetch(c echo.Context) error {

	ctx := c.Request().Context()
	au := helpers.GetAuthenticatedUser(c)
	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"status": c.QueryParam("status"),
	}

	data, total, err := h.MdpUcase.Fetch(ctx, au, &params)
	if err != nil {
		return err
	}

	// represent responses to the client
	pubRes := []domain.ListMasterDataResponse{}
	for _, row := range data {
		res := domain.ListMasterDataResponse{
			ID:          row.ID,
			ServiceName: row.DefaultInformation.ServiceName,
			OpdName:     row.DefaultInformation.OpdName,
			ServiceUser: row.DefaultInformation.ServiceUser,
			Technical:   row.DefaultInformation.Technical,
			UpdatedAt:   row.UpdatedAt,
			Status:      row.Status,
		}

		pubRes = append(pubRes, res)
	}

	res := helpers.Paginate(c, pubRes, total, params)

	return c.JSON(http.StatusOK, res)
}

// Delete will delete the master-data-publications by given id
func (h *MasterDataPublicationHandler) Delete(c echo.Context) (err error) {
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(reqID)
	ctx := c.Request().Context()

	err = h.MdpUcase.Delete(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
