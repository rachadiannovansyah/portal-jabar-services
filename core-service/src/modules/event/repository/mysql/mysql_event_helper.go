package mysql

import (
	"fmt"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

/**
 * this block of code is used to generate the query for the event
 * need to be refactored later to be more generic and reduce the code complexity (go generic tech debt)
 */
func filterEventQuery(params *domain.Request, binds *[]interface{}) string {
	var query string

	if params.Keyword != "" {
		*binds = append(*binds, `%`+params.Keyword+`%`)
		query += ` AND e.title LIKE ? `
	}

	if v, ok := params.Filters["unit_id"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND u.unit_id = ?`, query)
	}

	if params.StartDate != "" && params.EndDate != "" {
		*binds = append(*binds, params.StartDate, params.EndDate)
		query += ` AND (DATE(e.date) BETWEEN ? AND ?) `
	}

	if v, ok := params.Filters["type"]; ok && v != "" {
		types := params.Filters["type"].([]string)
		if len(types) > 0 {
			inBind := helpers.GetInBind(binds, types)
			query = fmt.Sprintf(`%s AND e.type IN %s`, query, inBind)
		}
	}

	if v, ok := params.Filters["categories"]; ok && v != "" {
		categories := params.Filters["categories"].([]string)
		if len(categories) > 0 {
			inBind := helpers.GetInBind(binds, categories)
			query = fmt.Sprintf(`%s AND e.category IN %s`, query, inBind)
		}
	}

	return query
}
