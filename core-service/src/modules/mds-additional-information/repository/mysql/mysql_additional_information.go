package mysql

import (
	"context"
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

type mysqlAdditionalInformationRepository struct {
	Conn *sql.DB
}

// NewMysqlAdditionalInformationRepository will create an object that represent the GeneralInformation.Repository interface
func NewMysqlAdditionalInformationRepository(Conn *sql.DB) domain.AdditionalInformationRepository {
	return &mysqlAdditionalInformationRepository{Conn}
}

func (m *mysqlAdditionalInformationRepository) Store(ctx context.Context, ms *domain.StoreMasterDataService) (ID int64, err error) {
	query := `
	INSERT additional_informations SET responsible_name=?, phone_number=?, email=?, social_media=?
	`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx,
		&ms.AdditionalInformation.ResponsibleName,
		&ms.AdditionalInformation.PhoneNumber,
		&ms.AdditionalInformation.Email,
		helpers.GetStringFromObject(&ms.AdditionalInformation.SocialMedia),
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

func (m *mysqlAdditionalInformationRepository) Update(ctx context.Context, aID int64, ms *domain.StoreMasterDataService) (err error) {
	query := `
	UPDATE additional_informations SET responsible_name=?, phone_number=?, email=?, social_media=? WHERE id=?
	`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx,
		&ms.AdditionalInformation.ResponsibleName,
		&ms.AdditionalInformation.PhoneNumber,
		&ms.AdditionalInformation.Email,
		helpers.GetStringFromObject(&ms.AdditionalInformation.SocialMedia),
		aID,
	)
	if err != nil {
		return
	}

	return
}
