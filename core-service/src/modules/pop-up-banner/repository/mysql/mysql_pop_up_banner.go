package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

type mysqlPopUpBannerRepository struct {
	Conn *sql.DB
}

// NewMysqlPopUpBannerRepository will create an object that represent the PopUpBanner.Repository interface
func NewMysqlPopUpBannerRepository(Conn *sql.DB) domain.PopUpBannerRepository {
	return &mysqlPopUpBannerRepository{Conn}
}

var querySelect = `SELECT id, title, button_label, image, link, status, is_live, duration, start_date, end_date, created_at, updated_at FROM pop_up_banners WHERE 1=1`

func (m *mysqlPopUpBannerRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.PopUpBanner, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	result = make([]domain.PopUpBanner, 0)
	for rows.Next() {
		pb := domain.PopUpBanner{}
		err = rows.Scan(
			&pb.ID,
			&pb.Title,
			&pb.ButtonLabel,
			&pb.Image,
			&pb.Link,
			&pb.Status,
			&pb.IsLive,
			&pb.Duration,
			&pb.StartDate,
			&pb.EndDate,
			&pb.CreatedAt,
			&pb.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, pb)
	}

	return result, nil
}

func (m *mysqlPopUpBannerRepository) count(ctx context.Context, query string, args ...interface{}) (total int64, err error) {

	err = m.Conn.QueryRow(query, args...).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}

func (m *mysqlPopUpBannerRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.PopUpBanner, total int64, err error) {
	// add binding optional params to mitigate sql injection
	binds := make([]interface{}, 0)
	queryFilter := filterPopUpBannerQuery(params, &binds)

	defaultSort := ` ORDER BY is_live DESC, status ASC, updated_at DESC`
	if params.SortBy != "" {
		defaultSort = ` ORDER BY ` + params.SortBy + ` ` + params.SortOrder
	}
	queryFilter += defaultSort

	// get count of data
	total, _ = m.count(ctx, ` SELECT COUNT(1) FROM pop_up_banners WHERE 1=1 `+queryFilter, binds...)

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

func (m *mysqlPopUpBannerRepository) GetByID(ctx context.Context, id int64) (res domain.PopUpBanner, err error) {
	query := querySelect + " AND id = ? LIMIT 1"
	err = m.Conn.QueryRowContext(ctx, query, id).Scan(
		&res.ID,
		&res.Title,
		&res.ButtonLabel,
		&res.Image,
		&res.Link,
		&res.Status,
		&res.IsLive,
		&res.Duration,
		&res.StartDate,
		&res.EndDate,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if err != nil {
		err = domain.ErrNotFound
	}

	return
}

func (m *mysqlPopUpBannerRepository) CheckStatus(ctx context.Context, status string) (id int64, err error) {
	query := `SELECT id FROM pop_up_banners WHERE status = ? LIMIT 1`

	err = m.Conn.QueryRowContext(ctx, query, status).Scan(
		&id,
	)

	if err != nil {
		return
	}

	return
}

func (m *mysqlPopUpBannerRepository) Store(ctx context.Context, body *domain.StorePopUpBannerRequest) (err error) {
	query := `INSERT pop_up_banners SET title=?, button_label=?, link=?, image=?, duration=?,
		start_date=?, end_date=?, status=?, updated_at=?, created_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx,
		body.Title,
		body.CustomButton.Label,
		body.CustomButton.Link,
		helpers.GetStringFromObject(body.Image),
		body.Scheduler.Duration,
		body.Scheduler.StartDate,
		helpers.ConvertStringToTime(body.Scheduler.StartDate).AddDate(0, 0, int(body.Scheduler.Duration)),
		body.Scheduler.Status,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}

	body.ID = lastID

	return
}

func (m *mysqlPopUpBannerRepository) Delete(ctx context.Context, id int64) (err error) {
	query := "DELETE FROM pop_up_banners WHERE id = ?"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
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

func (m *mysqlPopUpBannerRepository) UpdateStatus(ctx context.Context, id int64, body *domain.UpdateStatusPopUpBannerRequest) (err error) {
	query := `UPDATE pop_up_banners 
		SET status = ?, 
		is_live = ?,
		start_date = ?,
		end_date = ?, 
		updated_at = ? 
		WHERE id = ?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx,
		body.Status,
		body.IsLive,
		time.Now(),
		time.Now().AddDate(0, 0, int(body.Duration)),
		time.Now(),
		id,
	)

	return
}

func (m *mysqlPopUpBannerRepository) DeactiveStatus(ctx context.Context) (err error) {
	query := `UPDATE pop_up_banners SET status = ?, is_live = ?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx,
		"NON-ACTIVE",
		0,
	)

	return
}

func (m *mysqlPopUpBannerRepository) Update(ctx context.Context, id int64, body *domain.StorePopUpBannerRequest) (err error) {
	query := `UPDATE pop_up_banners SET title=?, button_label=?, link=?, image=?, duration=?,
	start_date=?, end_date=?, status=?, updated_at=? WHERE id=?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx,
		body.Title,
		body.CustomButton.Label,
		body.CustomButton.Link,
		helpers.GetStringFromObject(body.Image),
		body.Scheduler.Duration,
		body.Scheduler.StartDate,
		helpers.ConvertStringToTime(body.Scheduler.StartDate).AddDate(0, 0, int(body.Scheduler.Duration)),
		body.Scheduler.Status,
		time.Now(),
		id,
	)

	return
}

func (m *mysqlPopUpBannerRepository) LiveBanner(ctx context.Context) (res domain.PopUpBanner, err error) {
	query := querySelect + " AND status = ? AND is_live = ? LIMIT 1"
	err = m.Conn.QueryRowContext(ctx, query, "ACTIVE", 1).Scan(
		&res.ID,
		&res.Title,
		&res.ButtonLabel,
		&res.Image,
		&res.Link,
		&res.Status,
		&res.IsLive,
		&res.Duration,
		&res.StartDate,
		&res.EndDate,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if err != nil {
		err = domain.ErrNotFound
	}

	return
}
