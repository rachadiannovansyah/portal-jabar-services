package mysql

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type mysqlPublicServiceRepository struct {
	Conn *sql.DB
}

// NewMysqlPublicServiceRepository will create an object that represent the news.Repository interface
func NewMysqlPublicServiceRepository(Conn *sql.DB) domain.PublicServiceRepository {
	return &mysqlPublicServiceRepository{Conn}
}

var querySelectPservice = `SELECT id, name, description, unit, url, image, created_at, updated_at FROM public_services WHERE 1=1`

func (m *mysqlPublicServiceRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.PublicService, err error) {
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

	result = make([]domain.PublicService, 0)
	for rows.Next() {
		u := domain.PublicService{}
		err = rows.Scan(
			&u.ID,
			&u.Name,
			&u.Description,
			&u.Unit,
			&u.Url,
			&u.Image,
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

func (m *mysqlPublicServiceRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.PublicService, err error) {
	var query string

	if params.Keyword != "" {
		query += ` AND name LIKE '%` + params.Keyword + `%' `
	}

	if params.SortBy != "" {
		query += ` ORDER BY ` + params.SortBy + ` ` + params.SortOrder
	} else {
		query += ` ORDER BY created_at DESC`
	}

	query = querySelectPservice + query + ` LIMIT ?,? `

	res, err = m.fetch(ctx, query, params.Offset, params.PerPage)

	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlPublicServiceRepository) Store(ctx context.Context, ps *domain.StorePserviceRequest) (err error) {
	query := `INSERT public_service SET name=?, description=?, url=?, image=?, category=?, is_active=?, created_at=?, updated_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx,
		ps.Name,
		ps.Description,
		ps.Url,
		ps.Image,
		ps.Category,
		ps.IsActive,
		ps.CreatedAt,
		ps.UpdatedAt,
	)

	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}

	ps.ID = lastID

	return
}
