package mysql

import (
	"context"
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type mysqlMdsRepository struct {
	Conn *sql.DB
}

// NewMysqlMasterDataServiceRepository will create an object that represent the MasterDataService.Repository interface
func NewMysqlMasterDataServiceRepository(Conn *sql.DB) domain.MasterDataServiceRepository {
	return &mysqlMdsRepository{Conn}
}

func (m *mysqlMdsRepository) Store(ctx context.Context, mds *domain.StoreMasterDataService, tx *sql.Tx) (err error) {
	query := `INSERT masterdata_services SET main_service=?, application=?, additional_information=?`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx,
		mds.Services.ID,
		mds.Application.ID,
		mds.AdditionalInformation.ID,
	)

	if err != nil {
		return
	}

	return
}

func (m *mysqlMdsRepository) GetTx(ctx context.Context) (tx *sql.Tx, err error) {
	tx, err = m.Conn.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	return
}
