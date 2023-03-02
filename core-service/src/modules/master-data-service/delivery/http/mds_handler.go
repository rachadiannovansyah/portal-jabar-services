package http

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
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
