package http

import (
	"net/http"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

// PublicDocumentArchiveHandler is represented by domain.DocumentArchiveUsecase
type PublicDocumentArchiveHandler struct {
	DocumentArchiveUcase domain.DocumentArchiveUsecase
}

// NewDocumentArchiveHandler will initialize the document archive endpoint
func NewPublicDocumentArchiveHandler(p *echo.Group, us domain.DocumentArchiveUsecase) {
	handler := &PublicDocumentArchiveHandler{
		DocumentArchiveUcase: us,
	}
	p.GET("/document-archives", handler.Fetch)
}

// Fetch for fetching document archive data
func (h *PublicDocumentArchiveHandler) Fetch(c echo.Context) error {
	// init request by context
	ctx := c.Request().Context()
	params := helpers.GetRequestParams(c)
	params.Filters = map[string]interface{}{
		"category": helpers.RegexReplaceString(c, c.QueryParam("cat"), ""),
	}

	// getting data from usecase
	listDoc, total, err := h.DocumentArchiveUcase.Fetch(ctx, &params)
	if err != nil {
		return err
	}

	// preparing response
	listDocRes := []domain.ListDocumentArchive{}
	copier.Copy(&listDocRes, &listDoc)

	res := helpers.Paginate(c, listDocRes, total, params)

	return c.JSON(http.StatusOK, res)
}
