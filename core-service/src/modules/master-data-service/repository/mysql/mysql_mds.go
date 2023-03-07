package mysql

import (
	"context"
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type mysqlMdsRepository struct {
	Conn *sql.DB
}

// NewMysqlMasterDataServiceRepository will create an object that represent the MasterDataService.Repository interface
func NewMysqlMasterDataServiceRepository(Conn *sql.DB) domain.MasterDataServiceRepository {
	return &mysqlMdsRepository{Conn}
}

var querySelectJoin = `SELECT mds.id, ms.service_name, units.name, ms.service_user, ms.operational_status, mds.updated_at, mds.status, mds.main_service
FROM masterdata_services mds
LEFT JOIN main_services ms
ON mds.main_service = ms.id
LEFT JOIN units
ON ms.opd_name = units.id
WHERE 1=1`

func (m *mysqlMdsRepository) Store(ctx context.Context, mds *domain.StoreMasterDataService, tx *sql.Tx) (err error) {
	query := `INSERT masterdata_services SET main_service=?, application=?, additional_information=?, status=?`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx,
		mds.Services.ID,
		mds.Application.ID,
		mds.AdditionalInformation.ID,
		mds.Status,
	)

	if err != nil {
		return
	}

	return
}

func (m *mysqlMdsRepository) GetTx(ctx context.Context) (tx *sql.Tx, err error) {
	tx, err = m.Conn.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	return
}

func (m *mysqlMdsRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.MasterDataService, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	result = make([]domain.MasterDataService, 0)
	for rows.Next() {
		mds := domain.MasterDataService{}
		err = rows.Scan(
			&mds.ID,
			&mds.MainService.ServiceName,
			&mds.MainService.OpdName,
			&mds.MainService.ServiceUser,
			&mds.MainService.OperationalStatus,
			&mds.UpdatedAt,
			&mds.Status,
			&mds.MainService.ID,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, mds)
	}

	return result, nil
}

func (m *mysqlMdsRepository) count(ctx context.Context, query string, args ...interface{}) (total int64, err error) {
	err = m.Conn.QueryRow(query, args...).Scan(&total)
	if err != nil {
		logrus.Error(err)
		return
	}

	return total, nil
}

func (m *mysqlMdsRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.MasterDataService, total int64, err error) {
	binds := make([]interface{}, 0)
	query := filterMdsQuery(params, &binds)

	if params.SortBy != "" {
		query += ` ORDER BY ` + params.SortBy + ` ` + params.SortOrder
	} else {
		query += ` ORDER BY updated_at DESC`
	}

	total, _ = m.count(ctx, ` SELECT COUNT(1) FROM masterdata_services WHERE 1=1 `+query, binds...)
	query = querySelectJoin + query + ` LIMIT ?,? `

	binds = append(binds, params.Offset, params.PerPage)

	res, err = m.fetch(ctx, query, binds...)

	if err != nil {
		return nil, 0, err
	}

	return
}

func (m *mysqlMdsRepository) Delete(ctx context.Context, id int64) (err error) {
	query := "DELETE FROM masterdata_services WHERE id = ?"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	m.rowsAffected(res) // for fix code complexity exceed return on one func

	return
}

func (m *mysqlMdsRepository) rowsAffected(res sql.Result) (err error) {
	rowAffected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowAffected != 1 {
		logrus.Error(err)
		return
	}

	return
}
