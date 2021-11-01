package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type mysqlNewsRepository struct {
	Conn *sql.DB
}

// NewMysqlNewsRepository will create an object that represent the news.Repository interface
func NewMysqlNewsRepository(Conn *sql.DB) domain.NewsRepository {
	return &mysqlNewsRepository{Conn}
}

var querySelectNews = `SELECT id, category, title, excerpt, content, image, video, slug, author_id, type, tags, source, created_at, updated_at FROM news`

func (m *mysqlNewsRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.News, err error) {
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
		authorID := uuid.UUID{}
		err = rows.Scan(
			&t.ID,
			&t.Category,
			&t.Title,
			&t.Excerpt,
			&t.Content,
			&t.Image,
			&t.Video,
			&t.Slug,
			&authorID,
			&t.Type,
			&t.Tags,
			&t.Source,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		t.Author = domain.User{ID: authorID}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlNewsRepository) count(ctx context.Context, query string) (total int64, err error) {

	err = m.Conn.QueryRow(query).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}

func (m *mysqlNewsRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.News, total int64, err error) {
	query := querySelectNews + ` WHERE 1=1 `

	if params.Keyword != "" {
		query = query + ` AND title like '%` + params.Keyword + `%' `
	}

	if v, ok := params.Filters["highlight"]; ok && v != "" {
		query = fmt.Sprintf(`%s AND highlight = '%s'`, query, v)
	}

	if v, ok := params.Filters["category"]; ok && v != "" {
		query = fmt.Sprintf(`%s AND category = '%s'`, query, v)
	}

	if v, ok := params.Filters["type"]; ok && v != "" {
		query = fmt.Sprintf(`%s AND type = "%s"`, query, v)
	}

	if v, ok := params.Filters["tags"]; ok && v != "" {
		query = fmt.Sprintf(`%s AND tags = "%s"`, query, v)
	}

	if params.SortBy != "" {
		query = query + ` ORDER BY ` + params.SortBy + ` ` + params.SortOrder
	} else {
		query = query + ` ORDER BY created_at DESC`
	}

	query = query + ` LIMIT ?,? `

	res, err = m.fetch(ctx, query, params.Offset, params.PerPage)

	if err != nil {
		return nil, 0, err
	}

	total, _ = m.count(ctx, "SELECT COUNT(1) FROM news")

	return
}

func (m *mysqlNewsRepository) GetByID(ctx context.Context, id int64) (res domain.News, err error) {
	query := querySelectNews + ` WHERE id = ?`

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

func (m *mysqlNewsRepository) AddView(ctx context.Context, id int64) (err error) {
	query := `UPDATE news SET views = views + 1 WHERE id = ?`

	_, err = m.Conn.ExecContext(ctx, query, id)

	return
}

func (m *mysqlNewsRepository) FetchNewsBanner(ctx context.Context) (res []domain.News, err error) {
	query := querySelectNews + ` WHERE id IN (
		SELECT MAX(id) FROM news WHERE highlight = ? GROUP BY category 
	)`

	res, err = m.fetch(ctx, query, 1)
	if err != nil {
		return nil, err
	}

	return
}
