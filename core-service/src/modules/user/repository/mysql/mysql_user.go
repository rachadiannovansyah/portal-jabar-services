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

func (m *mysqlUserRepository) MemberList(ctx context.Context, params *domain.Request) (res []domain.MemberList, total int64, err error) {
	queryUnion := `SELECT member.id, member.name, member.email, member.role_name , member.status
	FROM (
		select users.id, users.name, users.email, roles.name as role_name, "active" as status
		FROM users
		LEFT JOIN roles ON roles.id = users.role_id 
		UNION ALL
		SELECT id, null, email, "Member", "waiting confirmation"
		FROM registration_invitations
	) member`

	if v, ok := params.Filters["name"]; ok && v != "" {
		queryUnion = fmt.Sprintf(`%s AND member.name = "%s"`, queryUnion, v)
	}

	if v, ok := params.Filters["email"]; ok && v != "" {
		queryUnion = fmt.Sprintf(`%s AND member.email = "%s"`, queryUnion, v)
	}

	if v, ok := params.Filters["status"]; ok && v != "" {
		queryUnion = fmt.Sprintf(`%s AND member.status = "%s"`, queryUnion, v)
	}

	queryUnion += ` ORDER BY name DESC`

	total, _ = m.count(ctx, `SELECT COUNT(1) 
							FROM (
									SELECT users.id
									FROM users
									LEFT JOIN roles ON roles.id = users.role_id 
									UNION ALL
									SELECT id
									FROM registration_invitations
							) member ORDER BY id ASC`)

	queryUnion = queryUnion + ` LIMIT ?,? `

	res, err = m.fetch(ctx, queryUnion, params.Offset, params.PerPage)
	if err != nil {
		return nil, 0, err
	}

	return
}

func (m *mysqlUserRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.MemberList, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
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

	result = make([]domain.MemberList, 0)
	for rows.Next() {
		t := domain.MemberList{}
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Email,
			&t.Role,
			&t.Status,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlUserRepository) count(ctx context.Context, query string) (total int64, err error) {

	err = m.Conn.QueryRow(query).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}
