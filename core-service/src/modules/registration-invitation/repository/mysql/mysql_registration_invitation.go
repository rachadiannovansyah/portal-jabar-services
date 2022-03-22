package mysql

import (
	"context"
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type mysqlRegInvitationRepository struct {
	Conn *sql.DB
}

func NewMysqlRegInvitationRepository(Conn *sql.DB) domain.RegistrationInvitationRepository {
	return &mysqlRegInvitationRepository{Conn}
}

var querySelectRegistrationInvitation = `
	SELECT id, email, token, expired_at, created_at, updated_at FROM registration_invitations WHERE 1=1
`

func (m *mysqlRegInvitationRepository) findOne(ctx context.Context,
	query string, args ...interface{}) (result domain.RegistrationInvitation, err error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return result, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	if rows.Next() {
		err = rows.Scan(
			&result.ID,
			&result.Email,
			&result.Token,
			&result.ExpiredAt,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
		if err != nil {
			logrus.Error(err)
			return result, err
		}
	}

	return result, nil
}

func (m *mysqlRegInvitationRepository) GetByEmail(ctx context.Context,
	email string) (result domain.RegistrationInvitation, err error) {

	query := querySelectRegistrationInvitation + " AND email=?"
	return m.findOne(ctx, query, email)
}

func (m *mysqlRegInvitationRepository) GetByToken(ctx context.Context,
	token string) (result domain.RegistrationInvitation, err error) {

	query := querySelectRegistrationInvitation + " AND token=?"
	return m.findOne(ctx, query, token)
}

func (m *mysqlRegInvitationRepository) Store(ctx context.Context,
	registrationInvitation *domain.RegistrationInvitation) (err error) {

	query := `INSERT INTO registration_invitations (email, token, expired_at) VALUES (?, ?, ?)`
	res, err := m.Conn.ExecContext(
		ctx,
		query,
		registrationInvitation.Email,
		registrationInvitation.Token,
		registrationInvitation.ExpiredAt,
	)

	if err != nil {
		logrus.Error(err)
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		logrus.Error(err)
		return
	}

	registrationInvitation.ID = id

	return
}

func (m *mysqlRegInvitationRepository) Update(ctx context.Context,
	id int64, registrationInvitation *domain.RegistrationInvitation) (err error) {

	query := `UPDATE registration_invitations SET email=?, token=?, expired_at=?, updated_at=? WHERE id=?`
	res, err := m.Conn.ExecContext(
		ctx,
		query,
		registrationInvitation.Email,
		registrationInvitation.Token,
		registrationInvitation.ExpiredAt,
		registrationInvitation.UpdatedAt,
		id,
	)

	if err != nil {
		logrus.Error(err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.Error(err)
		return
	}

	if rowsAffected == 0 {
		return domain.ErrNotFound
	}

	return
}
