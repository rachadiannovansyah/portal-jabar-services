package http

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

// TagHandler ...
type TagHandler struct {
	TagUsecase domain.TagUsecase
}

// NewTagHandler will initialize the contents/ resources endpoint
func NewTagHandler(e *echo.Group, r *echo.Group, us domain.TagUsecase) {
	handler := &TagHandler{
		TagUsecase: us,
	}
	e.GET("/tags", handler.FetchTag)
}

// FetchTag will fetch the content based on given params
func (h *TagHandler) FetchTag(c echo.Context) error {

	ctx := c.Request().Context()

	params := helpers.GetRequestParams(c)

	listTags, total, err := h.TagUsecase.FetchTag(ctx, &params)

	if err != nil {
		return err
	}

	res := helpers.Paginate(c, listTags, total, params)

	return c.JSON(http.StatusOK, res)
}
