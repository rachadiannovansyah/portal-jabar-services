package helpers

import (
	"errors"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

func IsInvitationTokenValid(regInvitation domain.RegistrationInvitation, token string) error {
	if regInvitation.Token != token {
		return errors.New("invalid token")
	}

	// token expired after 5 days from updated_at
	if time.Now().Sub(regInvitation.InvitedAt) > (time.Hour * 24 * 5) {
		return errors.New("token expired")
	}

	return nil
}
