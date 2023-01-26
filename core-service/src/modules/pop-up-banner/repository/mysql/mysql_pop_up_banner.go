package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type mysqlPopUpBannerRepository struct {
	Conn *sql.DB
}

// NewMysqlPopUpBannerRepository will create an object that represent the PopUpBanner.Repository interface
func NewMysqlPopUpBannerRepository(Conn *sql.DB) domain.PopUpBannerRepository {
	return &mysqlPopUpBannerRepository{Conn}
}

var querySelect = `SELECT id, title, button_label, image, link, status, duration, start_date, end_date, created_at, updated_at FROM pop_up_banners WHERE 1=1`

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
