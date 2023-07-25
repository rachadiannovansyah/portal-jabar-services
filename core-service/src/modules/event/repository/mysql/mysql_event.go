package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
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

var querySelectAgenda = `SELECT id, category, title, priority, type, status, address, url, date, start_hour, end_hour, 
	created_by, created_at, updated_at FROM events WHERE deleted_at is null`
var queryJoinAgenda = `SELECT e.id, e.category, e.title, e.priority, e.type, e.status, e.address, e.url, e.date, e.start_hour, e.end_hour, 
	e.created_by, e.created_at, e.updated_at FROM events e 
	LEFT JOIN users u
	ON e.created_by = u.id
	WHERE e.deleted_at is null`

func (m *mysqlEventRepository) GetTx(ctx context.Context) (tx *sql.Tx, err error) {
	tx, err = m.Conn.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	return
}

func (r *mysqlEventRepository) fetchQuery(ctx context.Context, query string, args ...interface{}) (result []domain.Event, err error) {
	rows, err := r.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	result = make([]domain.Event, 0)
	for rows.Next() {
		event := domain.Event{}
		userID := uuid.UUID{}
		err = rows.Scan(
			&event.ID,
			&event.Category,
			&event.Title,
			&event.Priority,
			&event.Type,
			&event.Status,
			&event.Address,
			&event.URL,
			&event.Date,
			&event.StartHour,
			&event.EndHour,
			&userID,
			&event.CreatedAt,
			&event.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		event.CreatedBy = domain.User{ID: userID}
		result = append(result, event)
	}

	return result, nil
}

func (r *mysqlEventRepository) fetchQueryCalendar(ctx context.Context, query string, args ...interface{}) (result []domain.Event, err error) {
	rows, err := r.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

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

func (r *mysqlEventRepository) count(_ context.Context, query string, args ...interface{}) (total int64, err error) {

	err = r.Conn.QueryRow(query, args...).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}

func (r *mysqlEventRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.Event, total int64, err error) {
	binds := make([]interface{}, 0)
	query := filterEventQuery(params, &binds)

	if params.SortBy != "" {
		query += ` ORDER BY ` + params.SortBy + ` ` + params.SortOrder
	} else {
		query += ` ORDER BY e.date DESC `
	}

	total, _ = r.count(ctx, ` SELECT COUNT(1) FROM events e LEFT JOIN users u ON e.created_by = u.id 
			WHERE e.deleted_at is NULL `+query, binds...)

	query = queryJoinAgenda + query + ` LIMIT ?,? `

	binds = append(binds, params.Offset, params.PerPage)
	res, err = r.fetchQuery(ctx, query, binds...)
	if err != nil {
		return nil, 0, err
	}

	return
}

func (r *mysqlEventRepository) GetByID(ctx context.Context, id int64) (res domain.Event, err error) {
	query := querySelectAgenda + ` AND ID =?`

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

func (r *mysqlEventRepository) GetByTitle(ctx context.Context, title string) (res domain.Event, err error) {
	query := querySelectAgenda + ` AND title = ?`

	list, err := r.fetchQuery(ctx, query, title)
	if err != nil {
		return
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}
	return
}

func (r *mysqlEventRepository) Store(ctx context.Context, m *domain.StoreRequestEvent, tx *sql.Tx) (err error) {
	query := `INSERT events SET title=? , type=? , url=? , address=? , date=? , start_hour=? , end_hour=? , category=? , created_by=?`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx,
		m.Title,
		m.Type,
		m.URL,
		m.Address,
		m.Date,
		m.StartHour,
		m.EndHour,
		m.Category,
		m.CreatedBy.ID.String(),
	)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}

	m.ID = lastID

	return
}

func (r *mysqlEventRepository) Update(ctx context.Context, id int64, m *domain.StoreRequestEvent, tx *sql.Tx) (err error) {
	query := `UPDATE events SET title=? , type=? , url=? , address=? , date=? , start_hour=? , end_hour=? , category=? , updated_at=? WHERE id = ?`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, m.Title, m.Type, m.URL, m.Address, m.Date, m.StartHour, m.EndHour, m.Category, m.UpdatedAt, id)
	if err != nil {
		return
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowAffected != 1 {
		err = fmt.Errorf("Weird Behavior. Total affected: %d", rowAffected)
		return
	}

	return
}

func (r *mysqlEventRepository) Delete(ctx context.Context, id int64) (err error) {
	query := "UPDATE events SET deleted_at=? WHERE id = ?"
	stmt, err := r.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	deletedAt := time.Now()
	res, err := stmt.ExecContext(ctx, deletedAt, id)
	if err != nil {
		return
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowAffected != 1 {
		err = fmt.Errorf("Weird Behavior. Total Affected: %d", rowAffected)
		return
	}

	return
}

func (r *mysqlEventRepository) AgendaPortal(ctx context.Context, params *domain.Request) (res []domain.Event, total int64, err error) {
	binds := make([]interface{}, 0)
	var query string

	if params.Keyword != "" {
		binds = append(binds, "%"+params.Keyword+"%")
		query += ` AND title LIKE ? `
	}

	if params.StartDate != "" && params.EndDate != "" {
		binds = append(binds, params.StartDate, params.EndDate)
		query += ` AND (DATE(date) BETWEEN ? AND ?) `
	}

	if params.SortBy != "" {
		query += ` ORDER BY ` + params.SortBy + ` ` + params.SortOrder
	} else {
		query += ` ORDER BY date, start_hour, priority DESC `
	}

	total, _ = r.count(ctx, ` SELECT COUNT(1) FROM events WHERE deleted_at is NULL `+query, binds...)

	query = querySelectAgenda + query + ` LIMIT ?,? `
	binds = append(binds, params.Offset, params.PerPage)
	res, err = r.fetchQuery(ctx, query, binds...)
	if err != nil {
		return nil, 0, err
	}

	return
}

func (r *mysqlEventRepository) ListCalendar(ctx context.Context, params *domain.Request) (res []domain.Event, err error) {
	binds := make([]interface{}, 0)
	query := `SELECT id, title, date FROM events WHERE deleted_at is NULL`

	if params.StartDate != "" && params.EndDate != "" {
		binds = append(binds, params.StartDate, params.EndDate)
		query = query + ` AND (DATE(date) BETWEEN ? AND ?) `
	}

	query = query + ` ORDER BY date DESC `

	res, err = r.fetchQueryCalendar(ctx, query, binds...)

	return
}
