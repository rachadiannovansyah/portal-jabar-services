package usecase

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
)

type visitorUsecase struct {
	externalRepository domain.ExternalVisitorRepository
	conn               utils.Conn
	contextTimeout     time.Duration
}

func NewVisitorUsecase(externalRepository domain.ExternalVisitorRepository, conn *utils.Conn, timeout time.Duration) *visitorUsecase {
	return &visitorUsecase{
		externalRepository: externalRepository,
		conn:               *conn,
		contextTimeout:     timeout,
	}
}

func (uc visitorUsecase) GetCounterVisitor(ctx context.Context, path string) (result domain.CounterVisitorResponse) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.externalRepository.GetCounterVisitor()
	if err != nil {
		redis, _ := uc.conn.Redis.Get(path).Result()
		json.Unmarshal([]byte(redis), &result)
		return
	}

	result = domain.CounterVisitorResponse{
		Data: res.Result,
	}

	uc.conn.Redis.Set(path, result, 0)

	return result
}
