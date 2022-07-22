package mysql_test

import (
	"context"
	"testing"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mysqlRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/news/repository/mysql"
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
			Title:     "title",
			Excerpt:   "excerpt",
			Content:   "content",
			Views:     10,
			Shared:    25,
			Image:     nil,
			StartDate: time.Now(),
			EndDate:   time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Title:     "title",
			Excerpt:   "excerpt",
			Content:   "content",
			Views:     15,
			Shared:    30,
			Image:     nil,
			StartDate: time.Now(),
			EndDate:   time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "category", "title", "excerpt", "content", "image", "video", "slug", "author", "reporter", "editor", "area_id", "type", "views", "shared", "source", "duration", "start_date", "end_date", "status", "is_live", "published_at", "created_by", "created_at", "updated_at"}).
		AddRow(mockNews[0].ID, mockNews[0].Category, mockNews[0].Title, mockNews[0].Excerpt, mockNews[0].Content,
			nil, nil, mockNews[0].Slug, "", "", "", 0, "", mockNews[0].Views, mockNews[0].Shared, mockNews[0].Source, 0, mockNews[0].StartDate, mockNews[0].EndDate, mockNews[0].Status, 0, mockNews[0].PublishedAt, nil, mockNews[0].CreatedAt, mockNews[0].UpdatedAt).
		AddRow(mockNews[1].ID, mockNews[1].Category, mockNews[1].Title, mockNews[1].Excerpt, mockNews[1].Content,
			nil, nil, mockNews[1].Slug, "", "", "", 0, "", mockNews[1].Views, mockNews[1].Shared, mockNews[1].Source, 0, mockNews[1].StartDate, mockNews[1].EndDate, mockNews[1].Status, 0, mockNews[1].PublishedAt, nil, mockNews[1].CreatedAt, mockNews[1].UpdatedAt)

	query := "SELECT n.id, n.category, n.title, n.excerpt, n.content, n.image, n.video, n.slug, n.author, n.reporter, n.editor, n.area_id, n.type, n.views, n.shared, n.source, n.duration, n.start_date, n.end_date, n.status, n.is_live, n.published_at, n.created_by, n.created_at, n.updated_at FROM news n LEFT JOIN users u ON n.created_by = u.id WHERE n.deleted_at is NULL"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := mysqlRepo.NewMysqlNewsRepository(db)

	params := &domain.Request{
		Keyword:   "",
		PerPage:   10,
		Offset:    0,
		SortBy:    "",
		SortOrder: "",
	}

	list, _, err := a.Fetch(context.TODO(), params)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}
