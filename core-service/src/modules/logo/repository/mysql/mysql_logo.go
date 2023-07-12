package mysql

import (
	"context"
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type logoRepository struct {
	Conn *sql.DB
}

var querySelect = `SELECT id, title, image FROM logos WHERE 1=1 `
var querySelectTotal = `SELECT COUNT(1) FROM logos WHERE 1=1 `

func NewMysqlLogoRepository(Conn *sql.DB) domain.LogoRepository {
	return &logoRepository{Conn}
}

func (m *logoRepository) Fetch(ctx context.Context, params domain.Request) (result []domain.Logo, total int64, err error) {
	binds := make([]interface{}, 0)
	queryFilter := filterLogoQuery(params, &binds)

	querySelectTotal := querySelectTotal + queryFilter

	if err = m.Conn.QueryRowContext(ctx, querySelectTotal, binds...).Scan(&total); err != nil {
		return
	}

	binds = append(binds, params.Offset, params.PerPage)

	querySelect := querySelect + queryFilter + ` ORDER BY title ASC LIMIT ?,?`
	rows, err := m.Conn.QueryContext(ctx, querySelect, binds...)
	if err != nil {
		return
	}
	result = make([]domain.Logo, 0)
	for rows.Next() {
		var item domain.Logo
		err = rows.Scan(
			&item.ID,
			&item.Title,
			&item.Image,
		)
		if err != nil {
			return
		}

		result = append(result, item)
	}
	return
}

func (m *logoRepository) Store(ctx context.Context, body *domain.StoreLogoRequest) (err error) {
	query := `INSERT logos SET title=?, image=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx,
		body.Title,
		body.Image,
	)

	return
}
