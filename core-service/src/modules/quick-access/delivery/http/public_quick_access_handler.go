package http

import (
	"net/http"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
	"github.com/labstack/echo/v4"
)

type publicQuickAccessHandler struct {
	IUsecase domain.QuickAccessUsecase
	apm      *utils.Apm
}

func NewPublicQuickAccessHandler(p *echo.Group, ucase domain.QuickAccessUsecase, apm *utils.Apm) {
	handler := &publicQuickAccessHandler{
		IUsecase: ucase,
		apm:      apm,
	}

	p.GET("/quick-accesses", handler.Fetch)
}

func (h *publicQuickAccessHandler) Fetch(c echo.Context) (err error) {
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

		helpers.GetObjectFromString(row.Image, &resp.Image)
		list = append(list, resp)
	}

	res := helpers.Paginate(c, list, total, params)

	return c.JSON(http.StatusOK, res)
}
