package mysql

import (
	"context"
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type mysqlAreaRepository struct {
	Conn *sql.DB
}

// NewMysqlAreaRepository will create an object that represent the area.Repository interface
func NewMysqlAreaRepository(Conn *sql.DB) domain.AreaRepository {
	return &mysqlAreaRepository{Conn}
}

func (m *mysqlAreaRepository) GetByID(ctx context.Context, id int64) (res domain.Area, err error) {
	query := `SELECT id, depth, name, parent_code_kemendagri, code_kemendagri, code_bps, latitude, longtitude, meta FROM categories WHERE id = ?`

	err = m.Conn.QueryRowContext(ctx, query, id).Scan(
		&res.ID,
		&res.Depth,
		&res.Name,
		&res.ParentCodeKemendagri,
		&res.CodeKemendagri,
		&res.CodeBps,
		&res.Latitude,
		&res.Longtitude,
		&res.Meta,
	)
	if err != nil {
		logrus.Error("Error found", err)
	}

	return
}
