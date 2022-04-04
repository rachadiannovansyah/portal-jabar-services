package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type mysqlRolePermissionRepository struct {
	Conn *sql.DB
}

// NewMysqlRoleRepository will create an object that represent the news.Repository interface
func NewMysqlRolePermissionRepository(Conn *sql.DB) domain.RolePermissionRepository {
	return &mysqlRolePermissionRepository{Conn}
}

var querySelectRolePermisson = `
	SELECT p.name FROM role_permissions rp 
	LEFT JOIN permissions p ON rp.permission_id = p.id where 1=1
`

func (m *mysqlRolePermissionRepository) GetPermissionsByRoleID(ctx context.Context, roleID int8) (result []string, err error) {
	var query = fmt.Sprintf("%s", querySelectRolePermisson) // if super admin, return all permissions
	if roleID != 1 {
		query = fmt.Sprintf("%s AND rp.role_id=?", querySelectRolePermisson)
	}

	rows, err := m.Conn.QueryContext(ctx, query, roleID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]string, 0)
	for rows.Next() {
		var permissionName string
		err = rows.Scan(&permissionName)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, permissionName)
	}

	return result, nil
}
