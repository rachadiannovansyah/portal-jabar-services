package mysql

import (
	"fmt"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

/**
 * this block of code is used to generate the query for the news
 * need to be refactored later to be more generic and reduce the code complexity (go generic tech debt)
 */
func filterMdsQuery(params *domain.Request, binds *[]interface{}) string {
	var query string

	if params.Keyword != "" {
		*binds = append(*binds, `%`+params.Keyword+`%`)
		query = fmt.Sprintf(`%s AND ms.service_name LIKE ?`, query)
	}

	if params.Filters["status"] != "" {
		*binds = append(*binds, params.Filters["status"])
		query = fmt.Sprintf(`%s AND status = ?`, query)
	}

	if v, ok := params.Filters["created_by"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND mds.created_by = ?`, query)
	}

	if v, ok := params.Filters["unit_id"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND u.unit_id = ?`, query)
	}

	return query
}
