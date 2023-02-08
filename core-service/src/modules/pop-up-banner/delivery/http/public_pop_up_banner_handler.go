package http

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
)

// PublicPopUpBanner ...
type PublicPopUpBanner struct {
	PUsecase domain.PopUpBannerUsecase
	apm      *utils.Apm
}

// NewPublicPopUpBanner will create a new PublicPopUpBanner
func NewPublicPopUpBanner(e *echo.Group, ucase domain.PopUpBannerUsecase, apm *utils.Apm) {
	handler := &PublicPopUpBanner{
		PUsecase: ucase,
		apm:      apm,
	}

	e.GET("/pop-up-banners/live", handler.LiveBanner)
}

// LiveBanner will get pop up banner with active status and is_live is 1
func (h *PublicPopUpBanner) LiveBanner(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := h.PUsecase.LiveBanner(ctx)
	if err != nil {
		return c.JSON(helpers.GetStatusCode(err), helpers.ResponseError{Message: err.Error()})
	}

	// re-presenting responses
	res := domain.DetailPopUpBannerResponse{
		ID:          data.ID,
		Title:       data.Title,
		ButtonLabel: data.ButtonLabel,
		Link:        data.Link,
		Status:      data.Status,
		Duration:    data.Duration,
		StartDate:   data.StartDate,
		EndDate:     data.EndDate,
		UpdateAt:    data.UpdatedAt,
	}

	helpers.GetObjectFromString(data.Image.String, &res.Image)

	return c.JSON(http.StatusOK, &domain.ResultData{Data: &res})
}
