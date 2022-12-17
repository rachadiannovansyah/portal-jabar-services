package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type mysqlRoleRepository struct {
	Conn *sql.DB
}

// NewMysqlRoleRepository will create an object that represent the news.Repository interface
func NewMysqlRoleRepository(Conn *sql.DB) domain.RoleRepository {
	return &mysqlRoleRepository{Conn}
}

var querySelectRole = `SELECT id, name, description, created_at, updated_at FROM roles WHERE 1=1`

func (m *mysqlRoleRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Role, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	result = make([]domain.Role, 0)
	for rows.Next() {
		r := domain.Role{}
		err = rows.Scan(
			&r.ID,
			&r.Name,
			&r.Description,
			&r.CreatedAt,
			&r.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, r)
	}

	return result, nil
}

func (m *mysqlRoleRepository) findOne(ctx context.Context, key string, value string) (res domain.Role, err error) {
	query := fmt.Sprintf("%s AND %v=?", querySelectRole, key)

	list, err := m.fetch(ctx, query, value)
	if err != nil {
		return domain.Role{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlRoleRepository) GetByID(ctx context.Context, id int8) (res domain.Role, err error) {
	query := querySelectRole + ` AND id = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.Role{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}
