package http

import (
	"fmt"
	"net/http"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/labstack/echo/v4"
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
	e.GET("/search/suggestion/:suggestion", handler.SearchSuggestion)
}

// FetchSearch will fetch the content based on given params
func (h *SearchHandler) FetchSearch(c echo.Context) error {
	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)
	listSearch, tot, err := h.SUsecase.Fetch(ctx, &params)
	if err != nil {
		return err
	}
	res := helpers.Paginate(c, listSearch, tot, params)
	return c.JSON(http.StatusOK, res)
}

// SearchSuggestion ...
func (h *SearchHandler) SearchSuggestion(c echo.Context) error {
	ctx := c.Request().Context()
	suggestion := c.Param("suggestion")
	fmt.Println(suggestion)
	if suggestion == "" {
		return c.JSON(http.StatusOK, "")
	}

	listSuggest, err := h.SUsecase.SearchSuggestion(ctx, suggestion)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, listSuggest)
}
