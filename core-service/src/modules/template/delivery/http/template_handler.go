package http

import (
	"net/http"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/labstack/echo/v4"
)

type MailHandler struct {
	MUsecase domain.TemplateUsecase
}

func NewMailHandler(_ *echo.Group, r *echo.Group, us domain.TemplateUsecase) {
	handler := &MailHandler{
		MUsecase: us,
	}

	r.GET("/template/account-submission", handler.GetTemplate)
}

func (h *MailHandler) GetTemplate(c echo.Context) error {
	ctx := c.Request().Context()
	au := helpers.GetAuthenticatedUser(c)

	getMail, err := h.MUsecase.GetByTemplate(ctx, au.ID, "administrator")
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, getMail)
}
