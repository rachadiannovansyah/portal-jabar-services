package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

// mysqlDistrictRepository ...
type mysqlDistrictRepository struct {
	Conn *sql.DB
}

// New mysqlDistrictRepository will create an object that represent the news.Repository interface
func NewMysqlDistrictRepository(Conn *sql.DB) domain.DistrictRepository {
	return &mysqlDistrictRepository{Conn}
}

var querySelectDistrict = `SELECT id, name, chief, address, website, logo, created_at, updated_at FROM districts WHERE 1=1`

func (m *mysqlDistrictRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.District, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	result = make([]domain.District, 0)
	for rows.Next() {
		u := domain.District{}
		err = rows.Scan(
			&u.ID,
			&u.Name,
			&u.Chief,
			&u.Address,
			&u.Website,
			&u.Logo,
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

func (m *mysqlDistrictRepository) count(_ context.Context, query string) (total int64, err error) {

	err = m.Conn.QueryRow(query).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}

func (m *mysqlDistrictRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.District, total int64, err error) {
	var query string

	if params.Keyword != "" {
		query += ` AND name LIKE '%` + params.Keyword + `%' `
	}

	if params.SortBy != "" {
		query += ` ORDER BY ` + params.SortBy + ` ` + params.SortOrder
	} else {
		query += ` ORDER BY created_at DESC`
	}

	total, _ = m.count(ctx, "SELECT COUNT(1) FROM districts WHERE 1=1"+query)

	query = querySelectDistrict + query + ` LIMIT ?,? `

	res, err = m.fetch(ctx, query, params.Offset, params.PerPage)

	if err != nil {
		return nil, 0, err
	}

	return
}
