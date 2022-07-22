package helpers

import (
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

// SetPropLiveNews ...
func SetPropLiveNews(news *domain.StoreNewsRequest) {

	currentTime := time.Now()
	endDate := ConvertStringToTime(news.EndDate)
	news.IsLive = 1
	news.PublishedAt = &currentTime
	// if endDate is exceed current time, set it to published at + duration
	if time.Now().Unix() > endDate.Unix() {
		news.StartDate = ConvertTimeToString(currentTime)
		news.EndDate = ConvertTimeToString(currentTime.AddDate(0, 0, int(news.Duration)))
	}
}
