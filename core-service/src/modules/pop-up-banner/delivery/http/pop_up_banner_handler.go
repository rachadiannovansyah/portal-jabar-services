package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
)

// PopUpBannerHandler ...
type PopUpBannerHandler struct {
	PUsecase domain.PopUpBannerUsecase
	apm      *utils.Apm
}

// NewPopUpBannerHandler will create a new PopUpBannerHandler
func NewPopUpBannerHandler(r *echo.Group, ucase domain.PopUpBannerUsecase, apm *utils.Apm) {
	handler := &PopUpBannerHandler{
		PUsecase: ucase,
		apm:      apm,
	}

	r.GET("/pop-up-banners", handler.Fetch)
	r.GET("/pop-up-banners/:id", handler.GetByID)
}

// Fetch will fetch the service-public
func (h *PopUpBannerHandler) Fetch(c echo.Context) error {
	// define requirements of request
	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"status": c.QueryParam("status"),
	}

	auth := helpers.GetAuthenticatedUser(c)

	// usecase needed
	data, total, err := h.PUsecase.Fetch(ctx, auth, &params)
	if err != nil {
		return err
	}

	// re-presenting responses
	listPopUpResponse := []domain.ListPopUpBannerResponse{}
	for _, row := range data {
		// attach object response
		resp := domain.ListPopUpBannerResponse{
			ID:        row.ID,
			Title:     row.Title,
			Link:      row.Link,
			Status:    row.Status,
			Duration:  row.Duration,
			StartDate: row.StartDate,
		}

		// un-marshalling object's string
		helpers.GetObjectFromString(row.Image.String, &resp.Image)

		// append element the end of slice
		listPopUpResponse = append(listPopUpResponse, resp)
	}

	res := helpers.Paginate(c, listPopUpResponse, total, params)

	return c.JSON(http.StatusOK, res)
}

// GetByID will get pop up banner by given id
func (h *PopUpBannerHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	data, err := h.PUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	// re-presenting responses
	res := domain.DetailPopUpBannerResponse{
		ID:          data.ID,
		Title:       data.Title,
		ButtonLabel: data.ButtonLabel,
		Link:        data.Link,
		Status:      data.Status,
		Duration:    data.Duration,
		StartDate:   data.StartDate,
		UpdateAt:    data.UpdatedAt,
	}

	helpers.GetObjectFromString(data.Image.String, &res.Image)

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &res})
}
