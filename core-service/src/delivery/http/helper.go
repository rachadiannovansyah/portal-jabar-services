package http

import (
	"fmt"
	"math"
	"strconv"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/labstack/echo/v4"
)

// GetCurrentURI ...
func GetCurrentURI(c echo.Context) string {
	req := c.Request()

	proto := "http://"
	if req.TLS != nil {
		proto = "https://"
	}

	return proto + req.Host + req.URL.Path
}

// GetRequestParams ...
func GetRequestParams(c echo.Context) domain.Request {
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	perPage, _ := strconv.ParseInt(c.QueryParam("per_page"), 10, 64)

	if page == 0 {
		page = 1
	}

	if perPage == 0 {
		perPage = 10
	}

	offset := (page - 1) * perPage

	params := domain.Request{
		Keyword: c.QueryParam("q"),
		Page:    page,
		PerPage: perPage,
		Offset:  offset,
		OrderBy: c.QueryParam("order_by"),
		SortBy:  c.QueryParam("sort_by"),
	}

	return params
}

// Paginate ...
func Paginate(c echo.Context, data interface{}, total int64, params domain.Request) *domain.ResultsData {
	totalPage := math.Ceil(float64(total) / float64(params.PerPage))
	requestURI := GetCurrentURI(c) + "?page=%v&per_page=" + strconv.FormatInt(params.PerPage, 10)

	var firstPage, prevPage, nextPage, lastPage string
	if total > 0 {
		firstPage = fmt.Sprintf(requestURI, 1)
		if params.Page > 1 {
			prevPage = fmt.Sprintf(requestURI, params.Page-1)
		}
		if float64(params.Page) < totalPage {
			nextPage = fmt.Sprintf(requestURI, params.Page+1)
		}
		lastPage = fmt.Sprintf(requestURI, int(totalPage))
	}

	if params.Keyword != "" {
		requestURI += "&q=" + c.QueryParam("q")
	}

	return &domain.ResultsData{
		Data: data,
		Meta: &domain.MetaData{
			TotalCount:  total,
			TotalPage:   totalPage,
			CurrentPage: params.Page,
			PerPage:     params.PerPage,
		},
		Links: &domain.LinksData{
			First: helpers.SetPointerString(firstPage),
			Last:  helpers.SetPointerString(lastPage),
			Next:  helpers.SetPointerString(nextPage),
			Prev:  helpers.SetPointerString(prevPage),
		},
	}
}
