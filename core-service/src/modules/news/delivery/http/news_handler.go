package http

import (
	"net/http"
	"strconv"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/go-playground/validator.v9"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

// NewsHandler ...
type NewsHandler struct {
	CUsecase domain.NewsUsecase
}

func isRequestValid(n *domain.StoreNewsRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(n)
	if err != nil {
		return false, err
	}
	return true, nil
}

// NewNewsHandler will initialize the contents/ resources endpoint
func NewNewsHandler(e *echo.Group, r *echo.Group, us domain.NewsUsecase) {
	handler := &NewsHandler{
		CUsecase: us,
	}
	e.GET("/news", handler.FetchNews)
	r.POST("/news", handler.Store)
	e.GET("/news/:id", handler.GetByID)
	r.PUT("/news/:id", handler.Update)
	r.PATCH("/news/:id/status", handler.UpdateStatus)
	e.GET("/news/slug/:slug", handler.GetBySlug)
	e.GET("/news/banner", handler.FetchNewsBanner)
	e.GET("/news/headline", handler.FetchNewsHeadline)
	e.PATCH("/news/:id/share", handler.AddShare)
	e.GET("/news/tabs", handler.TabStatus)
}

// FetchNews will fetch the content based on given params
func (h *NewsHandler) FetchNews(c echo.Context) error {

	ctx := c.Request().Context()

	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"highlight": c.QueryParam("highlight"),
		"category":  c.QueryParam("cat"),
		"type":      c.QueryParam("type"),
		"tags":      c.QueryParam("tags"),
		"status":    c.QueryParam("status"),
	}

	if c.Get("auth:user") == nil {
		params.Filters["is_live"] = "1"
	}

	listNews, total, err := h.CUsecase.Fetch(ctx, &params)

	if err != nil {
		return err
	}

	// Copy slice to slice
	listNewsRes := []domain.NewsListResponse{}
	copier.Copy(&listNewsRes, &listNews)

	res := helpers.Paginate(c, listNewsRes, total, params)

	return c.JSON(http.StatusOK, res)
}

// FetchNews will fetch the content based on given params
func (h *NewsHandler) FetchNewsBanner(c echo.Context) error {

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
func (h *NewsHandler) FetchNewsHeadline(c echo.Context) error {

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

// GetByID will get article by given id
func (h *NewsHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	news, err := h.CUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	// Copy slice to slice
	newsRes := domain.DetailNewsResponse{}
	copier.Copy(&newsRes, &news)

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &newsRes})
}

func (h *NewsHandler) TabStatus(c echo.Context) (err error) {
	ctx := c.Request().Context()

	tabs, err := h.CUsecase.TabStatus(ctx)

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &tabs})
}

// GetBySlug will get article by given slug
func (h *NewsHandler) GetBySlug(c echo.Context) error {
	slug := c.Param("slug")
	ctx := c.Request().Context()

	news, err := h.CUsecase.GetBySlug(ctx, slug)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	// Copy slice to slice
	newsRes := domain.DetailNewsResponse{}
	copier.Copy(&newsRes, &news)

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &newsRes})
}

// AddShare counter to share
func (h *NewsHandler) AddShare(c echo.Context) error {
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

// Store will store the news by given request body
func (h *NewsHandler) Store(c echo.Context) (err error) {
	n := new(domain.StoreNewsRequest)
	if err = c.Bind(n); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(n); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// FIXME: authenticated variables must be global variables to be accessible everywhere
	auth := domain.JwtCustomClaims{}
	mapstructure.Decode(c.Get("auth:user"), &auth)

	n.Author.ID = auth.ID
	n.Author.Name = auth.Name

	ctx := c.Request().Context()
	err = h.CUsecase.Store(ctx, n)
	if err != nil {
		return err
	}

	// Copy slice to slice
	res := []domain.DetailNewsResponse{}
	copier.Copy(&res, &n)

	return c.JSON(http.StatusCreated, res)
}

// Update will update the news by given request body
func (h *NewsHandler) Update(c echo.Context) (err error) {
	n := new(domain.StoreNewsRequest)
	if err = c.Bind(n); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(n); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	// FIXME: authenticated variables must be global variables to be accessible everywhere
	auth := domain.JwtCustomClaims{}
	mapstructure.Decode(c.Get("auth:user"), &auth)

	n.Author.ID = auth.ID
	n.Author.Name = auth.Name

	ctx := c.Request().Context()
	err = h.CUsecase.Update(ctx, int64(reqID), n)
	if err != nil {
		return err
	}

	// Copy slice to slice
	res := []domain.DetailNewsResponse{}
	copier.Copy(&res, &n)

	return c.JSON(http.StatusOK, res)
}

// UpdateStatus will update the news status by given request body
func (h *NewsHandler) UpdateStatus(c echo.Context) (err error) {
	n := new(domain.UpdateNewsStatusRequest)
	if err = c.Bind(n); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if err = validator.New().Struct(n); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	reqID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	ctx := c.Request().Context()
	err = h.CUsecase.UpdateStatus(ctx, int64(reqID), n.Status)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully update status",
	})
}
