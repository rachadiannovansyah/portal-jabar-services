package mysql

import (
	"fmt"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

/**
 * this block of code is used to generate the query for the news
 * need to be refactored later to be more generic and reduce the code complexity (go generic tech debt)
 */
func filterPublicationQuery(params *domain.Request, binds *[]interface{}) string {
	var query string

	if params.Keyword != "" {
		*binds = append(*binds, `%`+params.Keyword+`%`)
		query = fmt.Sprintf(`%s AND pub.portal_category LIKE ?`, query)
	}

	if v, ok := params.Filters["status"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND pub.status = ?`, query)
	}

	if v, ok := params.Filters["created_by"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND pub.created_by = ?`, query)
	}

	if v, ok := params.Filters["unit_id"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND u.unit_id = ?`, query)
	}

	return query
}

/**
 * this block of code is used to generate the query for the service public
 */
func filterPublicationPortalQuery(params *domain.Request, binds *[]interface{}) string {
	var queryFilter string

	if params.Keyword != "" {
		*binds = append(*binds, `%`+params.Keyword+`%`)
		queryFilter = fmt.Sprintf(`%s AND ms.service_name LIKE ?`, queryFilter)
	}

	if v, ok := params.Filters["category"]; ok && v != "" {
		*binds = append(*binds, v)
		queryFilter = fmt.Sprintf(`%s AND mdp.portal_category = ?`, queryFilter)
	}

	if v, ok := params.Filters["status"]; ok && v != "" {
		*binds = append(*binds, v)
		queryFilter = fmt.Sprintf(`%s AND mdp.status = ?`, queryFilter)
	}

	return queryFilter
}
