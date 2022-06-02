package mysql

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
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
		highlight := doc["highlight"]

		// mapstructure
		searchData := domain.SearchListResponse{}
		mapstructure.Decode(source, &searchData)

		// parsing the date string to time.Time
		searchData.CreatedAt = helpers.ParseESDate(source["created_at"].(string))
		searchData.Highlight = highlight

		res = append(res, searchData)
	}

	return
}

// type alias for map query
type q map[string]interface{}

func buildQuery(params *domain.Request) (buf bytes.Buffer) {
	domain := []string{"news", "featured_program", "public_service"}
	if domainFilter := params.Filters["domain"].([]string); len(domainFilter) > 0 {
		domain = domainFilter
	}

	paramsSort := q{"created_at": q{"order": params.SortOrder}}
	if sortFilter := params.SortBy; len(sortFilter) > 0 {
		paramsSort = q{params.SortBy: q{"order": params.SortOrder}}
	}

	query := q{
		"_source": q{
			"includes": []string{"id", "domain", "title", "excerpt", "slug", "category", "thumbnail", "content", "unit", "url", "created_at"},
		},
		"sort": []map[string]interface{}{
			paramsSort,
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
							"query":  params.Keyword,
							"fields": []string{"title", "content"},
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
		"highlight": q{
			"pre_tags":  []string{"<strong>"},
			"post_tags": []string{"</strong>"},
			"fields": q{
				"content": q{"number_of_fragments": 4},
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
	fmt.Println(query.String())

	// Pass the JSON query to the Golang client's Search() method
	resp, err := esClient.Search(
		esClient.Search.WithContext(ctx),
		esClient.Search.WithIndex(indices),
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

	//format time
	formatTime := "2006-01-02 15:04:05"

	body, err := goquery.NewDocumentFromReader(strings.NewReader(data.Content))
	if err != nil {
		return
	}
	re := regexp.MustCompile(`\r?\n`)
	content := strings.ReplaceAll(re.ReplaceAllString(body.Text(), " "), `"`, "")

	// prepare the data to be indexed
	doc := q{
		"id":         data.ID,
		"domain":     data.Domain,
		"title":      data.Title,
		"excerpt":    data.Excerpt,
		"content":    content,
		"slug":       data.Slug,
		"category":   data.Category,
		"thumbnail":  data.Thumbnail,
		"created_at": data.CreatedAt.Format(formatTime),
		"updated_at": data.UpdatedAt.Format(formatTime),
		"is_active":  data.IsActive,
	}

	jsonString, err := json.Marshal(doc)

	req := esapi.IndexRequest{
		Index:      indices,
		DocumentID: uuid.New().String(),
		Body:       strings.NewReader(string(jsonString)),
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

	body, err := goquery.NewDocumentFromReader(strings.NewReader(data.Content))
	if err != nil {
		return
	}
	re := regexp.MustCompile(`\r?\n`)
	content := strings.ReplaceAll(re.ReplaceAllString(body.Text(), " "), `"`, "")

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
		// update all fields
		"script": q{
			"source": "ctx._source.title = params.title; ctx._source.excerpt = params.excerpt; ctx._source.content = params.content; ctx._source.slug = params.slug; ctx._source.category = params.category; ctx._source.thumbnail = params.thumbnail; ctx._source.updated_at = params.updated_at; ctx._source.is_active = params.is_active;",
			"lang":   "painless",
			"params": q{
				"title":      data.Title,
				"excerpt":    data.Excerpt,
				"content":    content,
				"slug":       data.Slug,
				"category":   data.Category,
				"thumbnail":  data.Thumbnail,
				"updated_at": data.UpdatedAt.Format("2006-01-02 15:04:05"),
				"is_active":  data.IsActive,
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
