package mysql

import (
	"context"
	"database/sql"

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

func (m *mysqlSpbeRalsRepository) Fetch(ctx context.Context) (res []domain.SpbeRals, err error) {
	// exec query and params binding
	res, err = m.fetch(ctx, querySelect)
	if err != nil {
		return nil, err
	}

	return
}
