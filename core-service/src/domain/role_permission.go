package domain

import "context"

type RolePermissionRepository interface {
	GetPermissionsByRoleID(ctx context.Context, id int8) ([]string, error)
}
