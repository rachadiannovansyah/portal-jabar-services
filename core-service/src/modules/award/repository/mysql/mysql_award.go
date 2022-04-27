package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type mysqlAwardRepository struct {
	Conn *sql.DB
}

// NewMysqlAwardRepository will create an object that represent the news.Repository interface
func NewMysqlAwardRepository(Conn *sql.DB) domain.AwardRepository {
	return &mysqlAwardRepository{Conn}
}

var querySelectAward = `SELECT id, title, logo, appreciator, description, category, created_at, updated_at FROM awards`

func (m *mysqlAwardRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Award, err error) {
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

	result = make([]domain.Award, 0)
	for rows.Next() {
		u := domain.Award{}
		err = rows.Scan(
			&u.ID,
			&u.Title,
			&u.Logo,
			&u.Appreciator,
			&u.Description,
			&u.Category,
			&u.CreatedAt,
			&u.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, u)
	}

	return result, nil
}

func (m *mysqlAwardRepository) count(ctx context.Context, query string) (total int64, err error) {

	err = m.Conn.QueryRow(query).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}

func (m *mysqlAwardRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.Award, total int64, err error) {
	query := querySelectAward

	if params.Keyword != "" {
		query += ` WHERE title LIKE '%` + params.Keyword + `%' `
	}

	query += ` ORDER BY created_at LIMIT ?,? `

	res, err = m.fetch(ctx, query, params.Offset, params.PerPage)

	if err != nil {
		return nil, 0, err
	}

	total, _ = m.count(ctx, "SELECT COUNT(1) FROM awards")

	return
}

func (m *mysqlAwardRepository) GetByID(ctx context.Context, id int64) (res domain.Award, err error) {
	query := querySelectAward + ` WHERE id = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.Award{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}
