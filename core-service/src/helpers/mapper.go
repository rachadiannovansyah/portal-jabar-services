package helpers

func MapValue(arrMap []interface{}, key string) (value float64) {
	for _, m := range arrMap {
		el := m.(map[string]interface{})
		if el["key"] == key {
			return el["doc_count"].(float64)
		}
	}
	return
}
