package mysql

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

// NewMysqlNewsRepository will create an object that represent the news.Repository interface
func NewMysqlNewsRepository(Conn *sql.DB) domain.NewsRepository {
	return &mysqlRepository{Conn}
}

var querySelectNews = `SELECT id, title, content, image, video, slug, createdAt, updatedAt FROM news`

func (m *mysqlRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.News, err error) {
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

	result = make([]domain.News, 0)
	for rows.Next() {
		t := domain.News{}
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&t.Image,
			&t.Video,
			&t.Slug,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlRepository) Fetch(ctx context.Context, params *domain.FetchNewsRequest) (res []domain.News, total int64, err error) {
	query := querySelectNews

	if params.Keyword != "" {
		query = query + ` WHERE title like '%` + params.Keyword + `%' `
	}

	query = query + ` ORDER BY createdAt LIMIT ?,? `

	res, err = m.fetch(ctx, query, params.Offset, params.PerPage)

	if err != nil {
		return nil, 0, err
	}

	total, _ = m.count(ctx, "SELECT COUNT(1) FROM news")

	return
}

func (m *mysqlRepository) GetByID(ctx context.Context, id int64) (res domain.News, err error) {
	query := querySelectNews + ` WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.News{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}
