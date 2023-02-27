package mysql

import (
	"context"
	"database/sql"
	"fmt"

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

func (m *mysqlGovernmentAffairRepository) count(ctx context.Context, query string, args ...interface{}) (total int64, err error) {

	err = m.Conn.QueryRow(query, args...).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}

func (m *mysqlGovernmentAffairRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.GovernmentAffair, total int64, err error) {
	// add binding optional params to mitigate sql injection
	binds := make([]interface{}, 0)
	queryFilter := filterGovernmentAffairQuery(params, &binds)

	// get count of data
	total, _ = m.count(ctx, ` SELECT COUNT(1) FROM government_affairs WHERE 1=1 `+queryFilter, binds...)

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
