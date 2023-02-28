package mysql

import (
	"fmt"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

/**
 * this block of code is used to generate the query
 */
func filterGovernmentAffairQuery(params *domain.Request, binds *[]interface{}) string {
	var queryFilter string

	if params.Keyword != "" {
		*binds = append(*binds, `%`+params.Keyword+`%`)
		queryFilter = fmt.Sprintf(`%s AND main_affair LIKE ?`, queryFilter)
	}

	return queryFilter
}
