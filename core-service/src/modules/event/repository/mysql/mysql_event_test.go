package mysql_test

import (
	"context"
	"testing"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mysqlRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/event/repository/mysql"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockEvent := []domain.Event{
		{
			ID:        1,
			Title:     "Pembagian Minyak Goreng Sapawarga",
			Priority:  5,
			Date:      time.Now(),
			StartHour: "12:00:00",
			EndHour:   "15:00:00",
			Type:      "Agenda Gubernur",
			Status:    "online",
			Category:  "category",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Title:     "Parade Jodoh Jawa Barat",
			Priority:  1,
			Date:      time.Now(),
			StartHour: "15:00:00",
			EndHour:   "17:00:00",
			Type:      "Agenda Wakil Gubernur",
			Status:    "online",
			Category:  "category",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "category", "title", "priority", "type", "status", "address", "url", "date", "start_hour", "end_hour", "created_by", "created_at", "updated_at"}).
		AddRow(mockEvent[0].ID, mockEvent[0].Category, mockEvent[0].Title, mockEvent[0].Priority, mockEvent[0].Type, mockEvent[0].Status,
			"", "", mockEvent[0].Date, mockEvent[0].StartHour, mockEvent[0].EndHour, "", mockEvent[0].CreatedAt, mockEvent[0].UpdatedAt).
		AddRow(mockEvent[1].ID, mockEvent[1].Category, mockEvent[1].Title, mockEvent[1].Priority, mockEvent[1].Type, mockEvent[1].Status,
			"", "", mockEvent[1].Date, mockEvent[1].StartHour, mockEvent[1].EndHour, "", mockEvent[1].CreatedAt, mockEvent[1].UpdatedAt)

	query := "SELECT e.id, e.category, e.title, e.priority, e.type, e.status, e.address, e.url, e.date, e.start_hour, e.end_hour, e.created_by, e.created_at, e.updated_at FROM events e LEFT JOIN users u ON e.created_by = u.id WHERE e.deleted_at is null"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := mysqlRepo.NewMysqlEventRepository(db)

	params := &domain.Request{
		Keyword:   "",
		PerPage:   10,
		Offset:    0,
		SortBy:    "",
		SortOrder: "",
	}

	list, _, err := a.Fetch(context.TODO(), params)
	assert.NoError(t, err)
	assert.Len(t, list, len(mockEvent))
}
