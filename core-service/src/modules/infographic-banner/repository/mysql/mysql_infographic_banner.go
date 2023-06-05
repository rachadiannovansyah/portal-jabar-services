package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

type infographicBannerRepository struct {
	Conn *sql.DB
}

var querySelect = `SELECT id, title, sequence, link, is_active, image, created_at, updated_at FROM infographic_banners WHERE 1=1 `
var querySelectTotal = `SELECT COUNT(1) FROM infographic_banners WHERE 1=1 `

func NewMysqlInfographicBannerRepository(Conn *sql.DB) domain.InfographicBannerRepository {
	return &infographicBannerRepository{Conn}
}

func (m *infographicBannerRepository) GetTx(ctx context.Context) (tx *sql.Tx, err error) {
	tx, err = m.Conn.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	return
}

func (m *infographicBannerRepository) Store(ctx context.Context, body *domain.StoreInfographicBanner, tx *sql.Tx) (err error) {
	query := `INSERT infographic_banners SET title=?, sequence=?, link=?, image=?`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx,
		body.Title,
		1,
		body.Link,
		helpers.GetStringFromObject(body.Image),
	)

	return
}

func (m *infographicBannerRepository) Update(ctx context.Context, ID int64, body *domain.StoreInfographicBanner, tx *sql.Tx) (err error) {
	query := `UPDATE infographic_banners SET title=?, link=?, image=?, updated_at=? where id = ?`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx,
		body.Title,
		body.Link,
		helpers.GetStringFromObject(body.Image),
		time.Now(),
		ID,
	)

	return
}

func (m *infographicBannerRepository) GetLastSequence(ctx context.Context) (sequence int64) {
	query := `SELECT sequence FROM infographic_banners WHERE is_active = 1 order by sequence DESC`

	_ = m.Conn.QueryRowContext(ctx, query).Scan(&sequence)

	return
}

func (m *infographicBannerRepository) SyncSequence(ctx context.Context, sequence int64, tx *sql.Tx) (err error) {
	query := `SELECT id FROM infographic_banners WHERE is_active = 1 order by sequence ASC`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return
	}
	items := make([]domain.SyncSequence, 0)
	for rows.Next() {
		var ID int64
		err = rows.Scan(&ID)
		if err != nil {
			return
		}

		items = append(items, domain.SyncSequence{
			ID:       ID,
			Sequence: int8(sequence),
		})

		sequence++
	}

	for _, item := range items {
		err = m.UpdateSequence(ctx, item.ID, item.Sequence, tx)
		if err != nil {
			return
		}
	}
	return
}

func (m *infographicBannerRepository) Fetch(ctx context.Context, params domain.Request) (result []domain.InfographicBanner, total int64, err error) {
	binds := make([]interface{}, 0)
	queryFilter := filterInfographicBannerQuery(params, &binds)

	querySelectTotal := querySelectTotal + queryFilter

	if err = m.Conn.QueryRowContext(ctx, querySelectTotal, binds...).Scan(&total); err != nil {
		return
	}

	binds = append(binds, params.Offset, params.PerPage)

	querySelect := querySelect + queryFilter + ` ORDER BY is_active DESC, sequence ASC  LIMIT ?,?`
	rows, err := m.Conn.QueryContext(ctx, querySelect, binds...)
	if err != nil {
		return
	}
	result = make([]domain.InfographicBanner, 0)
	for rows.Next() {
		var item domain.InfographicBanner
		err = rows.Scan(
			&item.ID,
			&item.Title,
			&item.Sequence,
			&item.Link,
			&item.IsActive,
			&item.Image,
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

func (m *infographicBannerRepository) UpdateSequence(ctx context.Context, ID int64, sequence int8, tx *sql.Tx) (err error) {
	query := `UPDATE infographic_banners SET sequence=? WHERE id=?`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx,
		sequence,
		ID,
	)
	return
}

func (m *infographicBannerRepository) UpdateStatus(ctx context.Context, ID int64, body *domain.UpdateStatusInfographicBanner, tx *sql.Tx) (err error) {
	query := `UPDATE infographic_banners SET is_active=?, sequence=? WHERE id=?`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	sequence := body.IsActive

	_, err = stmt.ExecContext(ctx,
		body.IsActive,
		sequence,
		ID,
	)
	return
}

func (m *infographicBannerRepository) Delete(ctx context.Context, ID int64, tx *sql.Tx) (err error) {
	query := `DELETE FROM infographic_banners WHERE id = ?`

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, ID)

	return
}

func (m *infographicBannerRepository) GetByID(ctx context.Context, ID int64, tx *sql.Tx) (res domain.InfographicBanner, err error) {
	query := querySelect + `AND id = ? LIMIT 1`

	err = tx.QueryRowContext(ctx, query, ID).Scan(
		&res.ID,
		&res.Title,
		&res.Sequence,
		&res.Link,
		&res.IsActive,
		&res.Image,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if err != nil {
		err = errors.New(fmt.Sprintf("Your requested Item with ID %d is not found", ID))
	}

	return
}
