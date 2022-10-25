package mysql

import (
	"fmt"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

/**
 * this block of code is used to generate the query for the news
 * need to be refactored later to be more generic and reduce the code complexity (go generic tech debt)
 */
func filterNewsQuery(params *domain.Request, binds *[]interface{}) string {
	var query string

	if params.Keyword != "" {
		*binds = append(*binds, `%`+params.Keyword+`%`)
		query = fmt.Sprintf(`%s AND n.title LIKE ?`, query)
	}

	if v, ok := params.Filters["created_by"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND n.created_by = ?`, query)
	}

	if v, ok := params.Filters["unit_id"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND u.unit_id = ?`, query)
	}

	if v, ok := params.Filters["highlight"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND n.highlight = ?`, query)
	}

	if v, ok := params.Filters["type"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND n.type = ?`, query)
	}

	if v, ok := params.Filters["is_live"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND n.is_live = ?`, query)
	}

	if v, ok := params.Filters["status"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND n.status = ?`, query)
	}

	if v, ok := params.Filters["exclude"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND n.id <> ?`, query)
	}

	if v, ok := params.Filters["tag"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND n.id IN (SELECT data_id FROM data_tags WHERE type = 'news' AND tag_name = ?)`, query)
	}

	if v, ok := params.Filters["category"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND n.category = ?`, query)
	} else if v, ok := params.Filters["categories"]; ok && v != "" {
		categories := params.Filters["categories"].([]string)
		if len(categories) > 0 {
			inBind := helpers.GetInBind(binds, categories)
			query = fmt.Sprintf(`%s AND n.category IN %s`, query, inBind)
		}
	}

	if params.StartDate != "" && params.EndDate != "" {
		*binds = append(*binds, params.StartDate, params.EndDate)
		query = fmt.Sprintf(`%s AND (DATE(n.updated_at) BETWEEN ? AND ?)`, query)
	}

	if v, ok := params.Filters["is_published_last_weekly"]; ok && v != "" {
		date := helpers.GetRangeLastWeek()
		*binds = append(*binds, date.DayOfLastWeek, date.Today)
		query = fmt.Sprintf(`%s AND (n.published_at >= ? AND n.published_at <= ?)`, query)
	}

	return query
}
