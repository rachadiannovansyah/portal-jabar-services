package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type mysqlPublicServiceRepository struct {
	Conn *sql.DB
}

// NewMysqlPublicServiceRepository will create an object that represent the publicService.Repository interface
func NewMysqlPublicServiceRepository(Conn *sql.DB) domain.PublicServiceRepository {
	return &mysqlPublicServiceRepository{Conn}
}

var querySelect = `SELECT id, name, description, unit, url, images, category, is_active, slug, excerpt, social_media, website, service_type, video, purposes, facilities, info, logo, created_at, updated_at FROM public_services WHERE 1=1`

func (m *mysqlPublicServiceRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.PublicService, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	result = make([]domain.PublicService, 0)
	for rows.Next() {
		ps := domain.PublicService{}
		err = rows.Scan(
			&ps.ID,
			&ps.Name,
			&ps.Description,
			&ps.Unit,
			&ps.Url,
			&ps.Images,
			&ps.Category,
			&ps.IsActive,
			&ps.Slug,
			&ps.Excerpt,
			&ps.SocialMedia,
			&ps.Website,
			&ps.ServiceType,
			&ps.Video,
			&ps.Purposes,
			&ps.Facilities,
			&ps.Info,
			&ps.Logo,
			&ps.CreatedAt,
			&ps.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, ps)
	}

	return result, nil
}

func (m *mysqlPublicServiceRepository) count(ctx context.Context, query string, args ...interface{}) (total int64, err error) {

	err = m.Conn.QueryRow(query, args...).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}

func (m *mysqlPublicServiceRepository) getLastUpdated(ctx context.Context, query string) (lastUpdated string, err error) {
	err = m.Conn.QueryRow(query).Scan(&lastUpdated)

	if err == sql.ErrNoRows {
		// there were no rows, but otherwise no error occurred
		log.Println(err)
	}

	return lastUpdated, nil
}

func (m *mysqlPublicServiceRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.PublicService, err error) {
	queryFilter := filterPublicServiceQuery(params)
	query := querySelect + queryFilter + ` LIMIT ?,? `

	res, err = m.fetch(ctx, query, params.Offset, params.PerPage)
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlPublicServiceRepository) MetaFetch(ctx context.Context, params *domain.Request) (total int64, lastUpdated string, err error) {
	query := ` SELECT COUNT(1) FROM public_services WHERE category = ? ` // default list public service using filter by category
	total, _ = m.count(ctx, query, params.Filters["category"])
	lastUpdated, err = m.getLastUpdated(ctx, ` SELECT updated_at FROM public_services ORDER BY updated_at DESC LIMIT 1`)

	if err != nil {
		return 0, "", err
	}

	return
}

func (m *mysqlPublicServiceRepository) GetBySlug(ctx context.Context, slug string) (res domain.PublicService, err error) {
	return m.findOne(ctx, "slug", slug)
}

func (m *mysqlPublicServiceRepository) findOne(ctx context.Context, key string, value string) (res domain.PublicService, err error) {
	query := fmt.Sprintf("%s AND %s=?", querySelect, key)

	list, err := m.fetch(ctx, query, value)
	if err != nil {
		return domain.PublicService{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}
