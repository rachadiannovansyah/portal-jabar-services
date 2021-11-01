package mysql

import (
	"context"
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type mysqlTagsDataRepository struct {
	Conn *sql.DB
}

// NewMysqlTagsDataRepository ..
func NewMysqlTagsDataRepository(Conn *sql.DB) domain.TagsDataRepository {
	return &mysqlTagsDataRepository{Conn}
}

func (m *mysqlTagsDataRepository) GetByName(ctx context.Context, name string) (res []domain.TagsData, err error) {
	query := `SELECT id, data_id, tags_id, name, type from tags_data where name LIKE '% ? %'`

	rows, err := m.Conn.QueryContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(err)
		}
	}()

	res = make([]domain.TagsData, 0)
	for rows.Next() {
		tags := domain.TagsData{}
		err = rows.Scan(
			&tags.ID,
			&tags.Data,
			&tags.Tags,
			&tags.TagsName,
			&tags.Type,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		res = append(res, tags)
	}

	return res, nil
}
