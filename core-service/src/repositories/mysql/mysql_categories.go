package mysql

import (
	"context"
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type mysqlNewsCategoriesRepository struct {
	Conn *sql.DB
}

// NewMysqlNewsCategoriesRepository will create an object that represent the news.NewsCategoryRepository interface
func NewMysqlCategoriesRepository(Conn *sql.DB) domain.CategoriesRepository {
	return &mysqlNewsCategoriesRepository{Conn}
}

// GetByID ...
func (m *mysqlNewsCategoriesRepository) GetByID(ctx context.Context, id int64) (res domain.Categories, err error) {
	query := `SELECT id, title, description FROM news_categories WHERE id = ?`

	err = m.Conn.QueryRowContext(ctx, query, id).Scan(
		&res.ID,
		&res.Title,
		&res.Description,
	)
	if err != nil {
		logrus.Error("wkwk", err)
	}

	return
}
