package helpers

import (
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

// SetPropLiveNews ...
func SetPropLiveNews(news *domain.StoreNewsRequest) {

	startDate := ConvertStringToTime(news.StartDate)

	// if startDate is less than equal to today, set live to true
	if startDate.Unix() <= time.Now().Unix() {
		news.IsLive = 1
		news.PublishedAt.Time = time.Now()
	}

}
