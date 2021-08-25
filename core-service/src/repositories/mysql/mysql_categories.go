package mysql

import (
	"context"
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type mysqlCategoriesRepository struct {
	Conn *sql.DB
}

// NewMysqlNewsCategoriesRepository will create an object that represent the news.NewsCategoryRepository interface
func NewMysqlCategoriesRepository(Conn *sql.DB) domain.CategoriesRepository {
	return &mysqlCategoriesRepository{Conn}
}

// GetByID ...
func (m *mysqlCategoriesRepository) GetByID(ctx context.Context, id int64) (res domain.Category, err error) {
	query := `SELECT id, title, description, type FROM categories WHERE id = ?`

	err = m.Conn.QueryRowContext(ctx, query, id).Scan(
		&res.ID,
		&res.Title,
		&res.Description,
		&res.Type,
	)
	if err != nil {
		logrus.Error("Error found", err)
	}

	return
}
