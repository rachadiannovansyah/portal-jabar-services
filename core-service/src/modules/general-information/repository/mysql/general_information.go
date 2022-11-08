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

type mysqlGeneralInformationRepository struct {
	Conn *sql.DB
}

// NewMysqlGeneralInformationRepository will create an object that represent the GeneralInformation.Repository interface
func NewMysqlGeneralInformationRepository(Conn *sql.DB) domain.GeneralInformationRepository {
	return &mysqlGeneralInformationRepository{Conn}
}

var querySelectGenInfo = `SELECT id, name, description, slug, category, addresses, unit, phone, logo, operational_hours, media, social_media, type FROM general_informations WHERE 1=1`

func (m *mysqlGeneralInformationRepository) GetTx(ctx context.Context) (tx *sql.Tx, err error) {
	tx, err = m.Conn.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	return
}

func (m *mysqlGeneralInformationRepository) fetchGenInfo(ctx context.Context, query string, args ...interface{}) (result []domain.GeneralInformation, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
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

	result = make([]domain.GeneralInformation, 0)
	for rows.Next() {
		genInfo := domain.GeneralInformation{}
		err = rows.Scan(
			&genInfo.ID,
			&genInfo.Name,
			&genInfo.Description,
			&genInfo.Slug,
			&genInfo.Category,
			&genInfo.Addresses,
			&genInfo.Unit,
			&genInfo.Phone,
			&genInfo.Logo,
			&genInfo.OperationalHours,
			&genInfo.Media,
			&genInfo.SocialMedia,
			&genInfo.Type,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, genInfo)
	}

	return result, nil
}

func (m *mysqlGeneralInformationRepository) GetByID(ctx context.Context, id int64) (res domain.GeneralInformation, err error) {
	query := fmt.Sprintf("%s AND id=?", querySelectGenInfo)

	list, err := m.fetchGenInfo(ctx, query, id)
	if err != nil {
		return domain.GeneralInformation{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlGeneralInformationRepository) Store(ctx context.Context, ps domain.StorePublicService, tx *sql.Tx) (ID int64, err error) {
	query := `INSERT general_informations SET name=?, alias=?, email=?, description=?, category=?, 
	addresses=?, unit=?, phone=?, logo=?, operational_hours=?, link=?, media=?, social_media=?, type=?`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx,
		ps.GeneralInformation.Name,
		ps.GeneralInformation.Alias,
		ps.GeneralInformation.Email,
		ps.GeneralInformation.Description,
		ps.GeneralInformation.Category,
		helpers.GetStringFromObject(ps.GeneralInformation.Addresses),
		ps.GeneralInformation.Unit,
		helpers.GetStringFromObject(ps.GeneralInformation.Phone),
		ps.GeneralInformation.Logo,
		helpers.GetStringFromObject(ps.GeneralInformation.OperationalHours),
		helpers.GetStringFromObject(ps.GeneralInformation.Link),
		helpers.GetStringFromObject(ps.GeneralInformation.Media),
		helpers.GetStringFromObject(ps.GeneralInformation.SocialMedia),
		ps.GeneralInformation.Type,
	)
	if err != nil {
		return
	}
	ID, err = res.LastInsertId()
	if err != nil {
		return
	}

	err = m.UpdateSlug(ctx, ps, ID, tx)
	if err != nil {
		return
	}
	return
}

func (m *mysqlGeneralInformationRepository) UpdateSlug(ctx context.Context, ps domain.StorePublicService, ID int64, tx *sql.Tx) (err error) {
	query := `UPDATE general_informations SET slug=? WHERE id=?`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	slug := helpers.MakeSlug(ps.GeneralInformation.Name, ID)

	_, err = stmt.ExecContext(ctx, slug, ID)

	if err != nil {
		return
	}
	return
}

func (m *mysqlGeneralInformationRepository) Update(ctx context.Context, ps domain.UpdatePublicService, ID int64, tx *sql.Tx) (err error) {
	query := `UPDATE general_informations SET name=?, slug=?, alias=?, email=?, description=?, category=?, 
	addresses=?, unit=?, phone=?, logo=?, operational_hours=?, link=?, media=?, social_media=?, type=?, updated_at=? WHERE id = ?`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx,
		ps.GeneralInformation.Name,
		ps.GeneralInformation.Slug,
		ps.GeneralInformation.Alias,
		ps.GeneralInformation.Email,
		ps.GeneralInformation.Description,
		ps.GeneralInformation.Category,
		helpers.GetStringFromObject(ps.GeneralInformation.Addresses),
		ps.GeneralInformation.Unit,
		helpers.GetStringFromObject(ps.GeneralInformation.Phone),
		ps.GeneralInformation.Logo,
		helpers.GetStringFromObject(ps.GeneralInformation.OperationalHours),
		helpers.GetStringFromObject(ps.GeneralInformation.Link),
		helpers.GetStringFromObject(ps.GeneralInformation.Media),
		helpers.GetStringFromObject(ps.GeneralInformation.SocialMedia),
		ps.GeneralInformation.Type,
		time.Now(),
		ID,
	)

	return
}
