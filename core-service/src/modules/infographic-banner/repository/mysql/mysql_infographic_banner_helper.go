package mysql

import (
	"fmt"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

func filterInfographicBannerQuery(params domain.Request, binds *[]interface{}) string {
	var queryFilter string

	if params.Keyword != "" {
		*binds = append(*binds, `%`+params.Keyword+`%`)
		queryFilter = fmt.Sprintf(`%s AND title LIKE ?`, queryFilter)
	}

	if v, ok := params.Filters["is_active"]; ok && v != "" {
		*binds = append(*binds, v)
		queryFilter = fmt.Sprintf(`%s AND is_active = ?`, queryFilter)
	}

	return queryFilter
}
