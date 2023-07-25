package helpers

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// RangeLastWeek represent today and the day last of last week
type RangeLastWeek struct {
	DayOfLastWeek string
	Today         string
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
		StartDate: c.QueryParam("start_date"),
		EndDate:   c.QueryParam("end_date"),
	}

	return params
}

// Paginate ...
func Paginate(_ echo.Context, data interface{}, total int64, params domain.Request) *domain.ResultsData {
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
	case domain.ErrSlugAlreadyExist:
		return http.StatusBadRequest
	case domain.ErrBadRequest:
		return http.StatusBadRequest
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

// GetRangeLastWeek for get date within a week range ...
func GetRangeLastWeek() RangeLastWeek {
	format := "2006-01-02 15:04:05"
	return RangeLastWeek{
		Today:         time.Now().Format(format),
		DayOfLastWeek: time.Now().AddDate(0, 0, -7).Format(format),
	}
}

// GetObjectFromString ...
func GetObjectFromString(str string, obj interface{}) error {
	return json.Unmarshal([]byte(str), &obj)
}

// GetStringFromObject ...
func GetStringFromObject(obj interface{}) string {
	data, _ := json.Marshal(obj)
	return string(data)
}

func GetInBind(binds *[]interface{}, arr []string) string {
	bind := "("
	for i, str := range arr {
		*binds = append(*binds, str)
		bind += "?"
		if i < len(arr)-1 {
			bind += ","
		}
	}
	bind += ")"

	return bind
}

func GetMetaDataImage(link string) (meta domain.DetailMetaDataImage, err error) {
	subStringsSlice := strings.Split(link, "/")
	fileName := subStringsSlice[len(subStringsSlice)-1]

	resp, err := http.Head(link)
	if err != nil {
		logrus.Error(err)
		return domain.DetailMetaDataImage{}, err
	}

	if resp.StatusCode != http.StatusOK {
		logrus.Error(err)
		return domain.DetailMetaDataImage{}, err
	}

	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	fileSize := int64(size)

	meta.FileName = fileName
	meta.FileDownloadUri = link
	meta.Size = fileSize

	return meta, err
}
