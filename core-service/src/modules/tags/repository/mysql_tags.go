package mysql

import (
	"context"
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type mysqlTagsRepository struct {
	Conn *sql.DB
}

// NewMysqlTagsRepository ..
func NewMysqlTagsRepository(Conn *sql.DB) domain.TagsRepository {
	return &mysqlTagsRepository{Conn}
}

func (m *mysqlTagsRepository) StoreTags(ctx context.Context, t *domain.Tags) (err error) {
	query := `INSERT tags SET name=?`

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
