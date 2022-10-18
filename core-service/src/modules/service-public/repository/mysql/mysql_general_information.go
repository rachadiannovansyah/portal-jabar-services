package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type mysqlGeneralInformationRepository struct {
	Conn *sql.DB
}

// NewMysqlGeneralInformationRepository will create an object that represent the GeneralInformation.Repository interface
func NewMysqlGeneralInformationRepository(Conn *sql.DB) domain.GeneralInformationRepository {
	return &mysqlGeneralInformationRepository{Conn}
}

var querySelectGenInfo = `SELECT id, name, description, slug, category, address, unit, phone, logo, operational_hours, media, social_media, type FROM general_informations WHERE 1=1`

func (m *mysqlGeneralInformationRepository) fetchGenInfo(ctx context.Context, query string, args ...interface{}) (result []domain.GeneralInformation, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.GeneralInformation, 0)
	for rows.Next() {
		genInfo := domain.GeneralInformation{}
		err = rows.Scan(
			&genInfo.ID,
			&genInfo.Name,
			&genInfo.Description,
			&genInfo.Slug,
			&genInfo.Category,
			&genInfo.Address,
			&genInfo.Unit,
			&genInfo.Phone,
			&genInfo.Logo,
			&genInfo.OperationalHours,
			&genInfo.Media,
			&genInfo.SocialMedia,
			&genInfo.Type,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, genInfo)
	}

	return result, nil
}

func (m *mysqlGeneralInformationRepository) GetByID(ctx context.Context, id int64) (res domain.GeneralInformation, err error) {
	query := fmt.Sprintf("%s AND id=?", querySelectGenInfo)

	list, err := m.fetchGenInfo(ctx, query, id)
	if err != nil {
		return domain.GeneralInformation{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}
