package helpers

func GetESTotalCount(mapResp map[string]interface{}) float64 {
	return mapResp["hits"].(map[string]interface{})["total"].(interface{}).(map[string]interface{})["value"].(float64)
}
