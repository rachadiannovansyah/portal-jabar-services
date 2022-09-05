package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

// PublicServiceHandler ...
type PublicServiceHandler struct {
	PSUsecase domain.PublicServiceUsecase
}

// NewPublicServiceHandler will create a new PublicServiceHandler
func NewPublicServiceHandler(e *echo.Group, r *echo.Group, ps domain.PublicServiceUsecase) {
	handler := &PublicServiceHandler{
		PSUsecase: ps,
	}
	r.POST("/public-service", handler.Store)
	r.DELETE("/public-service/:id", handler.Delete)
}

func isRequestValid(f *domain.StorePserviceRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(f)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Store will store the public-service by given request body
func (h *PublicServiceHandler) Store(c echo.Context) (err error) {
	// binding a request body to domain.StorePserviceRequest struct
	ps := new(domain.StorePserviceRequest)
	if err = c.Bind(ps); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	// check an requests is valid or not
	var ok bool
	if ok, err = isRequestValid(ps); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// store public-service
	ctx := c.Request().Context()
	err = h.PSUsecase.Store(ctx, ps)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, ps)
}

// Delete will drop the public-service data by given id
func (h *PublicServiceHandler) Delete(c echo.Context) (err error) {
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	// set request
	id := int64(reqID)
	ctx := c.Request().Context()

	// destroy public-service data by given id
	err = h.PSUsecase.Delete(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
