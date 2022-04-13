package policies

import "github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"

func AllowNewsAccess(au *domain.JwtCustomClaims, n domain.News) bool {
	if au.Role.ID == domain.RoleContributor {
		return au.ID == n.CreatedBy.ID
	}

	if au.Role.ID == domain.RoleAdministrator || au.Role.ID == domain.RoleGroupAdmin {
		return au.Unit.Name.String == n.CreatedBy.UnitName
	}

	return false
}
