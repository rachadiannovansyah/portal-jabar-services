package http

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
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
