package mysql

import (
	"context"
	"database/sql"
	"fmt"

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

func (m *mysqlFeaturedProgramRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.FeaturedProgram, err error) {
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
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, fp)
	}

	return result, nil
}

func (m *mysqlFeaturedProgramRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.FeaturedProgram, err error) {

	query := `SELECT id, title, excerpt, description, organization, categories, service_type, websites, social_media, logo FROM featured_programs WHERE 1=1`

	if v, ok := params.Filters["categories"]; ok && v != "" {
		query = fmt.Sprintf(`%s AND json_search(categories, 'all', '%s') is not null limit 50`, query, v)
	}

	res, err = m.fetch(ctx, query)

	if err != nil {
		return nil, err
	}

	return
}
