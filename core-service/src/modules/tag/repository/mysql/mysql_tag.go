package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type mysqlTagRepository struct {
	Conn *sql.DB
}

// NewMysqlTagRepository ..
func NewMysqlTagRepository(Conn *sql.DB) domain.TagRepository {
	return &mysqlTagRepository{Conn}
}

var querySelectTags = `SELECT id, name FROM tags`

func (m *mysqlTagRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Tag, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	result = make([]domain.Tag, 0)
	for rows.Next() {
		t := domain.Tag{}
		err = rows.Scan(
			&t.ID,
			&t.Name,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlTagRepository) count(_ context.Context, query string) (total int64, err error) {

	err = m.Conn.QueryRow(query).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}

func (m *mysqlTagRepository) FetchTag(ctx context.Context, params *domain.Request) (res []domain.Tag, total int64, err error) {
	query := querySelectTags

	if params.Keyword != "" {
		query += ` WHERE name LIKE '%` + params.Keyword + `%' `
	}

	query += ` LIMIT 5`

	res, err = m.fetch(ctx, query)

	if err != nil {
		return nil, 0, err
	}

	total, _ = m.count(ctx, "SELECT COUNT(1) FROM tags")

	return
}

func (m *mysqlTagRepository) StoreTag(ctx context.Context, t *domain.Tag, tx *sql.Tx) (err error) {
	query := `REPLACE into tags (name) VALUES (?) `

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, t.Name)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	t.ID = lastID

	return
}

func (m *mysqlTagRepository) GetTagByName(ctx context.Context, name string) (res domain.Tag, err error) {
	query := querySelectTags + ` WHERE name = ?`

	list, err := m.fetch(ctx, query, name)
	if err != nil {
		return
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}
