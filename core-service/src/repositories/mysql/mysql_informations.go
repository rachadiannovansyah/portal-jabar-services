package mysql

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

func NewMysqlInformationsRepository(Conn *sql.DB) domain.InformationsRepo {
	return &mysqlRepository{Conn}
}

func (mr *mysqlRepository) fetchQuery(ctx context.Context, query string, args ...interface{}) (result []domain.Informations, err error) {
	rows, err := mr.Conn.QueryContext(ctx, query, args...)
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

	result = make([]domain.Informations, 0)
	for rows.Next() {
		infos := domain.Informations{}
		err = rows.Scan(
			&infos.ID,
			&infos.Title,
			&infos.Description,
			&infos.Slug,
			&infos.Image,
			&infos.ShowDate,
			&infos.EndDate,
			&infos.CreatedBy,
			&infos.UpdatedBy,
			&infos.CreatedAt,
			&infos.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, infos)
	}

	return result, nil
}

func (mr *mysqlRepository) FetchAll(ctx context.Context, params *domain.FetchInformationsRequest) (res []domain.Informations, total int64, err error) {
	query := `SELECT * FROM informations`

	if params.Keyword != "" {
		query = query + ` WHERE title like '%` + params.Keyword + `%' `
	}

	query = query + ` ORDER BY createdAt LIMIT ?,? `

	res, err = mr.fetchQuery(ctx, query, params.Offset, params.PerPage)

	if err != nil {
		return nil, 0, err
	}

	total, _ = mr.count(ctx, "SELECT COUNT(1) FROM informations")

	return
}

func (mr *mysqlRepository) FetchOne(ctx context.Context, id int64) (res domain.Informations, err error) {
	query := `SELECT * FROM informations` + ` WHERE ID = ?`

	list, err := mr.fetchQuery(ctx, query, id)
	if err != nil {
		return domain.Informations{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}
