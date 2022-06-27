package mysql

import (
	"context"
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type mysqlDataTagRepository struct {
	Conn *sql.DB
}

// NewMysqlDataTagRepository ..
func NewMysqlDataTagRepository(Conn *sql.DB) domain.DataTagRepository {
	return &mysqlDataTagRepository{Conn}
}

func (m *mysqlDataTagRepository) fetch(ctx context.Context, query string, args ...interface{}) (res []domain.DataTag, err error) {
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

	res = make([]domain.DataTag, 0)
	for rows.Next() {
		t := domain.DataTag{}
		err = rows.Scan(
			&t.ID,
			&t.DataID,
			&t.TagID,
			&t.TagName,
			&t.Type,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		res = append(res, t)
	}

	return res, nil
}

func (m *mysqlDataTagRepository) FetchDataTags(ctx context.Context, id int64, domain string) (res []domain.DataTag, err error) {
	query := `SELECT id, data_id, tag_id, tag_name, type FROM data_tags WHERE data_id = ? AND type = ?`

	// TO DO
	// Create filter data tag based on tags name parameter
	res, err = m.fetch(ctx, query, id, domain)
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlDataTagRepository) StoreDataTag(ctx context.Context, dt *domain.DataTag) (err error) {
	query := `INSERT data_tags SET data_id=?, tag_id=?, tag_name=?, type=?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, dt.DataID, dt.TagID, dt.TagName, dt.Type)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	dt.ID = lastID

	return
}

func (m *mysqlDataTagRepository) DeleteDataTag(ctx context.Context, id int64, domain string) (err error) {
	query := `DELETE FROM data_tags WHERE data_id=? AND type=?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, id, domain)
	if err != nil {
		return
	}

	return
}
