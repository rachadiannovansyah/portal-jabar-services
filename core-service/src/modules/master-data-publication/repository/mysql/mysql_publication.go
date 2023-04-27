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

var querySelectJoinDetail = `SELECT mdp.id, unit.name, ms.service_form, ms.service_name, ms.program_name, ms.description, ms.service_user, mdp.portal_category, ms.operational_status, ms.technical, ms.benefits, ms.facilities, mdp.slug,
mdp.cover, mdp.images, ms.terms_and_condition, ms.service_procedures, ms.service_fee, ms.operational_time, ms.hotline_number, ms.hotline_mail, mdp.infographics,
ms.location, ap.ID, ap.name, ap.status, ap.features, ap.title, ms.links, aif.social_media, mdp.keywords, mdp.faq, mdp.status, mdp.created_at, mdp.updated_at
FROM masterdata_publications as mdp
LEFT JOIN masterdata_services as mds
ON mdp.mds_id = mds.id
LEFT JOIN main_services as ms
on mds.main_service = ms.id
LEFT JOIN applications as ap
on mds.application = ap.id
LEFT JOIN additional_informations as aif
on mds.additional_information = aif.id
LEFT JOIN units as unit
on ms.opd_name = unit.id
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

func (m *mysqlMdpRepository) GetByID(ctx context.Context, id int64) (res domain.MasterDataPublication, err error) {
	query := querySelectJoinDetail + " AND mdp.id = ? LIMIT 1"

	err = m.Conn.QueryRowContext(ctx, query, id).Scan(
		&res.ID,
		&res.DefaultInformation.OpdName,
		&res.DefaultInformation.ServiceForm,
		&res.DefaultInformation.ServiceName,
		&res.DefaultInformation.ProgramName,
		&res.DefaultInformation.Description,
		&res.DefaultInformation.ServiceUser,
		&res.DefaultInformation.PortalCategory,
		&res.DefaultInformation.OperationalStatus,
		&res.DefaultInformation.Technical,
		&res.DefaultInformation.Benefits,
		&res.DefaultInformation.Facilities,
		&res.DefaultInformation.Slug,
		&res.ServiceDescription.Cover,
		&res.ServiceDescription.Images,
		&res.ServiceDescription.TermsAndConditions,
		&res.ServiceDescription.ServiceProcedures,
		&res.ServiceDescription.ServiceFee,
		&res.ServiceDescription.OperationalTimes,
		&res.ServiceDescription.HotlineNumber,
		&res.ServiceDescription.HotlineMail,
		&res.ServiceDescription.InfoGraphics,
		&res.ServiceDescription.Locations,
		&res.ServiceDescription.Application.ID,
		&res.ServiceDescription.Application.Name,
		&res.ServiceDescription.Application.Status,
		&res.ServiceDescription.Application.Features,
		&res.ServiceDescription.Application.Title,
		&res.ServiceDescription.Links,
		&res.ServiceDescription.SocialMedia,
		&res.AdditionalInformation.Keywords,
		&res.AdditionalInformation.FAQ,
		&res.Status,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if err != nil {
		return
	}

	return
}

func (m *mysqlMdpRepository) TabStatus(ctx context.Context) (res []domain.TabStatusResponse, err error) {
	query := `
		SELECT mdp.status, count(mdp.status)
		FROM masterdata_publications mdp
		WHERE 1=1
		GROUP BY mdp.status
	`

	res, err = m.fetchTabs(ctx, query)

	if err != nil {
		return []domain.TabStatusResponse{}, err
	}

	return
}

func (m *mysqlMdpRepository) fetchTabs(ctx context.Context, query string, args ...interface{}) (result []domain.TabStatusResponse, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	result = make([]domain.TabStatusResponse, 0)
	for rows.Next() {
		t := domain.TabStatusResponse{}
		err = rows.Scan(
			&t.Status,
			&t.Count,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlMdpRepository) Update(ctx context.Context, body *domain.StoreMasterDataPublication, ID int64) (err error) {
	query := `UPDATE masterdata_publications SET mds_id=?, portal_category=?, slug=?, cover=?, images=?, infographics=?, keywords=?, faq=?, status=?, created_at=?, updated_at=?
		WHERE id=?`
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
		ID,
	)

	if err != nil {
		return
	}

	return
}
