package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type mysqlAreaRepository struct {
	Conn *sql.DB
}

// NewMysqlAreaRepository will create an object that represent the area.Repository interface
func NewMysqlAreaRepository(Conn *sql.DB) domain.AreaRepository {
	return &mysqlAreaRepository{Conn}
}

var querySelectArea = `SELECT id, depth, name, parent_code_kemendagri, code_kemendagri, code_bps, latitude, longitude, meta FROM areas WHERE 1 = 1`

func (m *mysqlAreaRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Area, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	result = make([]domain.Area, 0)
	for rows.Next() {
		a := domain.Area{}
		err = rows.Scan(
			&a.ID,
			&a.Depth,
			&a.Name,
			&a.ParentCodeKemendagri,
			&a.CodeKemendagri,
			&a.CodeBps,
			&a.Latitude,
			&a.Longtitude,
			&a.Meta,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, a)
	}

	return result, nil
}

func (m *mysqlAreaRepository) count(ctx context.Context, query string) (total int64, err error) {

	err = m.Conn.QueryRow(query).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}

func (m *mysqlAreaRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.Area, total int64, err error) {
	query := querySelectArea

	if params.Keyword != "" {
		query += ` AND name LIKE '%` + params.Keyword + `%' `
	}

	if v, ok := params.Filters["code_kemendagri"]; ok && v != "" {
		query = fmt.Sprintf(`%s AND code_kemendagri = '%s'`, query, v)
	}

	if v, ok := params.Filters["parent_code_kemendagri"]; ok && v != "" {
		query = fmt.Sprintf(`%s AND parent_code_kemendagri = '%s'`, query, v)
	}

	if v, ok := params.Filters["depth"]; ok && v != "" {
		query = fmt.Sprintf(`%s AND depth = '%s'`, query, v)
	}

	query += ` ORDER BY name LIMIT ?,? `

	res, err = m.fetch(ctx, query, params.Offset, params.PerPage)

	if err != nil {
		return nil, 0, err
	}

	total, _ = m.count(ctx, "SELECT COUNT(1) FROM areas")

	return
}

func (m *mysqlAreaRepository) GetByID(ctx context.Context, id int64) (res domain.Area, err error) {
	query := querySelectArea + ` AND id = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.Area{}, err
	}

	if len(list) > 0 {
		res = list[0]
	}

	return
}
