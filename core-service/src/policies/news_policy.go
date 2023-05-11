package policies

import (
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

func AllowNewsAccess(au *domain.JwtCustomClaims, n domain.News) bool {
	if au.Role.ID == domain.RoleContributor {
		return au.ID == n.CreatedBy.ID
	}

	if helpers.IsAdminOPD(au) {
		return au.Unit.Name.String == n.CreatedBy.UnitName
	}

	if helpers.IsSuperAdmin(au) {
		return true
	}

	return false
}
