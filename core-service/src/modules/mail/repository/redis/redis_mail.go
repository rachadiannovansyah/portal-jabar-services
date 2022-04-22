package redis

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type redisMailRepository struct {
	redisClient *redis.Client
}

func NewRedisMailRepository(redisClient *redis.Client) domain.MailRepository {
	return &redisMailRepository{
		redisClient: redisClient,
	}
}

func (r *redisMailRepository) Enqueue(ctx context.Context, mail domain.Mail) error {
	serializedValue, _ := json.Marshal(mail)
	return r.redisClient.RPush(ctx, "email-queue", serializedValue).Err()
}
