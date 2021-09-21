package http

import (
	"math"
	"strconv"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
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
		Keyword:   c.QueryParam("q"),
		Page:      page,
		PerPage:   perPage,
		Offset:    offset,
		SortBy:    c.QueryParam("sort_by"),
		SortOrder: c.QueryParam("sort_order"),
		StartDate: c.QueryParam("start_date"),
		EndDate:   c.QueryParam("end_date"),
	}

	return params
}

// Paginate ...
func Paginate(c echo.Context, data interface{}, total int64, params domain.Request) *domain.ResultsData {
	return &domain.ResultsData{
		Data: data,
		Meta: &domain.MetaData{
			TotalCount:  total,
			TotalPage:   math.Ceil(float64(total) / float64(params.PerPage)),
			CurrentPage: params.Page,
			PerPage:     params.PerPage,
		},
	}
}
