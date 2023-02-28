package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type mysqlSpbeRalsRepository struct {
	Conn *sql.DB
}

// NewMysqlSpbeRalsRepository will create an object that represent the SpbeRals.Repository interface
func NewMysqlSpbeRalsRepository(Conn *sql.DB) domain.SpbeRalsRepository {
	return &mysqlSpbeRalsRepository{Conn}
}

var querySelect = `SELECT id, kode_ral_2, kode, item FROM spbe_rals WHERE 1=1`

func (m *mysqlSpbeRalsRepository) fetch(ctx context.Context, query string, args ...interface{}) (results []domain.SpbeRals, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	results = make([]domain.SpbeRals, 0)
	for rows.Next() {
		sr := domain.SpbeRals{}
		err = rows.Scan(
			&sr.ID,
			&sr.RalCode2,
			&sr.Code,
			&sr.Item,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		results = append(results, sr)
	}

	return results, nil
}

func (m *mysqlSpbeRalsRepository) count(ctx context.Context, query string, args ...interface{}) (total int64, err error) {

	err = m.Conn.QueryRow(query, args...).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}

func (m *mysqlSpbeRalsRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.SpbeRals, total int64, err error) {
	// add binding optional params to mitigate sql injection
	binds := make([]interface{}, 0)
	queryFilter := filterSpbeRalsQuery(params, &binds)

	// get count of data
	total, _ = m.count(ctx, ` SELECT COUNT(1) FROM spbe_rals WHERE 1=1 `+queryFilter, binds...)

	// appending final query
	query := querySelect + queryFilter + ` LIMIT ?,? `
	binds = append(binds, params.Offset, params.PerPage)

	// exec query and params binding
	res, err = m.fetch(ctx, query, binds...)
	if err != nil {
		return nil, 0, err
	}

	return
}
