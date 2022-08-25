package mysql

import (
	"fmt"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

/**
 * this block of code is used to generate the query for the event
 * need to be refactored later to be more generic and reduce the code complexity (go generic tech debt)
 */
func filterDocArchiveQuery(params *domain.Request) string {
	var query string

	if params.Keyword != "" {
		query += ` AND d.title LIKE '%` + params.Keyword + `%' `
	}

	if v, ok := params.Filters["category"]; ok && v != "" {
		query = fmt.Sprintf(`%s AND d.category = '%s'`, query, v)
	}

	return query
}
