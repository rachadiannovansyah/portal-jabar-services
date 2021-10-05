package mysql

import (
	"context"
	"database/sql"

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

// GetByID ...
func (m *mysqlUserRepository) GetByID(ctx context.Context, id uuid.UUID) (res domain.User, err error) {
	query := `SELECT id, name, username, email, password, unit_id, role_id FROM users WHERE id = ?`

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
