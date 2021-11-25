package mysql

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
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

func (es *elasticSearchRepository) Fetch(ctx context.Context, params *domain.Request) (res []domain.SearchListResponse, tot int64, err error) {

	var buf bytes.Buffer
	query := map[string]interface{}{}
	esclient := es.Conn

	// Pass the JSON query to the Golang client's Search() method
	resp, err := esclient.Search(
		esclient.Search.WithContext(ctx),
		esclient.Search.WithIndex("ipj-content-staging"),
		esclient.Search.WithBody(&buf),
		esclient.Search.WithFrom(int(params.Offset)),
		esclient.Search.WithSize(int(params.PerPage)),
	)

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		failOnError(err, "Error encoding query")
	}

	// Decode the JSON response and using a pointer
	if err := json.NewDecoder(resp.Body).Decode(&mapResp); err != nil {
		failOnError(err, "Error parsing the response body")

		// If no error, then convert response to a map[string]interface
	} else {
		tot = int64(mapResp["hits"].(map[string]interface{})["total"].(interface{}).(map[string]interface{})["value"].(float64))

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
	}

	return
}

func (es *elasticSearchRepository) SearchSuggestion(ctx context.Context, key string) (res []domain.SearchSuggestionResponse, err error) {

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  key,
				"fields": []string{"title", "content"},
			},
		},
	}

	esclient := es.Conn

	// Pass the JSON query to the Golang client's Search() method
	resp, err := esclient.Search(
		esclient.Search.WithContext(ctx),
		esclient.Search.WithIndex("ipj-content-staging"),
		esclient.Search.WithBody(&buf),
		esclient.Search.WithSize(5),
	)

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		failOnError(err, "Error encoding query")
	}

	// Decode the JSON response and using a pointer
	if err := json.NewDecoder(resp.Body).Decode(&mapResp); err != nil {

		failOnError(err, "Error parsing the response body")

	} else {
		// Iterate the document "hits" returned by API call
		for _, hit := range mapResp["hits"].(map[string]interface{})["hits"].([]interface{}) {

			// Parse the attributes/fields of the document
			doc := hit.(map[string]interface{})

			// The "_source" data is another map interface nested inside of doc
			source := doc["_source"]

			// mapstructure
			searchSuggestData := domain.SearchSuggestionResponse{}
			mapstructure.Decode(source, &searchSuggestData)

			res = append(res, searchSuggestData)
		}
	}

	return
}
