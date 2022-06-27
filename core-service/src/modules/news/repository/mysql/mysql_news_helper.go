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
func filterNewsQuery(params *domain.Request) string {
	var query string

	if params.Keyword != "" {
		query += ` AND n.title LIKE '%` + params.Keyword + `%' `
	}

	if v, ok := params.Filters["created_by"]; ok && v != "" {
		query = fmt.Sprintf(`%s AND n.created_by = '%v'`, query, v)
	}

	if v, ok := params.Filters["unit_id"]; ok && v != "" {
		query = fmt.Sprintf(`%s AND u.unit_id = '%v'`, query, v)
	}

	if v, ok := params.Filters["highlight"]; ok && v != "" {
		query = fmt.Sprintf(`%s AND n.highlight = '%s'`, query, v)
	}

	if v, ok := params.Filters["type"]; ok && v != "" {
		query = fmt.Sprintf(`%s AND n.type = "%s"`, query, v)
	}

	if v, ok := params.Filters["is_live"]; ok && v != "" {
		query = fmt.Sprintf(`%s AND n.is_live = "%s"`, query, v)
	}

	if v, ok := params.Filters["status"]; ok && v != "" {
		query = fmt.Sprintf(`%s AND n.status = "%s"`, query, v)
	}

	if v, ok := params.Filters["exclude"]; ok && v != "" {
		query = fmt.Sprintf(`%s AND n.id <> "%s"`, query, v)
	}

	if v, ok := params.Filters["tags"]; ok && v != "" {
		query = fmt.Sprintf(`%s AND n.id IN (SELECT data_id FROM data_tags WHERE tag_name = '%s')`, query, v)
	}

	if v, ok := params.Filters["category"]; ok && v != "" {
		query = fmt.Sprintf(`%s AND n.category = '%s'`, query, v)
	} else if v, ok := params.Filters["categories"]; ok && v != "" {
		categories := params.Filters["categories"].([]string)

		if len(categories) > 0 {
			query = fmt.Sprintf(`%s AND n.category IN ('%s')`, query, helpers.ConverSliceToString(categories, "','"))
		}
	}

	if params.StartDate != "" && params.EndDate != "" {
		query += ` AND DATE(n.updated_at) BETWEEN '` + params.StartDate + `' AND '` + params.EndDate + `'`
	}

	return query
}
