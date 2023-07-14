package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type quickAccessRepository struct {
	Conn *sql.DB
}

var querySelect = `SELECT id, title, description, is_active, image, link, created_at, updated_at FROM quick_accesses WHERE 1=1 `
var querySelectTotal = `SELECT COUNT(1) FROM quick_accesses WHERE 1=1 `

func NewMysqlQuickAccessRepository(Conn *sql.DB) domain.QuickAccessRepository {
	return &quickAccessRepository{Conn}
}

func (m *quickAccessRepository) GetTx(ctx context.Context) (tx *sql.Tx, err error) {
	tx, err = m.Conn.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	return
}

func (m *quickAccessRepository) Store(ctx context.Context, body *domain.StoreQuickAccess, tx *sql.Tx) (err error) {
	query := `INSERT quick_accesses SET title=?, description=?, link=?, image=?, is_active=?, created_at=?, updated_at=?`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx,
		body.Title,
		body.Description,
		body.Link,
		body.Image,
		0,
		time.Now(),
		time.Now(),
	)

	return
}

func (m *quickAccessRepository) Update(ctx context.Context, ID int64, body *domain.StoreQuickAccess, tx *sql.Tx) (err error) {
	query := `UPDATE quick_accesses SET title=?, description=?, link=?, image=?, updated_at=? where id = ?`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx,
		body.Title,
		body.Description,
		body.Link,
		body.Image,
		time.Now(),
		ID,
	)

	return
}

func (m *quickAccessRepository) Fetch(ctx context.Context, params domain.Request) (result []domain.QuickAccess, total int64, err error) {
	binds := make([]interface{}, 0)
	queryFilter := filterQuickAccessQuery(params, &binds)

	querySelectTotal := querySelectTotal + queryFilter

	if err = m.Conn.QueryRowContext(ctx, querySelectTotal, binds...).Scan(&total); err != nil {
		return
	}

	binds = append(binds, params.Offset, params.PerPage)

	querySelect := querySelect + queryFilter + ` ORDER BY is_active DESC LIMIT ?,?`
	rows, err := m.Conn.QueryContext(ctx, querySelect, binds...)
	if err != nil {
		return
	}
	result = make([]domain.QuickAccess, 0)
	for rows.Next() {
		var item domain.QuickAccess
		err = rows.Scan(
			&item.ID,
			&item.Title,
			&item.Description,
			&item.IsActive,
			&item.Image,
			&item.Link,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			return
		}

		result = append(result, item)
	}
	return
}

func (m *quickAccessRepository) UpdateStatus(ctx context.Context, ID int64, body *domain.UpdateStatusQuickAccess, tx *sql.Tx) (err error) {
	query := `UPDATE quick_accesses SET is_active=?, updated_at=? WHERE id=?`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx,
		body.IsActive,
		time.Now(),
		ID,
	)
	return
}

func (m *quickAccessRepository) Delete(ctx context.Context, ID int64, tx *sql.Tx) (err error) {
	query := `DELETE FROM quick_accesses WHERE id = ?`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, ID)

	return
}

func (m *quickAccessRepository) GetByID(ctx context.Context, ID int64) (res domain.QuickAccess, err error) {
	query := querySelect + `AND id = ? LIMIT 1`

	err = m.Conn.QueryRowContext(ctx, query, ID).Scan(
		&res.ID,
		&res.Title,
		&res.Description,
		&res.IsActive,
		&res.Image,
		&res.Link,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if err != nil {
		err = domain.ErrNotFound
	}

	return
}
