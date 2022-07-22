package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	middl "github.com/jabardigitalservice/portal-jabar-services/core-service/src/middleware"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
)

// PublicNewsHandler ...
type PublicNewsHandler struct {
	CUsecase domain.NewsUsecase
}

// NewPublicNewsHandler will initialize the /public/news handler
func NewPublicNewsHandler(p *echo.Group, us domain.NewsUsecase) {
	handler := &PublicNewsHandler{CUsecase: us}
	p.GET("/news", handler.FetchNews)
	p.GET("/news/slug/:slug", handler.GetBySlug, middl.CheckCache())
	p.GET("/news/banner", handler.FetchNewsBanner)
	p.GET("/news/headline", handler.FetchNewsHeadline)
	p.PATCH("/news/:id/share", handler.AddShare)
}

// FetchNews will fetch the content based on given params
func (h *PublicNewsHandler) FetchNews(c echo.Context) error {
	ctx := c.Request().Context()

	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"category":   c.QueryParam("cat"),
		"categories": c.Request().URL.Query()["cat[]"],
		"highlight":  c.QueryParam("highlight"),
		"type":       c.QueryParam("type"),
		"tag":        c.QueryParam("tag"),
		"status":     c.QueryParam("status"),
		"exclude":    c.QueryParam("exclude"),
		"is_aptika":  c.QueryParam("is_aptika"),
	}

	listNews, total, err := h.CUsecase.FetchPublished(ctx, &params)

	if err != nil {
		return err
	}

	// Set news response for API Aptika
	isAptika, _ := strconv.ParseBool(params.Filters["is_aptika"].(string))
	if isAptika {
		listAptikaNewsRes := []domain.NewsAptikaResponse{}
		copier.Copy(&listAptikaNewsRes, &listNews)

		res := helpers.Paginate(c, listAptikaNewsRes, total, params)

		return c.JSON(http.StatusOK, res)
	}

	// Copy slice to slice
	listNewsRes := []domain.NewsListResponse{}
	copier.Copy(&listNewsRes, &listNews)

	res := helpers.Paginate(c, listNewsRes, total, params)

	return c.JSON(http.StatusOK, res)
}

// GetBySlug will get article by given slug
func (h *PublicNewsHandler) GetBySlug(c echo.Context) error {
	cfg := config.NewConfig()
	cache := utils.NewDBConn(cfg).Redis
	slug := c.Param("slug")
	ctx := c.Request().Context()

	news, err := h.CUsecase.GetBySlug(ctx, slug)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	// Copy slice to slice
	newsRes := domain.DetailNewsResponse{}
	copier.Copy(&newsRes, &news)

	// set ttl cache 5 minutes after store in redis
	ttl := 5 * time.Minute
	cacheErr := cache.Set(slug, &newsRes, ttl).Err()
	if cacheErr != nil {
		return cacheErr
	}

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &newsRes})
}

// FetchNews will fetch the content based on given params
func (h *PublicNewsHandler) FetchNewsBanner(c echo.Context) error {

	ctx := c.Request().Context()

	listNews, err := h.CUsecase.FetchNewsBanner(ctx)

	if err != nil {
		return err
	}

	res := map[string]interface{}{
		"data": listNews,
	}

	return c.JSON(http.StatusOK, res)
}

// FetchNewsHeadline ...
func (h *PublicNewsHandler) FetchNewsHeadline(c echo.Context) error {

	ctx := c.Request().Context()

	headlineNews, err := h.CUsecase.FetchNewsHeadline(ctx)

	if err != nil {
		return err
	}

	// Copy slice to slice
	headlineNewsRes := []domain.NewsBanner{}
	copier.Copy(&headlineNewsRes, &headlineNews)

	res := map[string]interface{}{
		"data": headlineNewsRes,
	}

	return c.JSON(http.StatusOK, res)
}

// AddShare counter to share
func (h *PublicNewsHandler) AddShare(c echo.Context) error {
	// FIXME: Check and verify the recaptcha response token.

	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	err = h.CUsecase.AddShare(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully add share count",
	})
}
