package mysql

import (
	"context"
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type mysqlFeedbackRepository struct {
	Conn *sql.DB
}

// NewMysqlFeedbackRepository creates a new mysql feedback repository
func NewMysqlFeedbackRepository(Conn *sql.DB) domain.FeedbackRepository {
	return &mysqlFeedbackRepository{Conn}
}

func (m *mysqlFeedbackRepository) Store(ctx context.Context, f *domain.Feedback) (err error) {
	query := `INSERT feedback SET rating=? , compliments=? , criticism=?, suggestions=?, sector=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, f.Rating, f.Compliments, f.Criticism, f.Suggestions, f.Sector)
	if err != nil {
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	f.ID = lastID
	return
}
