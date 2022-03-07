package helpers

import (
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jinzhu/copier"
)

func MapValue(arrMap []interface{}, key string) (value float64) {
	for _, m := range arrMap {
		el := m.(map[string]interface{})
		if el["key"] == key {
			return el["doc_count"].(float64)
		}
	}
	return
}

// MapUserInfo ...
func MapUserInfo(u domain.User) *domain.UserInfo {
	userinfo := domain.UserInfo{}
	copier.Copy(&userinfo, &u)
	return &userinfo
}
