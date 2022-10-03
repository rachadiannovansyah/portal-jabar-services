package mysql

import (
	"fmt"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

/**
 * this block of code is used to generate the query for the public service
 */
func filterPublicServiceQuery(params *domain.Request) string {
	var queryFilter string

	if params.Keyword != "" {
		queryFilter += ` AND name LIKE '%` + params.Keyword + `%' `
	}

	if v, ok := params.Filters["service_type"]; ok && v != "" {
		queryFilter = fmt.Sprintf(`%s AND service_type = "%s"`, queryFilter, v)
	}

	return queryFilter
}
