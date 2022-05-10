package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/google/uuid"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	_newsRepo "github.com/jabardigitalservice/portal-jabar-services/core-service/src/modules/news/repository/mysql"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	var err error

	cfg := config.NewConfig()

	if len(os.Args) <= 2 {
		logrus.Error("Usage:", os.Args[1], "command", "argument")
		return errors.New("invalid command")
	}

	switch os.Args[1] {
	case "migrate":
		dsn := cfg.DB.DSN + "&multiStatements=true"
		err = Migrate(dsn, os.Args[2])
	case "seed":
		err = errors.New("to be develop")
	case "es:mapping":
		err = DoMapping(cfg, os.Args[2])
	case "es:sync":
		err = DoSyncElastic(cfg, os.Args[2])
	case "es:truncate":
		err = DoTruncate(cfg, os.Args[2])
	default:
		err = errors.New("must specify a command")
	}

	if err != nil {
		return err
	}

	return nil
}

// Migrate to run database migration up or down
func Migrate(dsn string, command string) error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logrus.Error(err)
		return err
	}
	path, err := os.Getwd()
	if err != nil {
		logrus.Error(err)
		return err
	}

	migrationPath := fmt.Sprintf("file://%s/migrations", path)
	logrus.Infof("migrationPath : %s", migrationPath)

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		logrus.Error(err)
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		migrationPath,
		"mysql",
		driver,
	)
	if err != nil {
		logrus.Error(err)
		return err
	}
	if command == "up" {
		logrus.Info("Migrate up")
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			logrus.Errorf("An error occurred while syncing the database.. %v", err)
			return err
		}
	}

	if command == "down" {
		logrus.Info("Migrate down")
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			logrus.Errorf("An error occurred while syncing the database.. %v", err)
			return err
		}
	}

	if err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Migrate complete")
	return nil
}

func DoTruncate(cfg *config.Config, command string) error {
	es, err := elasticsearch.NewClient(*cfg.ELastic.ElasticConfig)
	if err != nil {
		logrus.Error(err)
		return err
	}

	mappings := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}

	mappingsBytes, _ := json.Marshal(&mappings)

	req := esapi.DeleteByQueryRequest{
		Index: []string{cfg.ELastic.IndexContent},
		Body:  bytes.NewReader(mappingsBytes),
	}
	resMap, errMap := req.Do(context.Background(), es)
	if errMap != nil {
		log.Fatalf("Delete Error getting response: %s", errMap)
	}
	// Check response status
	if resMap.IsError() {
		log.Fatalf("Error: %s", resMap.String())
	}

	log.Println("truncate result: ", resMap.String())

	return nil
}

func DoSyncElastic(cfg *config.Config, command string) error {
	es, err := elasticsearch.NewClient(*cfg.ELastic.ElasticConfig)
	if err != nil {
		logrus.Error(err)
		return err
	}
	log.SetFlags(0)

	var (
		wg sync.WaitGroup
	)

	// 1. Get cluster info
	res, err := es.Info()
	if err != nil {
		logrus.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	logrus.Println(res)

	dbConn := utils.NewDBConn(cfg)

	defer func() {
		err := dbConn.Mysql.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	newsRepo := _newsRepo.NewMysqlNewsRepository(dbConn.Mysql)
	news, _, err := newsRepo.Fetch(context.TODO(), &domain.Request{PerPage: 100})

	if err != nil {
		logrus.Error(err)
	}

	// 3. Index documents concurrently
	for i, data := range news {
		wg.Add(1)

		go func(i int, data domain.News) {
			defer wg.Done()

			// Build the request body.
			var b strings.Builder
			re := regexp.MustCompile(`\r?\n`)

			excerpt := ""
			if err != nil {
				panic(err)
			}

			content := "" //html2text.HTML2Text(data.Content)
			if err != nil {
				panic(err)
			}
			excerpt = strings.ReplaceAll(re.ReplaceAllString(excerpt, " "), `"`, "")
			content = strings.ReplaceAll(re.ReplaceAllString(content, " "), `"`, "")

			//format time
			layout := "2006-01-02 15:04:05"

			b.WriteString(fmt.Sprintf(`{"id" : %v,`, data.ID))
			b.WriteString(fmt.Sprintf(`"domain" : "%v",`, "news"))
			b.WriteString(fmt.Sprintf(`"title" : "%v",`, data.Title))
			b.WriteString(fmt.Sprintf(`"excerpt" : "%v",`, excerpt))
			b.WriteString(fmt.Sprintf(`"content" : "%v",`, content))
			b.WriteString(fmt.Sprintf(`"slug" : "%v",`, data.Slug))
			b.WriteString(fmt.Sprintf(`"category" : "%v",`, data.Category))
			b.WriteString(fmt.Sprintf(`"views" : "%v",`, data.Views))
			b.WriteString(fmt.Sprintf(`"shared" : "%v",`, data.Shared))
			b.WriteString(fmt.Sprintf(`"created_at" : "%s",`, data.CreatedAt.Format(layout)))
			b.WriteString(fmt.Sprintf(`"updated_at" : "%s",`, data.UpdatedAt.Format(layout)))
			b.WriteString(fmt.Sprintf(`"thumbnail" : "%v",`, data.Image))
			b.WriteString(fmt.Sprintf(`"is_active" : %v}`, data.IsLive == 1))
			fmt.Println(b.String())

			// Set up the request object.
			req := esapi.IndexRequest{
				Index:      cfg.ELastic.IndexContent,
				DocumentID: uuid.New().String(),
				Body:       strings.NewReader(b.String()),
				Refresh:    "true",
			}

			// Perform the request with the client.
			res, err := req.Do(context.Background(), es)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				log.Println(res)
				log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
			} else {
				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
				}
			}
		}(i, data)
	}
	wg.Wait()

	return nil
}

func DoMapping(cfg *config.Config, command string) error {
	es, _ := elasticsearch.NewClient(*cfg.ELastic.ElasticConfig)

	mappings := `
	{
    "mappings" : {
      "properties" : {
        "category" : {
          "type" : "keyword"
        },
        "content" : {
          "type" : "text"
        },
        "created_at" : {
          "type" : "date",
          "format" : "yyyy-MM-dd HH:mm:ss"
        },
        "domain" : {
          "type" : "keyword"
        },
        "excerpt" : {
          "type" : "text"
        },
        "id" : {
          "type" : "long"
        },
        "shared" : {
          "type" : "long"
        },
        "slug" : {
          "type" : "keyword"
        },
        "thumbnail" : {
          "type" : "keyword"
        },
        "title" : {
          "type" : "text"
        },
        "updated_at" : {
          "type" : "date",
          "format" : "yyyy-MM-dd HH:mm:ss"
        },
        "views" : {
          "type" : "long"
        },
				"is_active" : {
					"type" : "boolean"
				}
      }
    }
  }
}`

	// indices.create

	req := esapi.IndicesCreateRequest{
		Index: cfg.ELastic.IndexContent,
		Body:  strings.NewReader(mappings),
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Println(res)
		log.Printf("[%s] Error creating index", res.Status())
		return err
	} else {
		log.Printf("[%s] Index created", res.Status())
	}

	return nil

}
