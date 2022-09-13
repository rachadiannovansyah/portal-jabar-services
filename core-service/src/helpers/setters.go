package helpers

import (
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

// SetPropLiveNews ...
func SetPropLiveNews(newsRequest *domain.StoreNewsRequest) {

	currentTime := time.Now()
	isLive := int8(1)
	endDate := ConvertStringToTime(newsRequest.EndDate)
	startDate := ConvertStringToTime(newsRequest.StartDate)
	newsRequest.IsLive = 0

	// if endDate is exceed current time, set it to published at + duration
	if time.Now().Unix() > endDate.Unix() {
		newsRequest.StartDate = ConvertTimeToString(currentTime)
		newsRequest.EndDate = ConvertTimeToString(currentTime.AddDate(0, 0, int(newsRequest.Duration)))
		setLiveAndPublish(newsRequest, isLive, currentTime)
	}

	// if startDate is less than equal to today, set live to true
	if startDate.Unix() <= time.Now().Unix() {
		setLiveAndPublish(newsRequest, isLive, currentTime)
	}
}

func setLiveAndPublish(newsRequest *domain.StoreNewsRequest, args ...interface{}) *domain.StoreNewsRequest {
	// type assertion params
	isLive := args[0].(int8)
	publishedAt := args[1].(time.Time)

	// set it to struct store request
	newsRequest.IsLive = isLive
	newsRequest.PublishedAt = &publishedAt

	return newsRequest
}
