package test

import (
	"context"
	"testing"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mysqlRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/repositories/mysql"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockNews := []domain.News{
		{
			ID:        1,
			Title:     domain.NullString{String: "title", Valid: true},
			Excerpt:   domain.NullString{String: "excerpt", Valid: true},
			Content:   domain.NullString{String: "content", Valid: true},
			Image:     domain.NullString{String: "image", Valid: true},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Title:     domain.NullString{String: "title 2", Valid: true},
			Excerpt:   domain.NullString{String: "excerpt 2", Valid: true},
			Content:   domain.NullString{String: "content 2", Valid: true},
			Image:     domain.NullString{String: "image 2", Valid: true},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "category_id", "title", "excerpt", "content", "image", "video", "slug", "created_at", "updated_at"}).
		AddRow(mockNews[0].ID, mockNews[0].Category.ID, mockNews[0].Title.String, mockNews[0].Excerpt.String, mockNews[0].Content.String,
			nil, nil, mockNews[0].Slug.String, mockNews[0].CreatedAt, mockNews[0].UpdatedAt).
		AddRow(mockNews[1].ID, mockNews[1].Category.ID, mockNews[1].Title.String, mockNews[1].Excerpt.String, mockNews[1].Content.String,
			nil, nil, mockNews[1].Slug.String, mockNews[1].CreatedAt, mockNews[1].UpdatedAt)

	query := "SELECT id, category_id, title, excerpt, content, image, video, slug, created_at, updated_at FROM news"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := mysqlRepo.NewMysqlNewsRepository(db)

	params := &domain.Request{
		Keyword: "",
		PerPage: 10,
		Offset:  0,
		OrderBy: "",
		SortBy:  "",
	}

	list, _, err := a.Fetch(context.TODO(), params)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}
