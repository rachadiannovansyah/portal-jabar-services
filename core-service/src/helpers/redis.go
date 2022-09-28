package helpers

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
	"github.com/labstack/echo/v4"
)

// func to set connection of redis
func RedisCache() *redis.Client {
	cfg := config.NewConfig()
	cache := utils.NewDBConn(cfg).Redis

	// check connection redis
	_, err := cache.Ping().Result()
	if err != nil {
		panic(err)
	}

	return cache
}

// func to set cache redis
func SetCache(key string, value interface{}, ttl time.Duration) error {
	// set key value and ttl on redis
	err := RedisCache().Set(key, value, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}

// func to get cache redis
func GetCache(c echo.Context) (memcached string) {
	// get url path from context
	path := c.Request().URL.RequestURI()

	// Get cached redis data
	memcached, _ = RedisCache().Get(path).Result()

	if memcached != "" {
		fmt.Println("Cached!")
		return
	}

	return
}
