package mysql

import (
	"bytes"
	"context"
	"encoding/json"

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

func buildQuery(params *domain.Request) (buf bytes.Buffer) {
	domain := []string{"news", "information", "public_service", "announcement", "about"}
	if domainFilter := params.Filters["domain"].([]string); len(domainFilter) > 0 {
		domain = domainFilter
	}

	query := q{
		"query": q{
			"multi_match": q{
				"query":  params.Keyword,
				"fields": []string{"title", "content"},
				"type":   "best_fields",
			},
		},
		"aggs": q{
			"agg_domain": q{
				"terms": q{
					"field": "domain.keyword",
				},
			},
		},
		"post_filter": q{
			"terms": q{
				"domain": domain,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		failOnError(err, "Error encoding query")
	}

	return
}

func (es *elasticSearchRepository) Fetch(ctx context.Context, params *domain.Request) (docs []domain.SearchListResponse, total int64, aggs interface{}, err error) {
	esClient := es.Conn
	query := buildQuery(params)

	// Pass the JSON query to the Golang client's Search() method
	resp, err := esClient.Search(
		esClient.Search.WithContext(ctx),
		esClient.Search.WithIndex("ipj-content-staging"), // FIXME: this should use env
		esClient.Search.WithBody(&query),
		esClient.Search.WithFrom(int(params.Offset)),
		esClient.Search.WithSize(int(params.PerPage)),
	)

	// Decode the JSON response and using a pointer
	if err := json.NewDecoder(resp.Body).Decode(&mapResp); err != nil {
		failOnError(err, "Error parsing the response body")

		// If no error, then convert response to a map[string]interface
	} else {
		docs = mapElasticDocs(mapResp)
		total = int64(helpers.GetESTotalCount(mapResp))
		aggs = mapResp["aggregations"].(map[string]interface{})
	}

	return
}

func (es *elasticSearchRepository) SearchSuggestion(ctx context.Context, params *domain.Request) (res []domain.SuggestResponse, err error) {
	var buf bytes.Buffer
	key := params.Filters["suggestions"]

	query := q{
		"_source": q{
			"includes": "title",
		},
		"query": q{
			"bool": q{
				"must": q{
					"term": q{"title": key},
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
		esclient.Search.WithIndex("ipj-content-staging"),
		esclient.Search.WithBody(&buf),
		esclient.Search.WithSize(5),
	)

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

			// mapstructure using suggest response
			suggestData := domain.SuggestResponse{}
			mapstructure.Decode(source, &suggestData)

			res = append(res, suggestData)
		}
	}

	return
}
