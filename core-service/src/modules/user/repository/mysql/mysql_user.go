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

var querySelect = `SELECT u.id, u.name, u.username, u.email, u.photo, u.password, last_password_changed, u.last_active, u.status, u.nip, u.occupation,
	u.unit_id, u.role_id, un.name as unit_name FROM users u LEFT JOIN units un ON un.id = u.unit_id WHERE 1=1`
var querySelectUnion = `SELECT member.id, member.name, member.email, member.role_id , member.status, member.last_active, member.occupation, member.nip
	FROM (
		select users.id, users.name, users.email, roles.id as role_id, users.status, users.last_active,
		users.occupation, users.nip, users.unit_id, users.created_at
		FROM users
		LEFT JOIN roles ON roles.id = users.role_id 
		UNION ALL
		SELECT id, "", email, 4, "` + domain.PendingUser + `" , null, null, null, unit_id, invited_at
		FROM registration_invitations
	) member WHERE 1=1`

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

func (m *mysqlUserRepository) GetByNip(ctx context.Context, nip *string) (res domain.User, err error) {
	query := querySelect + fmt.Sprintf(` AND nip = '%v'`, *nip)

	if err = m.scan(ctx, query, &res); err != sql.ErrNoRows {
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
		&res.LastActive,
		&res.Status,
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

func (m *mysqlUserRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.User, total int64, err error) {
	query := querySelectUnion

	if v, ok := params.Filters["unit_id"]; ok && v != "" {
		query = fmt.Sprintf(`%s AND member.unit_id = "%v"`, query, v)
	}

	if v, ok := params.Filters["exclude_user_id"]; ok && v != "" {
		query = fmt.Sprintf(`%s AND member.id <> "%v"`, query, v)
	}

	if params.SortBy != "" {
		query += ` ORDER BY ` + params.SortBy + ` ` + params.SortOrder
	} else {
		query += ` ORDER BY member.created_at DESC`
	}

	total, _ = m.count(ctx, query)

	query += ` LIMIT ?,? `

	res, err = m.fetch(ctx, query, params.Offset, params.PerPage)
	if err != nil {
		return nil, 0, err
	}

	return
}

func (m *mysqlUserRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.User, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	roleID := int8(0)
	result = make([]domain.User, 0)
	for rows.Next() {
		t := domain.User{}
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Email,
			&roleID,
			&t.Status,
			&t.LastActive,
			&t.Occupation,
			&t.Nip,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		t.Role.ID = roleID
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlUserRepository) count(ctx context.Context, query string) (total int64, err error) {

	err = m.Conn.QueryRow(fmt.Sprintf(`SELECT COUNT(1) FROM (%s) AS t`, query)).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}

func (m *mysqlUserRepository) WriteLastActive(ctx context.Context, time time.Time, user *domain.User) (err error) {
	query := `UPDATE users SET last_active = ? WHERE id = ?`

	_, err = m.Conn.ExecContext(ctx, query, time, user.ID)

	return
}

func (m *mysqlUserRepository) SetAsAdmin(ctx context.Context, id uuid.UUID, role int8) (err error) {
	query := `UPDATE users SET role_id=? WHERE id = ?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, role, id)
	if err != nil {
		return
	}

	return
}

func (m *mysqlUserRepository) ChangeEmail(ctx context.Context, id uuid.UUID, email string) (err error) {
	query := `UPDATE users SET email=? WHERE id = ?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, email, id)
	if err != nil {
		return
	}

	return
}

func (m *mysqlUserRepository) ChangeStatus(ctx context.Context, id uuid.UUID, status string) (err error) {
	query := `UPDATE users SET status = ? WHERE id = ?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, status, id)
	if err != nil {
		return
	}

	return
}
