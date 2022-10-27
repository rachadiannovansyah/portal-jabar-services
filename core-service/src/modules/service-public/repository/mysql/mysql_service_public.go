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
g.ID, g.name, g.alias, g.Description, g.slug, g.category, g.addresses, g.unit, g.phone, g.email, g.logo, g.operational_hours, g.link, g.media, g.social_media, g.type
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
			&ps.GeneralInformation.Alias,
			&ps.GeneralInformation.Description,
			&ps.GeneralInformation.Slug,
			&ps.GeneralInformation.Category,
			&ps.GeneralInformation.Addresses,
			&ps.GeneralInformation.Unit,
			&ps.GeneralInformation.Phone,
			&ps.GeneralInformation.Email,
			&ps.GeneralInformation.Logo,
			&ps.GeneralInformation.OperationalHours,
			&ps.GeneralInformation.Link,
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

func (m *mysqlServicePublicRepository) count(ctx context.Context, query string, args ...interface{}) (total int64, err error) {

	err = m.Conn.QueryRow(query, args...).Scan(&total)
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
	// add binding optional params to mitigate sql injection
	binds := make([]interface{}, 0)
	queryFilter := filterServicePublicQuery(params, &binds)

	query := querySelectJoin + queryFilter + ` LIMIT ?,? `
	binds = append(binds, params.Offset, params.PerPage)

	res, err = m.fetch(ctx, query, binds...)
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlServicePublicRepository) MetaFetch(ctx context.Context, params *domain.Request) (total int64, lastUpdated string, staticCount int64, err error) {
	binds := make([]interface{}, 0)
	queryFilter := filterServicePublicQuery(params, &binds)

	total, _ = m.count(ctx, ` SELECT COUNT(1) FROM service_public s LEFT JOIN general_informations g ON s.general_information_id = g.id WHERE 1=1 `+queryFilter, binds...)

	lastUpdated, err = m.getLastUpdated(ctx, ` SELECT updated_at FROM service_public ORDER BY updated_at DESC LIMIT 1`)

	staticCount, _ = m.count(ctx, ` SELECT COUNT(1) FROM service_public s LEFT JOIN general_informations g ON s.general_information_id = g.id WHERE 1=1 AND g.category = ?`, params.Filters["category"].(string))

	if err != nil {
		return 0, "", 0, err
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

func (m *mysqlServicePublicRepository) Store(ctx context.Context, ps domain.StorePublicService, tx *sql.Tx) (err error) {
	query := `INSERT service_public SET general_information_id=?, purpose=?, facility=?, requirement=?, 
		tos=?, info_graphic=?, faq=?`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx,
		ps.GeneralInformation.ID,
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

	return
}
