package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type mysqlEventRepository struct {
	Conn *sql.DB
}

// NewMysqlEventRepository will create an object that represent the event.Repository interface
func NewMysqlEventRepository(Conn *sql.DB) domain.EventRepository {
	return &mysqlEventRepository{Conn}
}

var querySelectAgenda = `select id, category, title, priority, type, address, url, date, start_hour, end_hour, created_at, updated_at FROM events`

func (r *mysqlEventRepository) fetchQuery(ctx context.Context, query string, args ...interface{}) (result []domain.Event, err error) {
	rows, err := r.Conn.QueryContext(ctx, query, args...)
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

	result = make([]domain.Event, 0)
	for rows.Next() {
		event := domain.Event{}
		err = rows.Scan(
			&event.ID,
			&event.Category,
			&event.Title,
			&event.Priority,
			&event.Type,
			&event.Address,
			&event.URL,
			&event.Date,
			&event.StartHour,
			&event.EndHour,
			&event.CreatedAt,
			&event.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, event)
	}

	return result, nil
}

func (r *mysqlEventRepository) count(ctx context.Context, query string) (total int64, err error) {

	err = r.Conn.QueryRow(query).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}

func (r *mysqlEventRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.Event, total int64, err error) {
	query := ` WHERE 1=1`

	if params.Keyword != "" {
		query = query + ` AND title like '%` + params.Keyword + `%' `
	}

	if params.StartDate != "" && params.EndDate != "" {
		query = query + ` AND date BETWEEN '` + params.StartDate + `' AND '` + params.EndDate + `'`
	}

	query = query + ` ORDER BY date, start_hour, priority DESC `

	query = querySelectAgenda + query + ` LIMIT ?,? `

	res, err = r.fetchQuery(ctx, query, params.Offset, params.PerPage)
	if err != nil {
		return nil, 0, err
	}

	total, _ = r.count(ctx, "SELECT COUNT(1) FROM events "+query)

	return
}

func (r *mysqlEventRepository) GetByID(ctx context.Context, id int64) (res domain.Event, err error) {
	query := querySelectAgenda + ` WHERE ID = ?`

	list, err := r.fetchQuery(ctx, query, id)
	if err != nil {
		return domain.Event{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (r *mysqlEventRepository) fetchQueryCalendar(ctx context.Context, query string) (result []domain.Event, err error) {
	rows, err := r.Conn.QueryContext(ctx, query)
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

	result = make([]domain.Event, 0)
	for rows.Next() {
		event := domain.Event{}
		err = rows.Scan(
			&event.ID,
			&event.Title,
			&event.Date,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, event)
	}

	return result, nil
}

func (r *mysqlEventRepository) ListCalendar(ctx context.Context, params *domain.Request) (res []domain.Event, err error) {
	query := `SELECT id, title, date from events where 1=1`

	if params.StartDate != "" && params.EndDate != "" {
		query = query + ` AND date BETWEEN '` + params.StartDate + `' and '` + params.EndDate + `'`
	}

	query = query + ` ORDER BY date DESC `

	res, err = r.fetchQueryCalendar(ctx, query)

	return
}
