package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/sirupsen/logrus"
)

type mysqlNewsRepository struct {
	Conn *sql.DB
}

// NewMysqlNewsRepository will create an object that represent the news.Repository interface
func NewMysqlNewsRepository(Conn *sql.DB) domain.NewsRepository {
	return &mysqlNewsRepository{Conn}
}

var querySelectNews = `SELECT id, category, title, excerpt, content, image, video, slug, author, reporter, editor, area_id, type, 
	views, shared, source, duration, start_date, end_date, status, is_live, published_at, created_by, created_at, updated_at FROM news WHERE deleted_at is null`

var queryJoinNews = `SELECT n.id, n.category, n.title, n.excerpt, n.content, n.image, n.video, n.slug, n.author, n.reporter, n.editor, n.area_id, n.type, 
	n.views, n.shared, n.source, n.duration, n.start_date, n.end_date, n.status, n.is_live, n.published_at, n.created_by, n.created_at, n.updated_at FROM news n
	LEFT JOIN users u
	ON n.created_by = u.id
	WHERE n.deleted_at is NULL`

func (m *mysqlNewsRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.News, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.News, 0)
	for rows.Next() {
		t := domain.News{}
		createdByID := uuid.UUID{}
		areaID := int64(0)
		err = rows.Scan(
			&t.ID,
			&t.Category,
			&t.Title,
			&t.Excerpt,
			&t.Content,
			&t.Image,
			&t.Video,
			&t.Slug,
			&t.Author,
			&t.Reporter,
			&t.Editor,
			&areaID,
			&t.Type,
			&t.Views,
			&t.Shared,
			&t.Source,
			&t.Duration,
			&t.StartDate,
			&t.EndDate,
			&t.Status,
			&t.IsLive,
			&t.PublishedAt,
			&createdByID,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		t.CreatedBy = domain.User{ID: createdByID}
		t.Area = domain.Area{ID: areaID}

		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlNewsRepository) findOne(ctx context.Context, key string, value string) (res domain.News, err error) {
	query := fmt.Sprintf("%s AND %s=?", querySelectNews, key)

	list, err := m.fetch(ctx, query, value)
	if err != nil {
		return domain.News{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlNewsRepository) TabStatus(ctx context.Context, id uuid.UUID, key string) (res []domain.TabStatusResponse, err error) {
	var query string
	var list []domain.TabStatusResponse

	if key == "contributor" {
		query = "SELECT status, COUNT(status) FROM news WHERE created_by = ? GROUP BY status"
		list, err = m.fetchTabs(ctx, query, id)
	} else {
		query = "SELECT status, COUNT(status) FROM news GROUP BY status"
		list, err = m.fetchTabs(ctx, query)
	}

	if err != nil {
		return []domain.TabStatusResponse{}, err
	}

	if len(list) > 0 {
		res = list
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlNewsRepository) fetchTabs(ctx context.Context, query string, args ...interface{}) (result []domain.TabStatusResponse, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.TabStatusResponse, 0)
	for rows.Next() {
		t := domain.TabStatusResponse{}
		err = rows.Scan(
			&t.Status,
			&t.Count,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlNewsRepository) count(ctx context.Context, query string) (total int64, err error) {

	err = m.Conn.QueryRow(query).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}

func (m *mysqlNewsRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.News, total int64, err error) {
	query := buildQueryFetchNews(params)
	total, _ = m.count(ctx, ` SELECT COUNT(1) FROM news n LEFT JOIN users u ON n.created_by = u.id WHERE n.deleted_at is NULL `+query)
	query = queryJoinNews + query + ` LIMIT ?,? `
	res, err = m.fetch(ctx, query, params.Offset, params.PerPage)

	if err != nil {
		return nil, 0, err
	}

	return
}

func (m *mysqlNewsRepository) GetByID(ctx context.Context, id int64) (res domain.News, err error) {
	return m.findOne(ctx, "id", fmt.Sprintf("%v", id))
}

func (m *mysqlNewsRepository) GetBySlug(ctx context.Context, slug string) (res domain.News, err error) {
	return m.findOne(ctx, "slug", slug)
}

func (m *mysqlNewsRepository) AddView(ctx context.Context, id int64) (err error) {
	query := `UPDATE news SET views = views + 1 WHERE id = ?`

	_, err = m.Conn.ExecContext(ctx, query, id)

	return
}

func (m *mysqlNewsRepository) AddShare(ctx context.Context, id int64) (err error) {
	query := `UPDATE news SET shared = shared + 1 WHERE id = ?`

	_, err = m.Conn.ExecContext(ctx, query, id)

	return
}

func (m *mysqlNewsRepository) FetchNewsBanner(ctx context.Context) (res []domain.News, err error) {
	query := querySelectNews + ` AND id IN (
		SELECT MAX(id) FROM news WHERE highlight = ? and is_live=1 GROUP BY category 
	)`

	res, err = m.fetch(ctx, query, 1)
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlNewsRepository) FetchNewsHeadline(ctx context.Context) (res []domain.News, err error) {
	query := querySelectNews + ` AND id IN (
		SELECT MAX(id) FROM news WHERE id NOT IN (
			SELECT id from news  WHERE id IN (
				SELECT MAX(id) FROM news WHERE highlight = 1 and is_live=1
				GROUP BY category 
			)
		) AND highlight = 1 and is_live=1 GROUP BY category
	)`

	res, err = m.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlNewsRepository) Store(ctx context.Context, n *domain.StoreNewsRequest) (err error) {
	query := `INSERT news SET title=?, excerpt=?, content=?, slug=?, image=?, category=?,
		source=?, status=?, type=?, duration=?, start_date=?, end_date=?, area_id=?, author=?, reporter=?, editor=?, created_by=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	// temporary unique slug
	slug := helpers.MakeSlug(n.Title, time.Now().Unix())

	res, err := stmt.ExecContext(ctx,
		n.Title,
		n.Excerpt,
		n.Content,
		slug,
		n.Image,
		n.Category,
		n.Source,
		n.Status,
		"article",
		n.Duration,
		n.StartDate,
		n.EndDate,
		n.AreaID,
		n.Author,
		n.Reporter,
		n.Editor,
		n.CreatedBy.ID.String(),
	)
	if err != nil {
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	n.ID = lastID
	return
}

func (m *mysqlNewsRepository) Update(ctx context.Context, id int64, n *domain.StoreNewsRequest) (err error) {
	query := `UPDATE news SET title=?, excerpt=?, content=?, image=?, category=?, slug=?, author=?, reporter=?, editor=?,
		source=?, status=?, type=?, duration=?, start_date=?, end_date=?, area_id=?, is_live=?, published_at=?, updated_by=?, updated_at=? WHERE id=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx,
		n.Title,
		n.Excerpt,
		n.Content,
		n.Image,
		n.Category,
		n.Slug,
		n.Author,
		n.Reporter,
		n.Editor,
		n.Source,
		n.Status,
		"article",
		n.Duration,
		n.StartDate,
		n.EndDate,
		n.AreaID,
		n.IsLive,
		n.PublishedAt,
		n.CreatedBy.ID,
		time.Now(),
		id,
	)
	if err != nil {
		return
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return
	}

	if affect != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", affect)
		return
	}

	return
}

func (m *mysqlNewsRepository) Delete(ctx context.Context, id int64) (err error) {
	query := "DELETE FROM news WHERE id = ? AND status = ?"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id, "DRAFT")
	if err != nil {
		return
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowAffected != 1 {
		err = fmt.Errorf("Weird Behavior. Total Affected: %d", rowAffected)
		return
	}

	return
}
