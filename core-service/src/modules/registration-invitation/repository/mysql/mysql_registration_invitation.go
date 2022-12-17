package mysql

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
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
	SELECT id, email, token, unit_id, invited_by, invited_at FROM registration_invitations WHERE 1=1
`

func (m *mysqlRegInvitationRepository) findOne(ctx context.Context,
	query string, args ...interface{}) (result domain.RegistrationInvitation, err error) {

	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return result, err
	}

	if rows.Next() {
		err = rows.Scan(
			&result.ID,
			&result.Email,
			&result.Token,
			&result.UnitID,
			&result.InvitedBy,
			&result.InvitedAt,
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

	userID := uuid.New()
	registrationInvitation.ID = &userID

	query := `INSERT INTO registration_invitations (id, email, token, unit_id, invited_by) VALUES (?, ?, ?, ?, ?)`
	_, err = m.Conn.ExecContext(
		ctx,
		query,
		registrationInvitation.ID,
		registrationInvitation.Email,
		registrationInvitation.Token,
		registrationInvitation.UnitID,
		registrationInvitation.InvitedBy,
	)

	if err != nil {
		return err
	}

	return
}

func (m *mysqlRegInvitationRepository) Update(ctx context.Context,
	id uuid.UUID, registrationInvitation *domain.RegistrationInvitation) (err error) {

	query := `UPDATE registration_invitations SET email=?, token=?, unit_id=?, invited_by=?, invited_at=? WHERE id=?`
	res, err := m.Conn.ExecContext(
		ctx,
		query,
		registrationInvitation.Email,
		registrationInvitation.Token,
		registrationInvitation.UnitID,
		registrationInvitation.InvitedBy,
		registrationInvitation.InvitedAt,
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

func (m *mysqlRegInvitationRepository) Delete(ctx context.Context,
	id uuid.UUID) (err error) {

	query := `DELETE FROM registration_invitations WHERE id=?`
	res, err := m.Conn.ExecContext(ctx, query, id)

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
