package http

import (
	"net/http"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/labstack/echo/v4"
)

type MailHandler struct {
	MUsecase domain.MailUsecase
}

func NewMailHandler(e *echo.Group, us domain.MailUsecase) {
	handler := &MailHandler{
		MUsecase: us,
	}

	e.GET("/mail/template", handler.GetTemplate)
}

func (h *MailHandler) GetTemplate(c echo.Context) error {
	ctx := c.Request().Context()

	getMail, err := h.MUsecase.GetByTemplate(ctx, "administrator")
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, getMail)
}
