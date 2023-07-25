package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type mysqlFeaturedProgramRepository struct {
	Conn *sql.DB
}

// NewMysqlFeaturedProgramRepository will create an object that represent the featuredProgram.Repository interface
func NewMysqlFeaturedProgramRepository(Conn *sql.DB) domain.FeaturedProgramRepository {
	return &mysqlFeaturedProgramRepository{Conn}
}

var querySelect = `SELECT id, title, excerpt, description, organization, categories, service_type, websites, social_media, logo, created_at, updated_at FROM featured_programs`

func getJSONSearch(params *domain.Request) (query string) {
	query = ` WHERE 1=1`

	categories := params.Filters["categories"].([]string)

	if len(categories) > 0 {
		for idx, cat := range categories {
			if idx == 0 {
				query = fmt.Sprintf(`%s AND (JSON_SEARCH(categories, 'all', '%s') IS NOT NULL`, query, cat)
			} else {
				query = fmt.Sprintf(`%s OR JSON_SEARCH(categories, 'all', '%s') IS NOT NULL`, query, cat)
			}
		}
		query += `)` // it's imagine end blocking query of loop json_search
	}

	return
}

func (m *mysqlFeaturedProgramRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.FeaturedProgram, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	result = make([]domain.FeaturedProgram, 0)
	for rows.Next() {
		fp := domain.FeaturedProgram{}
		err = rows.Scan(
			&fp.ID,
			&fp.Title,
			&fp.Excerpt,
			&fp.Description,
			&fp.Organization,
			&fp.Categories,
			&fp.ServiceType,
			&fp.Websites,
			&fp.SocialMedia,
			&fp.Logo,
			&fp.CreatedAt,
			&fp.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, fp)
	}

	return result, nil
}

func (m *mysqlFeaturedProgramRepository) count(_ context.Context, query string) (total int64, err error) {

	err = m.Conn.QueryRow(query).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}

func (m *mysqlFeaturedProgramRepository) getLastUpdated(_ context.Context, query string) (lastUpdated string, err error) {
	err = m.Conn.QueryRow(query).Scan(&lastUpdated)

	if err == sql.ErrNoRows {
		// there were no rows, but otherwise no error occurred
		log.Println(err)
	}

	return lastUpdated, nil
}

func (m *mysqlFeaturedProgramRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.FeaturedProgram, err error) {
	query := getJSONSearch(params)

	query = querySelect + query + ` LIMIT 50`

	res, err = m.fetch(ctx, query)

	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlFeaturedProgramRepository) MetaFetch(ctx context.Context, params *domain.Request) (total int64, lastUpdated string, err error) {
	query := getJSONSearch(params)

	total, err = m.count(ctx, ` SELECT COUNT(1) FROM featured_programs `)

	lastUpdated, err = m.getLastUpdated(ctx, ` SELECT updated_at FROM featured_programs`+query+` LIMIT 1`)

	if err != nil {
		return 0, "", err
	}

	return
}
