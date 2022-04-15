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
func filterEventQuery(params *domain.Request) string {
	var query string

	if params.Keyword != "" {
		query += ` AND e.title LIKE '%` + params.Keyword + `%' `
	}

	if v, ok := params.Filters["unit_id"]; ok && v != "" {
		query = fmt.Sprintf(`%s AND u.unit_id = '%v'`, query, v)
	}

	if params.StartDate != "" && params.EndDate != "" {
		query += ` AND e.date BETWEEN '` + params.StartDate + `' AND '` + params.EndDate + `'`
	}

	if v, ok := params.Filters["type"]; ok && v != "" {
		types := params.Filters["type"].([]string)
		if len(types) > 0 {
			query = fmt.Sprintf(`%s AND e.type IN ('%s')`, query, helpers.ConverSliceToString(types, "','"))
		}
	}

	if v, ok := params.Filters["categories"]; ok && v != "" {
		categories := params.Filters["categories"].([]string)
		if len(categories) > 0 {
			query = fmt.Sprintf(`%s AND e.category IN ('%s')`, query, helpers.ConverSliceToString(categories, "','"))
		}
	}

	return query
}
