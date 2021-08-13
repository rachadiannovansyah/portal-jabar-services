package mysql_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/jabardigitalservice/portal-jabar-api/domain"
	"github.com/jabardigitalservice/portal-jabar-api/news/repository"
	newsMysqlRepo "github.com/jabardigitalservice/portal-jabar-api/news/repository/mysql"
)

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockNews := []domain.News{
		domain.News{
			ID: 1, Title: "title 1", Content: "content 1",
			UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
		domain.News{
			ID: 2, Title: "title 2", Content: "content 2",
			UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "title", "content", "updated_at", "createdAt"}).
		AddRow(mockNews[0].ID, mockNews[0].Title, mockNews[0].Content,
			mockNews[0].UpdatedAt, mockNews[0].CreatedAt).
		AddRow(mockNews[1].ID, mockNews[1].Title, mockNews[1].Content,
			mockNews[1].UpdatedAt, mockNews[1].CreatedAt)

	query := "SELECT id,title,content, updated_at, createdAt FROM contents WHERE createdAt > \\? ORDER BY createdAt LIMIT \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := newsMysqlRepo.NewMysqlNewsRepository(db)
	cursor := repository.EncodeCursor(mockNews[1].CreatedAt)
	num := int64(2)
	list, nextCursor, err := a.Fetch(context.TODO(), cursor, num)
	assert.NotEmpty(t, nextCursor)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "title", "content", "updated_at", "createdAt"}).
		AddRow(1, "title 1", "News 1", 1, time.Now(), time.Now())

	query := "SELECT id,title,content, updated_at, createdAt FROM contents WHERE ID = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := newsMysqlRepo.NewMysqlNewsRepository(db)

	num := int64(5)
	aContents, err := a.GetByID(context.TODO(), num)
	assert.NoError(t, err)
	assert.NotNil(t, aContents)
}

func TestStore(t *testing.T) {
	now := time.Now()
	ar := &domain.News{
		Title:     "Judul",
		Content:   "News",
		CreatedAt: now,
		UpdatedAt: now,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "INSERT contents SET title=\\? , content=\\?, updated_at=\\? , createdAt=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.Title, ar.Content, ar.CreatedAt, ar.UpdatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

	a := newsMysqlRepo.NewMysqlNewsRepository(db)

	err = a.Store(context.TODO(), ar)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), ar.ID)
}

func TestGetBySlug(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "title", "content", "updated_at", "createdAt"}).
		AddRow(1, "title 1", "News 1", 1, time.Now(), time.Now())

	query := "SELECT id,title,content, updated_at, createdAt FROM contents WHERE slug = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := newsMysqlRepo.NewMysqlNewsRepository(db)

	slug := "slug-1"
	aContents, err := a.GetBySlug(context.TODO(), slug)
	assert.NoError(t, err)
	assert.NotNil(t, aContents)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "DELETE FROM contents WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(12).WillReturnResult(sqlmock.NewResult(12, 1))

	a := newsMysqlRepo.NewMysqlNewsRepository(db)

	num := int64(12)
	err = a.Delete(context.TODO(), num)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	now := time.Now()
	ar := &domain.News{
		ID:        12,
		Title:     "Judul",
		Content:   "News",
		CreatedAt: now,
		UpdatedAt: now,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "UPDATE contents set title=\\?, content=\\?, updated_at=\\? WHERE ID = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.Title, ar.Content, ar.UpdatedAt, ar.ID).WillReturnResult(sqlmock.NewResult(12, 1))

	a := newsMysqlRepo.NewMysqlNewsRepository(db)

	err = a.Update(context.TODO(), ar)
	assert.NoError(t, err)
}
