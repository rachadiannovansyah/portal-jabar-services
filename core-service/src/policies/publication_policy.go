package policies

import (
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

func AllowPublicationAccess(au *domain.JwtCustomClaims, pub domain.MasterDataPublication) bool {
	if au.Role.ID == domain.RoleContributor {
		return au.ID == pub.CreatedBy.ID
	}

	if helpers.IsAdminOPD(au) {
		return au.Unit.Name.String == pub.CreatedBy.UnitName
	}

	if helpers.IsSuperAdmin(au) {
		return true
	}

	return false
}
