package mysql

import (
	"context"
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type mysqlMailTemplateRepository struct {
	Conn *sql.DB
}

func NewMysqlMailTemplateRepository(Conn *sql.DB) domain.TemplateRepository {
	return &mysqlMailTemplateRepository{Conn}
}

func (m *mysqlMailTemplateRepository) fetchQuery(ctx context.Context, query string, args ...interface{}) (result []domain.Template, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	result = make([]domain.Template, 0)
	for rows.Next() {
		mail := domain.Template{}
		err = rows.Scan(
			&mail.ID,
			&mail.From,
			&mail.To,
			&mail.Subject,
			&mail.CC,
			&mail.Body,
			&mail.Key,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, mail)
	}

	return result, nil
}

func (m *mysqlMailTemplateRepository) GetByTemplate(ctx context.Context, key string) (res domain.Template, err error) {
	query := `SELECT * FROM mails WHERE template = ?`

	list, err := m.fetchQuery(ctx, query, key)
	if err != nil {
		return domain.Template{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}
