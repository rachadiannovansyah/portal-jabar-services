package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type mysqlUnitRepository struct {
	Conn *sql.DB
}

// NewMysqlUnitRepository will create an object that represent the news.Repository interface
func NewMysqlUnitRepository(Conn *sql.DB) domain.UnitRepository {
	return &mysqlUnitRepository{Conn}
}

var querySelectUnit = `SELECT id, parent_id, name, description, logo, website, phone, address, chief, 
	created_at, updated_at FROM units`

func (m *mysqlUnitRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Unit, err error) {
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

	result = make([]domain.Unit, 0)
	for rows.Next() {
		u := domain.Unit{}
		err = rows.Scan(
			&u.ID,
			&u.ParentID,
			&u.Name,
			&u.Description,
			&u.Logo,
			&u.Website,
			&u.Phone,
			&u.Address,
			&u.Chief,
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

func (m *mysqlUnitRepository) count(ctx context.Context, query string) (total int64, err error) {

	err = m.Conn.QueryRow(query).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}

func (m *mysqlUnitRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.Unit, total int64, err error) {
	query := querySelectUnit

	if params.Keyword != "" {
		query = query + ` WHERE name like '%` + params.Keyword + `%' `
	}

	query = query + ` ORDER BY created_at LIMIT ?,? `

	res, err = m.fetch(ctx, query, params.Offset, params.PerPage)

	if err != nil {
		return nil, 0, err
	}

	total, _ = m.count(ctx, "SELECT COUNT(1) FROM units")

	return
}

func (m *mysqlUnitRepository) GetByID(ctx context.Context, id int64) (res domain.Unit, err error) {
	query := querySelectUnit + ` WHERE id = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.Unit{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}
