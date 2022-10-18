package http

import (
	"net/http"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/labstack/echo/v4"
)

type VisitorUsecase struct {
	uc domain.VisitorUsecase
}

func NewCounterVisitorHandler(p *echo.Group, uc domain.VisitorUsecase) {
	handler := &VisitorUsecase{
		uc: uc,
	}

	p.GET("/counter-widget", handler.counterWigdet)
}

func (h *VisitorUsecase) counterWigdet(c echo.Context) (err error) {
	ctx := c.Request().Context()
	path := c.Request().URL.RequestURI()
	visitor := h.uc.GetCounterVisitor(ctx, path)
	return c.JSON(http.StatusOK, visitor)
}
