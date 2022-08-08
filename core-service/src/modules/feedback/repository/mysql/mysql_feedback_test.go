package mysql_test

import (
	"context"
	"testing"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mysqlRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/feedback/repository/mysql"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	body := domain.Feedback{
		// payload here
		Rating:      5,
		Compliments: "compliments of ..",
		Criticism:   "criticism of ..",
		Suggestions: "suggestions of ..",
		Sector:      "sector",
		CreatedAt:   time.Now(),
	}

	// query := "INSERT INTO users \\(id, name, email, phone\\) VALUES \\(\\?, \\?, \\?, \\?\\)"
	query := "INSERT feedback SET rating=? , compliments=? , criticism=?, suggestions=?, sector=?, created_at=?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(body.Rating,
		body.Compliments,
		body.Criticism,
		body.Suggestions,
		body.Sector,
		body.CreatedAt).WillReturnResult(sqlmock.NewResult(1, 1))

	e := mysqlRepo.NewMysqlFeedbackRepository(db)
	err = e.Store(context.TODO(), &body)
	assert.NotNil(t, err)
}
