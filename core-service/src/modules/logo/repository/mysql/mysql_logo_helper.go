package mysql

import (
	"fmt"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

func filterLogoQuery(params domain.Request, binds *[]interface{}) string {
	var queryFilter string

	if params.Keyword != "" {
		*binds = append(*binds, `%`+params.Keyword+`%`)
		queryFilter = fmt.Sprintf(`%s AND title LIKE ?`, queryFilter)
	}

	return queryFilter
}
