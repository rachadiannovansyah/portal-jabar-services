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

	// throw if keyword is empty and saving our repo!
	if len(params.Keyword) == 0 {
		return c.JSON(http.StatusOK, []string{})
	}

	listTags, _, err := h.TagUsecase.FetchTag(ctx, &params)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, listTags)
}
