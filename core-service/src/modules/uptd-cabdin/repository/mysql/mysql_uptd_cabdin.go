package mysql

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type mysqlUptdCabdinRepository struct {
	Conn *sql.DB
}

// NewMysqlUptdCabdinRepository will create an object that represent the UptdCabdin.Repository interface
func NewMysqlUptdCabdinRepository(Conn *sql.DB) domain.UptdCabdinRepository {
	return &mysqlUptdCabdinRepository{Conn}
}

var querySelect = `SELECT id, prk_name, cbg_name, cbg_kotakab, cbg_alamat, cbg_notlp, cbg_jenis FROM uptd_cabdins WHERE 1=1`

func (m *mysqlUptdCabdinRepository) fetch(ctx context.Context, query string, args ...interface{}) (results []domain.UptdCabdin, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	results = make([]domain.UptdCabdin, 0)
	for rows.Next() {
		uc := domain.UptdCabdin{}
		err = rows.Scan(
			&uc.ID,
			&uc.PrkName,
			&uc.CbgName,
			&uc.CbgKotaKab,
			&uc.CbgAlamat,
			&uc.CbgNoTlp,
			&uc.CbgJenis,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		results = append(results, uc)
	}

	return results, nil
}

func (m *mysqlUptdCabdinRepository) Fetch(ctx context.Context) (res []domain.UptdCabdin, err error) {
	// exec query and params binding
	res, err = m.fetch(ctx, querySelect)
	if err != nil {
		return nil, err
	}

	return
}
