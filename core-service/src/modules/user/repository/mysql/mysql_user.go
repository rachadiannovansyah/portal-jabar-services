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

var querySelect = `SELECT u.id, u.name, u.username, u.email, u.photo, u.password, u.unit_id, u.role_id, un.name as unit_name
	FROM users u LEFT JOIN units un ON un.id = u.unit_id WHERE 1=1`

// GetByID ...
func (m *mysqlUserRepository) GetByID(ctx context.Context, id uuid.UUID) (res domain.User, err error) {
	query := querySelect + fmt.Sprintf(` AND u.id = '%s'`, id)
	err = m.scan(ctx, query, &res)

	return
}

// GetByEmail ...
func (m *mysqlUserRepository) GetByEmail(ctx context.Context, email string) (res domain.User, err error) {
	query := querySelect + fmt.Sprintf(` AND email = '%s'`, email)

	err = m.scan(ctx, query, &res)

	if err != sql.ErrNoRows {
		return
	}

	return res, nil
}

func (m *mysqlUserRepository) scan(ctx context.Context, query string, res *domain.User) (err error) {
	err = m.Conn.QueryRowContext(ctx, query).Scan(
		&res.ID,
		&res.Name,
		&res.Username,
		&res.Email,
		&res.Photo,
		&res.Password,
		&res.Unit.ID,
		&res.RoleID,
		&res.UnitName,
	)
	if err != nil {
		logrus.Error("Error found", err)
	}

	return
}

func (m *mysqlUserRepository) Store(ctx context.Context, u *domain.User) (err error) {
	query := `INSERT users SET id = ?, name = ?, username = ?, email = ?, password = ?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, u.ID, u.Name, u.Username, u.Email, u.Password)
	if err != nil {
		return
	}

	return
}
