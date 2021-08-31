package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type mysqlAgendaRepository struct {
	Conn *sql.DB
}

func NewMysqlAgendaRepository(Conn *sql.DB) domain.AgendaRepository {
	return &mysqlAgendaRepository{Conn}
}

func (mr *mysqlAgendaRepository) fetchQuery(ctx context.Context, query string, args ...interface{}) (result []domain.Agenda, err error) {
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

	result = make([]domain.Agenda, 0)
	for rows.Next() {
		agenda := domain.Agenda{}
		categoryID := int64(0)
		err = rows.Scan(
			&agenda.ID,
			&categoryID,
			&agenda.Title,
			&agenda.Description,
			&agenda.Date,
			&agenda.Address,
			&agenda.StartHour,
			&agenda.EndHour,
			&agenda.Image,
			&agenda.PublishedBy,
			&agenda.Province,
			&agenda.City,
			&agenda.District,
			&agenda.Village,
			&agenda.CreatedAt,
			&agenda.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		agenda.Category = domain.Category{ID: categoryID}
		result = append(result, agenda)
	}

	return result, nil
}

func (mr *mysqlAgendaRepository) count(ctx context.Context, query string) (total int64, err error) {

	err = mr.Conn.QueryRow(query).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}

func (mr *mysqlAgendaRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.Agenda, total int64, err error) {
	query := `SELECT id, category_id, title, description, date, address, start_hour, end_hour, image, published_by, province_code, city_code, district_code, village_code, created_at, updated_at FROM agenda`

	if params.Keyword != "" {
		query = query + ` WHERE title like '%` + params.Keyword + `%' `
	}

	query = query + ` ORDER BY created_at LIMIT ?,? `

	res, err = mr.fetchQuery(ctx, query, params.Offset, params.PerPage)

	if err != nil {
		return nil, 0, err
	}

	total, _ = mr.count(ctx, "SELECT COUNT(1) FROM agenda")

	return
}

func (mr *mysqlAgendaRepository) GetByID(ctx context.Context, id int64) (res domain.Agenda, err error) {
	query := `SELECT id, category_id, title, description, date, address, start_hour, end_hour, image, published_by, province_code, city_code, district_code, village_code, created_at, updated_at FROM informations` + ` WHERE ID = ?`

	list, err := mr.fetchQuery(ctx, query, id)
	if err != nil {
		return domain.Agenda{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}
