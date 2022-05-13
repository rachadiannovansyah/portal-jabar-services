package mysql

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/google/uuid"
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
		source := doc["_source"].(map[string]interface{})

		// mapstructure
		searchData := domain.SearchListResponse{}
		mapstructure.Decode(source, &searchData)

		// parsing the date string to time.Time
		searchData.CreatedAt = helpers.ParseESDate(source["created_at"].(string))

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
		"_source": q{
			"includes": []string{"id", "domain", "title", "excerpt", "slug", "category", "thumbnail", "created_at"},
		},
		"sort": []map[string]interface{}{
			q{"created_at": q{"order": params.SortOrder}},
		},
		"query": q{
			"bool": q{
				"must": []q{
					q{
						"match": q{
							"is_active": true,
						},
					},
					q{
						"multi_match": q{
							"query":         params.Keyword,
							"fields":        []string{"title", "content"},
							"type":          "best_fields",
							"fuzziness":     "AUTO",
							"prefix_length": 2,
						},
					},
				},
			},
		},
		"aggs": q{
			"agg_domain": q{
				"terms": q{
					"field": "domain",
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

func (es *elasticSearchRepository) Fetch(ctx context.Context, indices string, params *domain.Request) (docs []domain.SearchListResponse, total int64, aggs interface{}, err error) {
	esClient := es.Conn
	query := buildQuery(params)

	// Pass the JSON query to the Golang client's Search() method
	resp, err := esClient.Search(
		esClient.Search.WithContext(ctx),
		esClient.Search.WithIndex(indices), // FIXME: this should use env
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

func (es *elasticSearchRepository) SearchSuggestion(ctx context.Context, indices string, params *domain.Request) (res []domain.SuggestResponse, err error) {
	var buf bytes.Buffer
	key := params.Filters["suggestions"]

	query := q{
		"_source": q{
			"includes": "title",
		},
		"query": q{
			"multi_match": q{
				"query":         key,
				"fields":        "title",
				"type":          "best_fields",
				"fuzziness":     "AUTO",
				"prefix_length": 2,
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
		esclient.Search.WithIndex(indices),
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
			source := doc["_source"].(map[string]interface{})

			// mapstructure using suggest response
			suggestData := domain.SuggestResponse{}
			mapstructure.Decode(source, &suggestData)

			res = append(res, suggestData)
		}
	}

	return
}

func (es *elasticSearchRepository) Store(ctx context.Context, indices string, data *domain.Search) (err error) {
	esclient := es.Conn

	req := esapi.IndexRequest{
		Index:      indices,
		DocumentID: uuid.New().String(),
		Body:       strings.NewReader(data.Content),
		Refresh:    "true",
	}

	res, err := req.Do(ctx, esclient)
	if err != nil {
		return
	}

	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("[%s] %s", res.Status(), res.String())
	} else {
		logrus.Printf("[%s] Index created", res.Status())
	}

	return
}

func (es *elasticSearchRepository) Update(ctx context.Context, indices string, id int, data *domain.Search) (err error) {
	esclient := es.Conn

	doc := q{
		"query": q{
			"bool": q{
				"must": []q{
					q{
						"match": q{
							"id": id,
						},
					},
					q{
						"match": q{
							"domain": data.Domain,
						},
					},
				},
			},
		},
		"script": q{
			"source": "ctx._source.content = params.content",
			"params": q{
				"content": data.Content,
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		return err
	}

	esRes, esErr := esclient.UpdateByQuery(
		[]string{indices},
		esclient.UpdateByQuery.WithBody(&buf),
		esclient.UpdateByQuery.WithContext(context.Background()),
	)

	if esErr != nil {
		return esErr
	}

	defer esRes.Body.Close()
	fmt.Println(esRes.String())

	return
}

func (es *elasticSearchRepository) Delete(ctx context.Context, indices string, id int, domain string) (err error) {
	esclient := es.Conn

	doc := q{
		"query": q{
			"bool": q{
				"must": []q{
					q{
						"match": q{
							"id": id,
						},
					},
					q{
						"match": q{
							"domain": domain,
						},
					},
				},
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		return err
	}

	esRes, esErr := esclient.DeleteByQuery([]string{indices}, &buf)

	if esErr != nil {
		return esErr
	}

	defer esRes.Body.Close()
	fmt.Println(esRes.String())

	return
}
