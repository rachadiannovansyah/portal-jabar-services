package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
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
	ps := new(domain.StorePserviceRequest)
	if err = c.Bind(ps); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(ps); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = h.PSUsecase.Store(ctx, ps)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, ps)
}
