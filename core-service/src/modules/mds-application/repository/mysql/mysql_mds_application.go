package mysql

import (
	"context"
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

type mysqlApplicationRepository struct {
	Conn *sql.DB
}

// NewMysqlApplicationRepository will create an object that represent the Application.Repository interface
func NewMysqlApplicationRepository(Conn *sql.DB) domain.ApplicationRepository {
	return &mysqlApplicationRepository{Conn}
}

func (m *mysqlApplicationRepository) Store(ctx context.Context, ms *domain.StoreMasterDataService) (ID int64, err error) {
	query := `
	INSERT applications SET name=?, status=?, features=?
	`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx,
		&ms.Application.Name,
		&ms.Application.Status,
		helpers.GetStringFromObject(&ms.Application.Features),
	)
	if err != nil {
		return
	}
	ID, err = res.LastInsertId()
	if err != nil {
		return
	}

	return
}
