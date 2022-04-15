package policies

import (
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

func AllowUserAccess(au *domain.JwtCustomClaims, n domain.User) bool {
	if !helpers.IsSuperAdmin(au) {
		return au.Unit.ID == n.Unit.ID
	}

	return false
}
