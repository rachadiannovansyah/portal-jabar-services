package helpers

import (
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

// SetPropLiveNews ...
func SetPropLiveNews(newsRequest *domain.StoreNewsRequest) {

	currentTime := time.Now()
	endDate := ConvertStringToTime(newsRequest.EndDate)
	startDate := ConvertStringToTime(newsRequest.StartDate)
	newsRequest.IsLive = 0

	// if endDate is exceed current time, set it to published at + duration
	if time.Now().Unix() > endDate.Unix() {
		newsRequest.StartDate = ConvertTimeToString(currentTime)
		newsRequest.EndDate = ConvertTimeToString(currentTime.AddDate(0, 0, int(newsRequest.Duration)))
		newsRequest.IsLive = 1
		newsRequest.PublishedAt = &currentTime
	}

	// if startDate is less than equal to today, set live to true
	if startDate.Unix() <= time.Now().Unix() {
		newsRequest.IsLive = 1
		newsRequest.PublishedAt = &currentTime
	}
}
