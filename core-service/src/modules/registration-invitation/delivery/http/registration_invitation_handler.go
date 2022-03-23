package http

import (
	"net/http"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/labstack/echo/v4"
)

type RegistrationInvitationHandler struct {
	IUsecase domain.RegistrationInvitationUsecase
}

func NewRegistrationInvitationHandler(e *echo.Group, r *echo.Group, us domain.RegistrationInvitationUsecase) {
	handler := &RegistrationInvitationHandler{
		IUsecase: us,
	}
	r.POST("/registration-invitations", handler.Invite)
}

func (r *RegistrationInvitationHandler) Invite(c echo.Context) error {
	ctx := c.Request().Context()

	var regInvitation domain.RegistrationInvitation
	err := c.Bind(&regInvitation)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.NewErrResponse(err))
	}

	res, err := r.IUsecase.Invite(ctx, regInvitation.Email)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.NewErrResponse(err))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "invitation sent",
		"token":   res.Token, // FIXME: should be delete this response after implement email service
	})
}
