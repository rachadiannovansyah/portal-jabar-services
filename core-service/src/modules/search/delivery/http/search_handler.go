package http

import (
	"fmt"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
)

// SearchHandler ...
type SearchHandler struct {
	SUsecase domain.SearchUsecase
}

// NewSearchHandler will initialize the search/ resources endpoint
func NewSearchHandler(e *echo.Group, r *echo.Group, us domain.SearchUsecase) {
	handler := &SearchHandler{
		SUsecase: us,
	}
	e.GET("/search", handler.FetchSearch)
}

// FetchNews will fetch the content based on given params
func (h *SearchHandler) FetchSearch(c echo.Context) error {
	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)
	listSearch, tot, aggs, err := h.SUsecase.Fetch(ctx, &params)
	if err != nil {
		return err
	}
	res := helpers.Paginate(c, listSearch, tot, params)
	meta := res.Meta.(*domain.MetaData)

	// FIXME: meta aggregation structure in the next PR
	aggDomain := aggs.(map[string]interface{})["agg_domain"].(map[string]interface{})["buckets"]
	fmt.Println("aggDomain", aggDomain)
	meta.Aggregations = &domain.MetaAggregations{
		Domain: domain.AggDomain{
			News: 1, // FIXME: remove hardcoded aggregate
		},
	}
	return c.JSON(http.StatusOK, res)
}
