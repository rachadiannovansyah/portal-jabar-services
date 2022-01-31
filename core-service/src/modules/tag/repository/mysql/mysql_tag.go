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

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

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

func (m *mysqlTagRepository) count(ctx context.Context, query string) (total int64, err error) {

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

	query += ` ORDER BY created_at LIMIT ?,? `

	res, err = m.fetch(ctx, query, params.Offset, params.PerPage)

	if err != nil {
		return nil, 0, err
	}

	total, _ = m.count(ctx, "SELECT COUNT(1) FROM tags")

	return
}

func (m *mysqlTagRepository) StoreTag(ctx context.Context, t *domain.Tag) (err error) {
	query := `REPLACE into tags (name) VALUES (?) `

	stmt, err := m.Conn.PrepareContext(ctx, query)
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
