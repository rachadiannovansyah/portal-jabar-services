package mysql

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
)

type elasticSearchRepository struct {
	Conn *elasticsearch.Client
}

// Instantiate a mapping interface for API response
var mapResp map[string]interface{}

func failOnError(err error, msg string) {
	if err != nil {
		logrus.Fatalf("%s: %s", msg, err)
	}
}

// NewElasticSearchRepository will create an object that represent the news.Repository interface
func NewElasticSearchRepository(es *elasticsearch.Client) domain.SearchRepository {
	return &elasticSearchRepository{es}
}

func mapElasticDocs(mapResp map[string]interface{}) (res []domain.SearchListResponse) {
	// Iterate the document "hits" returned by API call
	for _, hit := range mapResp["hits"].(map[string]interface{})["hits"].([]interface{}) {

		// Parse the attributes/fields of the document
		doc := hit.(map[string]interface{})

		// The "_source" data is another map interface nested inside of doc
		source := doc["_source"]

		// mapstructure
		searchData := domain.SearchListResponse{}
		mapstructure.Decode(source, &searchData)
		res = append(res, searchData)
	}

	return
}

// type alias for map query
type q map[string]interface{}

func (es *elasticSearchRepository) Fetch(ctx context.Context, params *domain.Request) (docs []domain.SearchListResponse, total int64, aggs interface{}, err error) {

	var buf bytes.Buffer
	query := q{
		"query": q{
			"multi_match": q{
				"query":  params.Keyword,
				"fields": []string{"title", "content"},
			},
		},
		"aggs": q{
			"agg_domain": q{
				"terms": q{
					"field": "domain.keyword",
				},
			},
		},
	}

	esclient := es.Conn

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		failOnError(err, "Error encoding query")
	}

	// Pass the JSON query to the Golang client's Search() method
	resp, err := esclient.Search(
		esclient.Search.WithContext(ctx),
		esclient.Search.WithIndex("ipj-content-staging"), // FIXME: this should use env
		esclient.Search.WithBody(&buf),
		esclient.Search.WithFrom(int(params.Offset)),
		esclient.Search.WithSize(int(params.PerPage)),
	)

	// Decode the JSON response and using a pointer
	if err := json.NewDecoder(resp.Body).Decode(&mapResp); err != nil {
		failOnError(err, "Error parsing the response body")

		// If no error, then convert response to a map[string]interface
	} else {
		fmt.Println("aggs", mapResp["aggregations"].(map[string]interface{})["agg_domain"].(map[string]interface{}))
		docs = mapElasticDocs(mapResp)
		total = int64(helpers.GetESTotalCount(mapResp))
		aggs = mapResp["aggregations"].(map[string]interface{})
	}

	return
}
