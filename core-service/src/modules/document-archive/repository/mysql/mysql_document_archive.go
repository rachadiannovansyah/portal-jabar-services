package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type mysqlDocumentArchiveRepository struct {
	Conn *sql.DB
}

// NewMysqlDocumentArchiveRepository will create an object that represent the documentArchiveRepository interface
func NewMysqlDocumentArchiveRepository(Conn *sql.DB) domain.DocumentArchiveRepository {
	return &mysqlDocumentArchiveRepository{Conn}
}

var queryJoinDocArchive = `SELECT d.id, d.title, d.excerpt, d.description, d.source, d.mimetype, d.category,
	d.created_by, d.created_at, d.updated_at FROM document_archives d 
	LEFT JOIN users u
	ON d.created_by = u.id
	WHERE 1=1`

func (r *mysqlDocumentArchiveRepository) fetchQuery(ctx context.Context, query string, args ...interface{}) (result []domain.DocumentArchive, err error) {
	rows, err := r.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	result = make([]domain.DocumentArchive, 0)
	for rows.Next() {
		docArc := domain.DocumentArchive{}
		userID := uuid.UUID{}
		err = rows.Scan(
			&docArc.ID,
			&docArc.Title,
			&docArc.Excerpt,
			&docArc.Description,
			&docArc.Source,
			&docArc.Mimetype,
			&docArc.Category,
			&userID,
			&docArc.CreatedAt,
			&docArc.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		docArc.CreatedBy = domain.User{ID: userID}

		result = append(result, docArc)
	}

	return result, nil
}

func (r *mysqlDocumentArchiveRepository) count(ctx context.Context, query string) (total int64, err error) {

	err = r.Conn.QueryRow(query).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}

func (r *mysqlDocumentArchiveRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.DocumentArchive, total int64, err error) {
	query := filterDocArchiveQuery(params)

	if params.SortBy != "" {
		query += ` ORDER BY ` + params.SortBy + ` ` + params.SortOrder
	} else {
		query += ` ORDER BY d.created_at DESC `
	}

	total, _ = r.count(ctx, ` SELECT COUNT(1) FROM document_archives d LEFT JOIN users u ON d.created_by = u.id WHERE 1=1 `+query)
	query = queryJoinDocArchive + query + ` LIMIT ?,? `

	res, err = r.fetchQuery(ctx, query, params.Offset, params.PerPage)
	if err != nil {
		return nil, 0, err
	}

	return
}
