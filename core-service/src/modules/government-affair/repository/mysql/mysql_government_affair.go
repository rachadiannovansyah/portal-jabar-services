package mysql

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type mysqlGovernmentAffairRepository struct {
	Conn *sql.DB
}

// NewMysqlGovernmentAffairRepository will create an object that represent the GovernmentAffair.Repository interface
func NewMysqlGovernmentAffairRepository(Conn *sql.DB) domain.GovernmentAffairRepository {
	return &mysqlGovernmentAffairRepository{Conn}
}

var querySelect = `SELECT id, main_affair, sub_main_affair FROM government_affairs WHERE 1=1`

func (m *mysqlGovernmentAffairRepository) fetch(ctx context.Context, query string, args ...interface{}) (results []domain.GovernmentAffair, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	results = make([]domain.GovernmentAffair, 0)
	for rows.Next() {
		ga := domain.GovernmentAffair{}
		err = rows.Scan(
			&ga.ID,
			&ga.MainAffair,
			&ga.SubMainAffair,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		results = append(results, ga)
	}

	return results, nil
}

func (m *mysqlGovernmentAffairRepository) Fetch(ctx context.Context) (res []domain.GovernmentAffair, err error) {
	// exec query and params binding
	res, err = m.fetch(ctx, querySelect)
	if err != nil {
		return nil, err
	}

	return
}
