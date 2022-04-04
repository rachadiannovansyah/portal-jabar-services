package http

import (
	"net/http"

	middl "github.com/jabardigitalservice/portal-jabar-services/core-service/src/middleware"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/labstack/echo/v4"
)

type RegistrationInvitationHandler struct {
	IUsecase domain.RegistrationInvitationUsecase
}

func NewRegistrationInvitationHandler(e *echo.Group, r *echo.Group, us domain.RegistrationInvitationUsecase) {
	handler := &RegistrationInvitationHandler{
		IUsecase: us,
	}
	r.POST("/registration-invitations", handler.Invite, middl.CheckPermission(domain.PermissionInviteUser))
	e.POST("/registration-invitations/authorize", handler.Authorize)
}

func (r *RegistrationInvitationHandler) Invite(c echo.Context) error {
	ctx := c.Request().Context()
	au := helpers.GetAuthenticatedUser(c)

	var regInvitation domain.RegistrationInvitation
	err := c.Bind(&regInvitation)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.NewErrResponse(err))
	}

	regInvitation.InvitedBy = au.ID
	regInvitation.UnitID = au.Unit.ID

	res, err := r.IUsecase.Invite(ctx, regInvitation)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.NewErrResponse(err))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "invitation sent",
		"token":   res.Token, // FIXME: should be delete this response after implement email service
	})
}

func (r *RegistrationInvitationHandler) Authorize(c echo.Context) error {
	ctx := c.Request().Context()

	var regInvitation domain.RegistrationInvitation
	err := c.Bind(&regInvitation)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.NewErrResponse(err))
	}

	res, err := r.IUsecase.Authorize(ctx, regInvitation.Token)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, domain.NewErrResponse(err))
	}

	return c.JSON(http.StatusOK, res)
}
