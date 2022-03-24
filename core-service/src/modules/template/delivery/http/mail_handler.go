package http

import (
	"net/http"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/labstack/echo/v4"
)

type MailHandler struct {
	MUsecase domain.TemplateUsecase
}

func NewMailHandler(e *echo.Group, us domain.TemplateUsecase) {
	handler := &MailHandler{
		MUsecase: us,
	}

	e.GET("/template/mail", handler.GetTemplate)
}

func (h *MailHandler) GetTemplate(c echo.Context) error {
	ctx := c.Request().Context()

	getMail, err := h.MUsecase.GetByTemplate(ctx, "administrator")
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, getMail)
}
