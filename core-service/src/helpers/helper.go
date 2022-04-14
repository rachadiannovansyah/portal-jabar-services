package helpers

import (
	"errors"
	"net/mail"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

func IsInvitationTokenValid(regInvitation domain.RegistrationInvitation, token string) error {
	if regInvitation.Token != token {
		return errors.New("invalid token")
	}
	return nil
}

func IsValidMailAddress(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}

func IsAdminOPD(au *domain.JwtCustomClaims) bool {
	return au.Role.ID == domain.RoleAdministrator || au.Role.ID == domain.RoleGroupAdmin
}
