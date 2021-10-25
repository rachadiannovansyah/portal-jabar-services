package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type mysqlUserRepository struct {
	Conn *sql.DB
}

// NewMysqlUserRepository will create an object that represent the news.NewsUserRepository interface
func NewMysqlUserRepository(Conn *sql.DB) domain.UserRepository {
	return &mysqlUserRepository{Conn}
}

var querySelect = `SELECT id, name, username, email, password, unit_id, role_id FROM users WHERE 1=1`

// GetByID ...
func (m *mysqlUserRepository) GetByID(ctx context.Context, id uuid.UUID) (res domain.User, err error) {
	query := querySelect + ` AND id = ?`

	err = m.Conn.QueryRowContext(ctx, query, id).Scan(
		&res.ID,
		&res.Name,
		&res.Username,
		&res.Email,
		&res.Password,
		&res.UnitID,
		&res.RoleID,
	)
	if err != nil {
		logrus.Error("Error found", err)
	}

	return
}

// GetByEmail ...
func (m *mysqlUserRepository) GetByEmail(ctx context.Context, email string) (res domain.User, err error) {
	query := querySelect + fmt.Sprintf(` AND email = %s`, email)

	err = m.scan(ctx, query, res)

	return
}

func (m *mysqlUserRepository) scan(ctx context.Context, query string, res domain.User) (err error) {
	err = m.Conn.QueryRowContext(ctx, query).Scan(
		&res.ID,
		&res.Name,
		&res.Username,
		&res.Email,
		&res.Password,
		&res.UnitID,
		&res.RoleID,
	)
	if err != nil {
		logrus.Error("Error found", err)
	}

	return
}
