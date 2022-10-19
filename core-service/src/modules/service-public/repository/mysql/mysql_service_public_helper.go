package mysql

import (
	"fmt"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

/**
 * this block of code is used to generate the query for the service public
 */
func filterServicePublicQuery(params *domain.Request, binds *[]interface{}) string {
	var queryFilter string

	if params.Keyword != "" {
		*binds = append(*binds, `%`+params.Keyword+`%`)
		queryFilter = fmt.Sprintf(`%s AND g.name LIKE ?`, queryFilter)
	}

	if v, ok := params.Filters["category"]; ok && v != "" {
		*binds = append(*binds, v)
		queryFilter = fmt.Sprintf(`%s AND g.category = ?`, queryFilter)
	}

	if v, ok := params.Filters["type"]; ok && v != "" {
		*binds = append(*binds, v)
		queryFilter = fmt.Sprintf(`%s AND g.type = ?`, queryFilter)
	}

	return queryFilter
}
