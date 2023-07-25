package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type mysqlInformationRepository struct {
	Conn *sql.DB
}

// NewMysqlInformationRepository ..
func NewMysqlInformationRepository(Conn *sql.DB) domain.InformationRepository {
	return &mysqlInformationRepository{Conn}
}

func (mr *mysqlInformationRepository) fetchQuery(ctx context.Context, query string, args ...interface{}) (result []domain.Information, err error) {
	rows, err := mr.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	result = make([]domain.Information, 0)
	for rows.Next() {
		infos := domain.Information{}
		categoryID := int64(0)
		err = rows.Scan(
			&infos.ID,
			&categoryID,
			&infos.Title,
			&infos.Content,
			&infos.Slug,
			&infos.Image,
			&infos.ShowDate,
			&infos.EndDate,
			&infos.Status,
			&infos.CreatedAt,
			&infos.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		infos.Category = domain.Category{ID: categoryID}
		result = append(result, infos)
	}

	return result, nil
}

func (mr *mysqlInformationRepository) count(_ context.Context, query string) (total int64, err error) {

	err = mr.Conn.QueryRow(query).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}

func (mr *mysqlInformationRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.Information, total int64, err error) {
	query := `SELECT id, category_id, title, content, slug, image, show_date, end_date, status, created_at, updated_at FROM informations`

	if params.Keyword != "" {
		query += ` WHERE title LIKE '%` + params.Keyword + `%' `
	}

	query += ` ORDER BY created_at LIMIT ?,? `

	res, err = mr.fetchQuery(ctx, query, params.Offset, params.PerPage)

	if err != nil {
		return nil, 0, err
	}

	total, _ = mr.count(ctx, "SELECT COUNT(1) FROM informations")

	return
}

func (mr *mysqlInformationRepository) GetByID(ctx context.Context, id int64) (res domain.Information, err error) {
	query := `SELECT id, category_id, title, content, slug, image, show_date, end_date, status, created_at, updated_at FROM informations` + ` WHERE ID = ?`

	list, err := mr.fetchQuery(ctx, query, id)
	if err != nil {
		return domain.Information{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}
