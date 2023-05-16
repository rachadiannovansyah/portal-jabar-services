package policies

import (
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

func AllowMdsAccess(au *domain.JwtCustomClaims, mds domain.MasterDataService) bool {
	if au.Role.ID == domain.RoleContributor {
		return au.ID == mds.CreatedBy.ID
	}

	if helpers.IsAdminOPD(au) {
		return au.Unit.Name.String == mds.CreatedBy.UnitName
	}

	if helpers.IsSuperAdmin(au) {
		return true
	}

	return false
}
