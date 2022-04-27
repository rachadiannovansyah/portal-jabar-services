package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

// AwardHandler ...
type AwardHandler struct {
	UUsecase domain.AwardUsecase
}

// NewAwardHandler will initialize the contents/ resources endpoint
func NewAwardHandler(e *echo.Group, us domain.AwardUsecase) {
	handler := &AwardHandler{
		UUsecase: us,
	}
	e.GET("/awards", handler.FetchAwards)
	e.GET("/awards/:id", handler.GetByID)
}

// FetchAwards will fetch the content based on given params
func (h *AwardHandler) FetchAwards(c echo.Context) error {

	ctx := c.Request().Context()

	params := helpers.GetRequestParams(c)

	listAward, total, err := h.UUsecase.Fetch(ctx, &params)

	if err != nil {
		return err
	}

	res := helpers.Paginate(c, listAward, total, params)

	return c.JSON(http.StatusOK, res)
}

// GetByID will get awards by given id
func (h *AwardHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	award, err := h.UUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &award})
}
