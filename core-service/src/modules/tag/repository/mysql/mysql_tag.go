package mysql

import (
	"context"
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type mysqlTagRepository struct {
	Conn *sql.DB
}

// NewMysqlTagRepository ..
func NewMysqlTagRepository(Conn *sql.DB) domain.TagRepository {
	return &mysqlTagRepository{Conn}
}

func (m *mysqlTagRepository) StoreTag(ctx context.Context, t *domain.Tag) (err error) {
	query := `INSERT tag SET name=?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, t.Name)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	t.ID = lastID

	return
}
