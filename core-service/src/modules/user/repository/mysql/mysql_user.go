package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

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

var querySelect = `SELECT u.id, u.name, u.username, u.email, u.photo, u.password, last_password_changed, u.nip, u.occupation,
	u.unit_id, u.role_id, un.name as unit_name FROM users u LEFT JOIN units un ON un.id = u.unit_id WHERE 1=1`

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
	roleID := int8(0)
	err = m.Conn.QueryRowContext(ctx, query).Scan(
		&res.ID,
		&res.Name,
		&res.Username,
		&res.Email,
		&res.Photo,
		&res.Password,
		&res.LastPasswordChanged,
		&res.Nip,
		&res.Occupation,
		&res.Unit.ID,
		&roleID,
		&res.UnitName,
	)
	if err != nil {
		logrus.Error("Error found", err)
	}

	res.Role.ID = roleID

	return
}

func (m *mysqlUserRepository) Store(ctx context.Context, u *domain.User) (err error) {
	query := `INSERT users SET id = ?, name = ?, username = ?, email = ?, nip = ?, 
		occupation = ?, unit_id = ?, role_id = ?, password = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, u.ID, u.Name, u.Username, u.Email,
		u.Nip, u.Occupation, u.Unit.ID, u.Role.ID, u.Password)

	if err != nil {
		return
	}

	return
}

func (m *mysqlUserRepository) Update(ctx context.Context, u *domain.User) (err error) {
	query := `UPDATE users SET name=?, username=?, email=?, password=?, last_password_changed=?, nip=?, occupation=?,
		photo=?, unit_id=?, role_id=?, updated_at=? WHERE id = ?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, u.Name, u.Username, u.Email, u.Password, u.LastPasswordChanged, u.Nip, u.Occupation,
		u.Photo, u.Unit.ID, u.Role.ID, time.Now(), u.ID)
	if err != nil {
		return
	}

	return
}
