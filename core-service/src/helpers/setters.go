package helpers

import (
	"encoding/json"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
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
	if time.Now().Unix() > endDate.Unix() && (endDate != time.Time{}) { // its mean end date is not empty
		newsRequest.StartDate = ConvertTimeToString(currentTime)
		newsRequest.EndDate = ConvertTimeToString(currentTime.AddDate(0, 0, int(newsRequest.Duration)))
		setLiveAndPublish(newsRequest, isLive, currentTime)
	}

	// if startDate is less than equal to today, set live to true
	if startDate.Unix() <= time.Now().Unix() {
		setLiveAndPublish(newsRequest, isLive, currentTime)
	}
}

func Cache(key string, data interface{}, meta interface{}) (err error) {
	cacheData := domain.Cache{
		Data: data,
		Meta: meta,
	}

	// set cache from dependency injection redis
	ttl := time.Duration(config.NewConfig().Redis.TTL) * time.Second
	value, _ := json.Marshal(cacheData)
	cacheErr := SetCache(key, value, ttl)
	if cacheErr != nil {
		return cacheErr
	}

	return
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
