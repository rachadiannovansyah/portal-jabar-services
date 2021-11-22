package mysql

import (
	"context"
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type mysqlDataTagsRepository struct {
	Conn *sql.DB
}

// NewMysqlDataTagsRepository ..
func NewMysqlDataTagsRepository(Conn *sql.DB) domain.DataTagsRepository {
	return &mysqlDataTagsRepository{Conn}
}

func (m *mysqlDataTagsRepository) fetch(ctx context.Context, query string, args ...interface{}) (res []domain.DataTags, err error) {
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

	res = make([]domain.DataTags, 0)
	for rows.Next() {
		t := domain.DataTags{}
		err = rows.Scan(
			&t.DataID,
			&t.TagsName,
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

func (m *mysqlDataTagsRepository) FetchDataTags(ctx context.Context, id int64) (res []domain.DataTags, err error) {
	query := `SELECT data_id, tags_name, type FROM tags_data WHERE data_id = ?`

	res, err = m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	return
}
