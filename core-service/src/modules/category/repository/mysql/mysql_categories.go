package mysql

import (
	"context"
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type mysqlCategoryRepository struct {
	Conn *sql.DB
}

// NewMysqlCategoryRepository will create an object that represent the news.NewsCategoryRepository interface
func NewMysqlCategoryRepository(Conn *sql.DB) domain.CategoryRepository {
	return &mysqlCategoryRepository{Conn}
}

// GetByID ...
func (m *mysqlCategoryRepository) GetByID(ctx context.Context, id int64) (res domain.Category, err error) {
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
