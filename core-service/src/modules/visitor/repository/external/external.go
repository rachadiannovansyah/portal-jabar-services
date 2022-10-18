package external

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type externalRepository struct {
	cfg config.Config
}

func NewExternalVisitorRepository(cfg *config.Config) *externalRepository {
	return &externalRepository{
		cfg: *cfg,
	}
}

func (r *externalRepository) GetCounterVisitor() (result domain.ExternalCounterVisitor, err error) {
	res, err := http.Get(r.cfg.External.CoreDataUrl + "/portal-jabar/data")

	if err != nil {
		return
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	json.Unmarshal(responseData, &result)

	return
}
