package http

import (
	"net/http"
	"strconv"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/labstack/echo/v4"
)

// InformationHandler ...
type InformationHandler struct {
	InformationsUcase domain.InformationUsecase
}

// NewInformationHandler ...
func NewInformationHandler(e *echo.Group, r *echo.Group, us domain.InformationUsecase) {
	handler := &InformationHandler{
		InformationsUcase: us,
	}

	e.GET("/informations", handler.Fetch)
	e.GET("/informations/:id", handler.GetByID)
}

// Fetch ...
func (h *InformationHandler) Fetch(c echo.Context) error {

	ctx := c.Request().Context()

	params := GetRequestParams(c)

	listInformations, total, err := h.InformationsUcase.Fetch(ctx, &params)

	if err != nil {
		return err
	}

	res := Paginate(c, listInformations, total, params)

	return c.JSON(http.StatusOK, res)
}

// GetByID ...
func (h *InformationHandler) GetByID(c echo.Context) error {
	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(reqID)
	ctx := c.Request().Context()

	informations, err := h.InformationsUcase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &informations})
}
