package mysql

import (
	"context"
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type mysqlMailRepository struct {
	Conn *sql.DB
}

func NewMysqlMailRepository(Conn *sql.DB) domain.MailRepository {
	return &mysqlMailRepository{Conn}
}

func (m *mysqlMailRepository) fetchQuery(ctx context.Context, query string, args ...interface{}) (result []domain.Mail, err error) {
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

	result = make([]domain.Mail, 0)
	for rows.Next() {
		mail := domain.Mail{}
		err = rows.Scan(
			&mail.ID,
			&mail.From,
			&mail.To,
			&mail.Subject,
			&mail.CC,
			&mail.Body,
			&mail.Template,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, mail)
	}

	return result, nil
}

func (m *mysqlMailRepository) GetByTemplate(ctx context.Context, key string) (res domain.Mail, err error) {
	query := `SELECT * FROM mails WHERE template = ?`

	list, err := m.fetchQuery(ctx, query, key)
	if err != nil {
		return domain.Mail{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}
