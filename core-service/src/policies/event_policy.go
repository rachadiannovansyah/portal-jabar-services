package policies

import (
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

func AllowEventAccess(au *domain.JwtCustomClaims, e domain.Event) bool {
	if !helpers.IsSuperAdmin(au) {
		return au.Unit.Name.String == e.CreatedBy.UnitName
	}

	return false
}
