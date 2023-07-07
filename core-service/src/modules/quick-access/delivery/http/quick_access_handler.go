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

type quickAccessHandler struct {
	IUsecase domain.QuickAccessUsecase
	apm      *utils.Apm
}

func NewQuickAccessHandler(r *echo.Group, ucase domain.QuickAccessUsecase, apm *utils.Apm) {
	handler := &quickAccessHandler{
		IUsecase: ucase,
		apm:      apm,
	}

	r.POST("/quick-accesses", handler.Store)
	r.GET("/quick-accesses", handler.Fetch)
	r.DELETE("/quick-accesses/:id", handler.Delete)
	r.GET("/quick-accesses/:id", handler.GetByID)
	r.PUT("/quick-accesses/:id", handler.Update)
	r.PATCH("/quick-accesses/:id/status", handler.UpdateStatus)
}

func (h *quickAccessHandler) Store(c echo.Context) (err error) {
	req := new(domain.StoreQuickAccess)
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

func (h *quickAccessHandler) Fetch(c echo.Context) (err error) {
	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"is_active": c.QueryParam("is_active"),
	}

	rows, total, err := h.IUsecase.Fetch(ctx, params)
	if err != nil {
		return
	}

	list := make([]domain.QuickAccessResponse, 0)
	for _, row := range rows {
		resp := domain.QuickAccessResponse{
			ID:          row.ID,
			Title:       row.Title,
			Description: row.Description,
			Image:       row.Image,
			Link:        row.Link,
			IsActive:    row.IsActive == 1,
			CreatedAt:   row.CreatedAt,
			UpdatedAt:   row.UpdatedAt,
		}

		list = append(list, resp)
	}

	res := helpers.Paginate(c, list, total, params)

	return c.JSON(http.StatusOK, res)
}

func (h *quickAccessHandler) UpdateStatus(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))

	ID := int64(id)
	ctx := c.Request().Context()

	req := new(domain.UpdateStatusQuickAccess)
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

func (h *quickAccessHandler) Update(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))

	ID := int64(id)
	ctx := c.Request().Context()

	req := new(domain.StoreQuickAccess)
	if err = c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(req); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = h.IUsecase.Update(ctx, ID, req); err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, domain.MessageResponse{
		Message: "Successfully updated.",
	})
}

func (h *quickAccessHandler) Delete(c echo.Context) (err error) {
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

func (h *quickAccessHandler) GetByID(c echo.Context) (err error) {
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

	var res domain.QuickAccessResponse

	copier.Copy(&res, &row)

	res.IsActive = row.IsActive == 1

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
