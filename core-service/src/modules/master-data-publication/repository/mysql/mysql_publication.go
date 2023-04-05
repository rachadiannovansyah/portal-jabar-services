package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

type mysqlMdpRepository struct {
	Conn *sql.DB
}

// NewMysqlMasterDataPublicationRepository will create an object that represent the MasterDataPublication.Repository interface
func NewMysqlMasterDataPublicationRepository(Conn *sql.DB) domain.MasterDataPublicationRepository {
	return &mysqlMdpRepository{Conn}
}

func (m *mysqlMdpRepository) Store(ctx context.Context, body *domain.StoreMasterDataPublication) (err error) {
	query := `INSERT masterdata_publications SET mds_id=?, portal_category=?, slug=?, cover=?, images=?, infographics=?, keywords=?, faq=?, status=?, created_at=?, updated_at=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx,
		body.DefaultInformation.MdsID,
		body.DefaultInformation.PortalCategory,
		body.DefaultInformation.Slug,
		helpers.GetStringFromObject(body.ServiceDescription.Cover),
		helpers.GetStringFromObject(body.ServiceDescription.Images),
		helpers.GetStringFromObject(body.ServiceDescription.InfoGraphics),
		helpers.GetStringFromObject(body.AdditionalInformation.Keywords),
		helpers.GetStringFromObject(body.AdditionalInformation.FAQ),
		body.Status,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return
	}

	return
}

func (m *mysqlMdpRepository) GetTx(ctx context.Context) (tx *sql.Tx, err error) {
	tx, err = m.Conn.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	return
}
