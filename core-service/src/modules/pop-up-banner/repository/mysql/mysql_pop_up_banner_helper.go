package mysql

import (
	"fmt"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

/**
 * this block of code is used to generate the query for the service public
 */
func filterPopUpBannerQuery(params *domain.Request, binds *[]interface{}) string {
	var queryFilter string

	if params.Keyword != "" {
		*binds = append(*binds, `%`+params.Keyword+`%`)
		queryFilter = fmt.Sprintf(`%s AND title LIKE ?`, queryFilter)
	}

	if v, ok := params.Filters["status"]; ok && v != "" {
		*binds = append(*binds, v)
		queryFilter = fmt.Sprintf(`%s AND status = ?`, queryFilter)
	}

	return queryFilter
}
