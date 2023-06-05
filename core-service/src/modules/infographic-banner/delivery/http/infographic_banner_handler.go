package http

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type infographicBannerHandler struct {
	IUsecase domain.InfographicBannerUsecase
	apm      *utils.Apm
}

func NewInfographicBannerHandler(r *echo.Group, ucase domain.InfographicBannerUsecase, apm *utils.Apm) {
	handler := &infographicBannerHandler{
		IUsecase: ucase,
		apm:      apm,
	}

	r.POST("/infographic-banners", handler.Store)
	r.GET("/infographic-banners", handler.Fetch)
	r.DELETE("/infographic-banners/:id", handler.Delete)
	r.GET("/infographic-banners/:id", handler.GetByID)
	r.PATCH("/infographic-banners/sequences", handler.UpdateSequence)
	r.PATCH("/infographic-banners/:id/status", handler.UpdateStatus)
}

func (h *infographicBannerHandler) Store(c echo.Context) (err error) {
	req := new(domain.StoreInfographicBanner)
	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(req); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = h.IUsecase.Store(ctx, req)
	if err != nil {
		return err
	}

	res := domain.MessageResponse{
		Message: "successfully stored.",
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *infographicBannerHandler) Fetch(c echo.Context) (err error) {
	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"is_active": c.QueryParam("is_active"),
	}

	rows, total, err := h.IUsecase.Fetch(ctx, params)
	if err != nil {
		return
	}

	list := make([]domain.InfographicBannerResponse, 0)
	for _, row := range rows {
		resp := domain.InfographicBannerResponse{
			ID:        row.ID,
			Title:     row.Title,
			Sequence:  row.Sequence,
			Link:      row.Link,
			IsActive:  row.IsActive == 1,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		}

		helpers.GetObjectFromString(row.Image, &resp.Image)
		list = append(list, resp)
	}

	res := helpers.Paginate(c, list, total, params)

	return c.JSON(http.StatusOK, res)
}

func (h *infographicBannerHandler) UpdateStatus(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))

	ID := int64(id)
	ctx := c.Request().Context()

	req := new(domain.UpdateStatusInfographicBanner)
	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(req); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.IUsecase.UpdateStatus(ctx, ID, req); err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, domain.MessageResponse{
		Message: "Successfully updated.",
	})
}

func (h *infographicBannerHandler) UpdateSequence(c echo.Context) (err error) {
	ctx := c.Request().Context()

	req := new(domain.UpdateSequenceInfographicBanner)
	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(req); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.IUsecase.UpdateSequence(ctx, req); err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, domain.MessageResponse{
		Message: "Successfully updated.",
	})
}

func (h *infographicBannerHandler) Delete(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseError{Message: domain.ErrNotFound.Error()})
	}

	ID := int64(id)
	ctx := c.Request().Context()

	if err = h.IUsecase.Delete(ctx, ID); err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, domain.MessageResponse{
		Message: "Successfully deleted.",
	})
}

func (h *infographicBannerHandler) GetByID(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseError{Message: domain.ErrNotFound.Error()})
	}

	ID := int64(id)
	ctx := c.Request().Context()

	row, err := h.IUsecase.GetByID(ctx, ID)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	var res domain.InfographicBannerResponse

	copier.Copy(&res, &row)

	helpers.GetObjectFromString(row.Image, &res.Image)

	metaDesktop, _ := helpers.GetMetaDataImage(res.Image.Desktop) // for desktop
	metaMobile, _ := helpers.GetMetaDataImage(res.Image.Mobile)   // for mobile
	res.ImageMetaData = &domain.ImageMetaDataBanner{
		Desktop: metaDesktop,
		Mobile:  metaMobile,
	}

	return c.JSON(http.StatusOK, domain.ResultData{Data: res})
}

func isRequestValid(ps interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(ps)
	if err != nil {
		return false, err
	}
	return true, nil
}
