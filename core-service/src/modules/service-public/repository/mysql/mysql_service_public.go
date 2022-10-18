package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

type mysqlServicePublicRepository struct {
	Conn *sql.DB
}

// NewMysqlServicePublicRepository will create an object that represent the ServicePublic.Repository interface
func NewMysqlServicePublicRepository(Conn *sql.DB) domain.ServicePublicRepository {
	return &mysqlServicePublicRepository{Conn}
}

var querySelectJoin = `SELECT s.id, s.purpose, s.facility, s.requirement, s.tos, s.info_graphic, s.faq, s.created_at, s.updated_at,
g.ID, g.name, g.Description, g.slug, g.category, g.address, g.unit, g.phone, g.logo, g.operational_hours, g.media, g.social_media, g.type
FROM service_public s
LEFT JOIN general_informations g
ON s.general_information_id = g.id
WHERE 1=1`

func (m *mysqlServicePublicRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.ServicePublic, err error) {
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

	result = make([]domain.ServicePublic, 0)
	for rows.Next() {
		ps := domain.ServicePublic{}
		err = rows.Scan(
			&ps.ID,
			&ps.Purpose,
			&ps.Facility,
			&ps.Requirement,
			&ps.ToS,
			&ps.InfoGraphic,
			&ps.FAQ,
			&ps.CreatedAt,
			&ps.UpdatedAt,
			&ps.GeneralInformation.ID,
			&ps.GeneralInformation.Name,
			&ps.GeneralInformation.Description,
			&ps.GeneralInformation.Slug,
			&ps.GeneralInformation.Category,
			&ps.GeneralInformation.Address,
			&ps.GeneralInformation.Unit,
			&ps.GeneralInformation.Phone,
			&ps.GeneralInformation.Logo,
			&ps.GeneralInformation.OperationalHours,
			&ps.GeneralInformation.Media,
			&ps.GeneralInformation.SocialMedia,
			&ps.GeneralInformation.Type,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, ps)
	}

	return result, nil
}

func (m *mysqlServicePublicRepository) count(ctx context.Context, query string) (total int64, err error) {

	err = m.Conn.QueryRow(query).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}

func (m *mysqlServicePublicRepository) getLastUpdated(ctx context.Context, query string) (lastUpdated string, err error) {
	err = m.Conn.QueryRow(query).Scan(&lastUpdated)

	if err == sql.ErrNoRows {
		// there were no rows, but otherwise no error occurred
		log.Println(err)
	}

	return lastUpdated, nil
}

func (m *mysqlServicePublicRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.ServicePublic, err error) {
	query := querySelectJoin + ` LIMIT ?,? `

	res, err = m.fetch(ctx, query, params.Offset, params.PerPage)
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlServicePublicRepository) MetaFetch(ctx context.Context, params *domain.Request) (total int64, lastUpdated string, err error) {
	total, _ = m.count(ctx, ` SELECT COUNT(1) FROM service_public `)

	lastUpdated, err = m.getLastUpdated(ctx, ` SELECT updated_at FROM service_public ORDER BY updated_at DESC LIMIT 1`)

	if err != nil {
		return 0, "", err
	}

	return
}

func (m *mysqlServicePublicRepository) GetBySlug(ctx context.Context, slug string) (res domain.ServicePublic, err error) {
	return m.findOne(ctx, slug)
}

func (m *mysqlServicePublicRepository) findOne(ctx context.Context, value string) (res domain.ServicePublic, err error) {
	query := fmt.Sprintf("%s AND g.slug=?", querySelectJoin)

	list, err := m.fetch(ctx, query, value)
	if err != nil {
		return domain.ServicePublic{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlServicePublicRepository) Store(ctx context.Context, ps domain.StorePublicService) (err error) {
	tx, err := m.Conn.BeginTx(ctx, nil)
	id, err := m.StoreGeneralInformation(ctx, tx, ps)

	query := `INSERT service_public SET general_information_id=?, purpose=?, facility=?, requirement=?, 
		tos=?, info_graphic=?, faq=?`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx,
		id,
		helpers.GetStringFromObject(ps.Purpose),
		helpers.GetStringFromObject(ps.Facility),
		helpers.GetStringFromObject(ps.Requirement),
		helpers.GetStringFromObject(ps.Tos),
		helpers.GetStringFromObject(ps.Infographic),
		helpers.GetStringFromObject(ps.Faq),
	)
	if err != nil {
		return
	}
	_, err = res.LastInsertId()
	if err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}

func (m *mysqlServicePublicRepository) StoreGeneralInformation(ctx context.Context, tx *sql.Tx, ps domain.StorePublicService) (id int64, err error) {
	query := `INSERT general_informations SET name=?, description=?, slug=?, category=?, 
	address=?, unit=?, phone=?, logo=?, operational_hours=?, media=?, social_media=?, type=?`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx,
		ps.GeneralInformation.Name,
		ps.GeneralInformation.Description,
		ps.GeneralInformation.Slug,
		ps.GeneralInformation.Category,
		ps.GeneralInformation.Address,
		ps.GeneralInformation.Unit,
		helpers.GetStringFromObject(ps.GeneralInformation.Phone),
		ps.GeneralInformation.Logo,
		helpers.GetStringFromObject(ps.GeneralInformation.OperationalHours),
		helpers.GetStringFromObject(ps.GeneralInformation.Media),
		helpers.GetStringFromObject(ps.GeneralInformation.SocialMedia),
		ps.GeneralInformation.Type,
	)
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	if err != nil {
		return
	}
	return
}
