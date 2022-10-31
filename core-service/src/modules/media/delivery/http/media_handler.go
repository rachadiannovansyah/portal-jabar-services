package http

import (
	"bytes"
	"io"
	"net/http"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// MediaHandler ...
type MediaHandler struct {
	MUsecase domain.MediaUsecase
}

// NewMediaHandler will create a new MediaHandler
func NewMediaHandler(e *echo.Group, r *echo.Group, mu domain.MediaUsecase) {
	handler := &MediaHandler{
		MUsecase: mu,
	}
	r.POST("/media/upload", handler.Store)
}

// Store will store the feedback by given request body
func (h *MediaHandler) Store(c echo.Context) (err error) {
	// validate for certain allowed bucket name of domain
	domain := c.QueryParam("domain")
	domainBucketName := []string{
		"news",
		"events",
		"public-service",
		"units",
		"featured_program",
		"informations",
	}
	domainExists, domainIndex := helpers.InArray(domain, domainBucketName)
	domain = ""
	if domainExists {
		domain = domainBucketName[domainIndex] + "/"
	}

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, src); err != nil {
		logrus.Println(err)
		return err
	}

	ctx := c.Request().Context()
	res, err := h.MUsecase.Store(ctx, file, buf, domain)

	if err != nil {
		logrus.Fatal(err)
	}

	return c.JSON(http.StatusCreated, res)
}
