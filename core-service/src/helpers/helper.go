package helpers

import (
	"errors"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

func IsInvitationTokenValid(regInvitation domain.RegistrationInvitation, token string) error {
	if regInvitation.Token != token {
		return errors.New("invalid token")
	}
	return nil
}
