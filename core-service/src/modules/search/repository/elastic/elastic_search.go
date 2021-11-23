package mysql

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/sirupsen/logrus"
)

type elasticSearchRepository struct {
	Conn *elasticsearch.Client
}

// NewElasticSearchRepository will create an object that represent the news.Repository interface
func NewElasticSearchRepository(es *elasticsearch.Client) domain.SearchRepository {
	return &elasticSearchRepository{es}
}

func (es *elasticSearchRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.SearchListResponse, tot int, err error) {
	logrus.Print("MASHOK")
	esClient := es.Conn

	// Pass the JSON query to the Golang client's Search() method
	resp, err := esClient.Search(
		esClient.Search.WithContext(ctx),
		esClient.Search.WithIndex("ipj-content-staging"),
	)

	fmt.Println("res", resp)

	return
}
