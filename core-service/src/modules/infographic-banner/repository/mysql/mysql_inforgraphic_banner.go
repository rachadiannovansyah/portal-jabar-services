package mysql

import (
	"context"
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

type infographicBannerRepository struct {
	Conn *sql.DB
}

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

func (m *infographicBannerRepository) GetLastSequence(ctx context.Context) (sequence int64) {
	query := `SELECT sequence FROM infographic_banners where is_active = 1 order by sequence DESC`

	_ = m.Conn.QueryRowContext(ctx, query).Scan(&sequence)

	return
}

func (m *infographicBannerRepository) SyncSequence(ctx context.Context, sequence int64, tx *sql.Tx) (err error) {
	query := `SELECT id FROM infographic_banners where is_active = 1 order by sequence ASC`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return
	}
	for rows.Next() {
		var ID int64
		err = rows.Scan(&ID)
		if err != nil {
			return
		}

		err = m.UpdateSequence(ctx, ID, int8(sequence), tx)
		if err != nil {
			return
		}
		sequence++
	}
	return
}

func (m *infographicBannerRepository) UpdateSequence(ctx context.Context, ID int64, sequence int8, tx *sql.Tx) (err error) {
	query := `UPDATE infographic_banners SET sequence=? where id=?`

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
