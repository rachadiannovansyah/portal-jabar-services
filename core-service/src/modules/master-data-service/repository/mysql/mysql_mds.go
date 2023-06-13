package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
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

var querySelectJoin = `SELECT mds.id, ms.service_name, units.name, ms.service_user, ms.technical, mds.updated_at, mds.status, mds.has_publication, mds.main_service, mds.created_by
FROM masterdata_services mds
LEFT JOIN main_services ms
ON mds.main_service = ms.id
LEFT JOIN units
ON ms.opd_name = units.id
LEFT JOIN users u
ON mds.created_by = u.id
WHERE mds.deleted_at is NULL`

var querySelectCount = `SELECT COUNT(1) FROM masterdata_services mds
LEFT JOIN main_services ms
ON mds.main_service = ms.id
LEFT JOIN units
ON ms.opd_name = units.id
LEFT JOIN users u
ON mds.created_by = u.id
WHERE mds.deleted_at is NULL `

var querySelectJoinDetail = `SELECT mds.id, ms.service_name, units.name, ms.service_user, ms.operational_status, mds.updated_at, mds.status, mds.main_service,
ms.government_affair, ms.sub_government_affair, ms.service_form, ms.service_type, ms.program_name,
ms.description, ms.sub_service_spbe, ms.technical, ms.benefits, ms.facilities, ms.website, ms.links, ms.terms_and_condition, ms.service_procedures,
ms.service_fee, ms.operational_time, ms.hotline_number, ms.hotline_mail, ms.location,
apl.status, apl.name, apl.features, apl.title, mds.application,
aif.id, aif.responsible_name, aif.phone_number, aif.email, aif.social_media,
mds.status, mds.has_publication, mds.updated_at, mds.created_at, mds.created_by
FROM masterdata_services mds
LEFT JOIN main_services ms
ON mds.main_service = ms.id
LEFT JOIN units
ON ms.opd_name = units.id
LEFT JOIN applications apl
on mds.application = apl.id
LEFT JOIN additional_informations aif
on mds.additional_information = aif.id
WHERE deleted_at is NULL`

func (m *mysqlMdsRepository) Store(ctx context.Context, mds *domain.StoreMasterDataService, tx *sql.Tx) (err error) {
	query := `INSERT masterdata_services SET main_service=?, application=?, additional_information=?, status=?, updated_at=?, created_at=?, created_by=?`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx,
		mds.Services.ID,
		mds.Application.ID,
		mds.AdditionalInformation.ID,
		mds.Status,
		time.Now(),
		time.Now(),
		mds.CreatedBy.ID.String(),
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
			&mds.MainService.Technical,
			&mds.UpdatedAt,
			&mds.Status,
			&mds.HasPublication,
			&mds.MainService.ID,
			&mds.CreatedBy.ID,
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
		query += ` ORDER BY mds.updated_at DESC`
	}

	total, _ = m.count(ctx, querySelectCount+query, binds...)
	query = querySelectJoin + query + ` LIMIT ?,? `

	binds = append(binds, params.Offset, params.PerPage)

	res, err = m.fetch(ctx, query, binds...)

	if err != nil {
		return nil, 0, err
	}

	return
}

func (m *mysqlMdsRepository) Delete(ctx context.Context, id int64) (err error) {
	query := "UPDATE masterdata_services SET deleted_at = ? WHERE id = ?"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, time.Now(), id)
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

func (m *mysqlMdsRepository) GetByID(ctx context.Context, id int64) (res domain.MasterDataService, err error) {
	query := querySelectJoinDetail + " AND mds.id = ? LIMIT 1"

	createdByID := uuid.UUID{}
	err = m.Conn.QueryRowContext(ctx, query, id).Scan(
		&res.ID,
		&res.MainService.ServiceName,
		&res.MainService.OpdName,
		&res.MainService.ServiceUser,
		&res.MainService.OperationalStatus,
		&res.UpdatedAt,
		&res.Status,
		&res.MainService.ID,
		&res.MainService.GovernmentAffair,
		&res.MainService.SubGovernmentAffair,
		&res.MainService.ServiceForm,
		&res.MainService.ServiceType,
		&res.MainService.ProgramName,
		&res.MainService.Description,
		&res.MainService.SubServiceSpbe,
		&res.MainService.Technical,
		&res.MainService.Benefits,
		&res.MainService.Facilities,
		&res.MainService.Website,
		&res.MainService.Links,
		&res.MainService.TermsAndConditions,
		&res.MainService.ServiceProcedures,
		&res.MainService.ServiceFee,
		&res.MainService.OperationalTimes,
		&res.MainService.HotlineNumber,
		&res.MainService.HotlineMail,
		&res.MainService.Locations,
		&res.Application.Status,
		&res.Application.Name,
		&res.Application.Features,
		&res.Application.Title,
		&res.Application.ID,
		&res.AdditionalInformation.ID,
		&res.AdditionalInformation.ResponsibleName,
		&res.AdditionalInformation.PhoneNumber,
		&res.AdditionalInformation.Email,
		&res.AdditionalInformation.SocialMedia,
		&res.Status,
		&res.HasPublication,
		&res.UpdatedAt,
		&res.CreatedAt,
		&createdByID,
	)

	res.CreatedBy = domain.User{ID: createdByID}

	if err != nil {
		return domain.MasterDataService{}, domain.ErrNotFound
	}

	return
}

func (m *mysqlMdsRepository) Update(ctx context.Context, mds *domain.StoreMasterDataService, entityID *domain.MasterDataServiceEntityID, tx *sql.Tx) (err error) {
	query := `UPDATE masterdata_services SET main_service=?, application=?, additional_information=?, status=?, updated_at=? WHERE id=?`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	_, err = stmt.ExecContext(ctx,
		entityID.MainServiceID,
		entityID.ApplicationID,
		entityID.AdditionalInformationID,
		mds.Status,
		time.Now(),
		entityID.ID,
	)

	if err != nil {
		return
	}

	return
}

func (m *mysqlMdsRepository) TabStatus(ctx context.Context, params *domain.Request) (res []domain.TabStatusResponse, err error) {
	queryTabs := `
		SELECT mds.status, count(mds.status)
		FROM masterdata_services mds
		LEFT JOIN users u
		ON mds.created_by = u.id
		WHERE mds.deleted_at is NULL
	`

	binds := make([]interface{}, 0)
	query := filterMdsQuery(params, &binds)

	query = queryTabs + query + " GROUP BY mds.status"

	res, err = m.fetchTabs(ctx, query, binds...)

	if err != nil {
		return []domain.TabStatusResponse{}, err
	}

	return
}

func (m *mysqlMdsRepository) fetchTabs(ctx context.Context, query string, args ...interface{}) (result []domain.TabStatusResponse, err error) {
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

func (m *mysqlMdsRepository) Archive(ctx context.Context, params *domain.Request) (res []domain.MasterDataService, err error) {
	binds := make([]interface{}, 0)
	query := filterMdsQuery(params, &binds)
	query += ` AND has_publication = 0 ORDER BY mds.updated_at DESC`

	query = querySelectJoin + query
	res, err = m.fetch(ctx, query, binds...)
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlMdsRepository) CheckHasPublication(ctx context.Context, ID int64) (res domain.MasterDataService, err error) {
	query := "SELECT id, has_publication FROM masterdata_services WHERE has_publication != 1 AND id = ?"

	err = m.Conn.QueryRowContext(ctx, query, ID).Scan(
		&res.ID,
		&res.HasPublication,
	)

	if err != nil {
		return domain.MasterDataService{}, domain.ErrHasPublication
	}

	return
}

func (m *mysqlMdsRepository) UpdateHasPublication(ctx context.Context, ID int64, HasPublication int8) (err error) {
	query := "UPDATE masterdata_services SET has_publication = ? WHERE id = ?"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx,
		HasPublication,
		ID,
	)

	return
}
