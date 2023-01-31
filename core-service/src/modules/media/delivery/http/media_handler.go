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
	r.DELETE("/media/delete", handler.Delete)
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
		"featured-program",
		"informations",
		"pop-up-banners",
		"infographic-banners",
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

// Delete will delete the object S3 based on key name
func (h *MediaHandler) Delete(c echo.Context) (err error) {
	ctx := c.Request().Context()
	reqBody := new(domain.DeleteMediaRequest)
	if err = c.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.MUsecase.Delete(ctx, reqBody)
	if err != nil {
		logrus.Fatal(err)
	}

	res := domain.MessageResponse{
		Message: "succesfully deleted file from cloud storage.",
	}

	return c.JSON(http.StatusOK, res)
}
