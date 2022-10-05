package helpers

import (
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

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
	sortOrder := c.QueryParam("sort_order")
	page, _ := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	perPage, _ := strconv.ParseInt(c.QueryParam("per_page"), 10, 64)

	if page == 0 {
		page = 1
	}

	if perPage == 0 {
		perPage = 10
	}

	offset := (page - 1) * perPage

	sortOrderValues := []string{"ASC", "DESC"}
	sortOrderExist, sortOrderIndex := InArray(sortOrder, sortOrderValues)
	sortOrder = "DESC"
	if sortOrderExist {
		sortOrder = sortOrderValues[sortOrderIndex]
	}

	params := domain.Request{
		Keyword:   RegexReplaceString(c, c.QueryParam("q"), ""),
		Page:      page,
		PerPage:   perPage,
		Offset:    offset,
		SortBy:    RegexReplaceString(c, c.QueryParam("sort_by"), ""),
		SortOrder: sortOrder,
		StartDate: RegexReplaceString(c, c.QueryParam("start_date"), ""),
		EndDate:   RegexReplaceString(c, c.QueryParam("end_date"), ""),
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

// ESAggregate ...
func ESAggregate(aggs interface{}) *domain.MetaAggregations {
	aggDomain := aggs.(map[string]interface{})["agg_domain"].(map[string]interface{})["buckets"].([]interface{})

	return &domain.MetaAggregations{
		Domain: domain.AggDomain{
			News:          int64(MapValue(aggDomain, "news")),
			Information:   int64(MapValue(aggDomain, "information")),
			PublicService: int64(MapValue(aggDomain, "public_service")),
			Announcement:  int64(MapValue(aggDomain, "announcement")),
			About:         int64(MapValue(aggDomain, "about")),
		},
	}
}

// GetStatusCode ...
func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	case domain.ErrInvalidCredentials:
		return http.StatusUnauthorized
	case domain.ErrUserIsNotActive:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

// GetAuthenticatedUser ...
func GetAuthenticatedUser(c echo.Context) *domain.JwtCustomClaims {
	auth := domain.JwtCustomClaims{}
	mapstructure.Decode(c.Get("auth:user"), &auth)
	return &auth
}

// GetUnitInfo ...
func GetUnitInfo(u domain.Unit) domain.UnitInfo {
	return domain.UnitInfo{
		ID:   u.ID,
		Name: u.Name,
	}
}

// GetRoleInfo ...
func GetRoleInfo(r domain.Role) domain.RoleInfo {
	return domain.RoleInfo{
		ID:   r.ID,
		Name: r.Name,
	}
}

type RangeLastWeek struct {
	DayOfLastWeek string
	Today         string
}

func GetRangeLastWeek() RangeLastWeek {
	format := "2006-01-02 15:04:05"
	return RangeLastWeek{
		Today:         time.Now().Format(format),
		DayOfLastWeek: time.Now().AddDate(0, 0, -7).Format(format),
	}
}
