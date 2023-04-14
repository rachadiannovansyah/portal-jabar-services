package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/sirupsen/logrus"
)

type mysqlMdpRepository struct {
	Conn *sql.DB
}

// NewMysqlMasterDataPublicationRepository will create an object that represent the MasterDataPublication.Repository interface
func NewMysqlMasterDataPublicationRepository(Conn *sql.DB) domain.MasterDataPublicationRepository {
	return &mysqlMdpRepository{Conn}
}

var querySelectJoin = `SELECT pub.id, pub.mds_id, ms.service_name, units.name, ms.service_user, ms.technical, pub.updated_at, pub.status
FROM masterdata_publications pub
LEFT JOIN masterdata_services mds
ON pub.mds_id = mds.id
LEFT JOIN main_services ms
ON mds.main_service = ms.id
LEFT JOIN units
ON ms.opd_name = units.id
WHERE 1=1`

func (m *mysqlMdpRepository) Store(ctx context.Context, body *domain.StoreMasterDataPublication) (err error) {
	query := `INSERT masterdata_publications SET mds_id=?, portal_category=?, slug=?, cover=?, images=?, infographics=?, keywords=?, faq=?, status=?, created_at=?, updated_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx,
		body.DefaultInformation.MdsID,
		body.DefaultInformation.PortalCategory,
		body.DefaultInformation.Slug,
		helpers.GetStringFromObject(body.ServiceDescription.Cover),
		helpers.GetStringFromObject(body.ServiceDescription.Images),
		helpers.GetStringFromObject(body.ServiceDescription.InfoGraphics),
		helpers.GetStringFromObject(body.AdditionalInformation.Keywords),
		helpers.GetStringFromObject(body.AdditionalInformation.FAQ),
		body.Status,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return
	}

	return
}

func (m *mysqlMdpRepository) GetTx(ctx context.Context) (tx *sql.Tx, err error) {
	tx, err = m.Conn.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	return
}

func (m *mysqlMdpRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.MasterDataPublication, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	result = make([]domain.MasterDataPublication, 0)
	for rows.Next() {
		pub := domain.MasterDataPublication{}
		err = rows.Scan(
			&pub.ID,
			&pub.DefaultInformation.MdsID,
			&pub.DefaultInformation.ServiceName,
			&pub.DefaultInformation.OpdName,
			&pub.DefaultInformation.ServiceUser,
			&pub.DefaultInformation.Technical,
			&pub.UpdatedAt,
			&pub.Status,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, pub)
	}

	return result, nil
}

func (m *mysqlMdpRepository) count(ctx context.Context, query string, args ...interface{}) (total int64, err error) {
	err = m.Conn.QueryRow(query, args...).Scan(&total)
	if err != nil {
		logrus.Error(err)
		return
	}

	return total, nil
}

func (m *mysqlMdpRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.MasterDataPublication, total int64, err error) {
	binds := make([]interface{}, 0)
	query := filterPublicationQuery(params, &binds)

	if params.SortBy != "" {
		query += ` ORDER BY ` + params.SortBy + ` ` + params.SortOrder
	} else {
		query += ` ORDER BY updated_at DESC`
	}

	total, _ = m.count(ctx, ` SELECT COUNT(1) FROM masterdata_publications pub WHERE 1=1 `+query, binds...)
	query = querySelectJoin + query + ` LIMIT ?,? `

	binds = append(binds, params.Offset, params.PerPage)

	res, err = m.fetch(ctx, query, binds...)

	if err != nil {
		return nil, 0, err
	}

	return
}

func (m *mysqlMdpRepository) Delete(ctx context.Context, id int64) (err error) {
	query := "DELETE FROM masterdata_publications WHERE id = ?"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	err = m.rowsAffected(res)
	if err != nil {
		return
	}

	return
}

func (m *mysqlMdpRepository) rowsAffected(res sql.Result) (err error) {
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
