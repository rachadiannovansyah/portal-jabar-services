package http

import (
	"bytes"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
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
	err = h.MUsecase.Store(ctx, file, buf)

	if err != nil {
		logrus.Fatal(err)
	}

	return c.JSON(http.StatusCreated, "File Uploaded successfully.")
}
