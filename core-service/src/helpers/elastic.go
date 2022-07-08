package helpers

import "time"

func GetESTotalCount(mapResp map[string]interface{}) float64 {
	return mapResp["hits"].(map[string]interface{})["total"].(interface{}).(map[string]interface{})["value"].(float64)
}

func ParseESDate(strDateTime string) time.Time {
	tLayout := "2006-01-02 15:04:05" // timestamp layout
	tCreatedAt, _ := time.Parse(tLayout, strDateTime)
	return tCreatedAt
}

func ParseESPointerDate(strDateTime string) *time.Time {
	tLayout := "2006-01-02 15:04:05" // timestamp layout
	tPublishedAt, _ := time.Parse(tLayout, strDateTime)
	return &tPublishedAt
}
