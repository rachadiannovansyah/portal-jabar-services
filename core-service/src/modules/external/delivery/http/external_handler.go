package http

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// NewExternalHandler will initialize the event endpoint
func NewExternalHandler(e *echo.Group, r *echo.Group) {
	r.GET("/public/link-checker", CheckLinkHandler)
}

// Fetch will get events data
func CheckLinkHandler(c echo.Context) error {
	link := c.QueryParam("link")
	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		link = "http://" + link
	}

	resp, err := http.Get(link)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"link":  link,
			"valid": false,
			"ssl":   false,
		})
	}
	defer resp.Body.Close()

	valid := true
	ssl := false

	if resp.TLS != nil && len(resp.TLS.PeerCertificates) > 0 {
		for _, cert := range resp.TLS.PeerCertificates {
			if cert.IsCA {
				ssl = true
				break
			}
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"link":  link,
		"valid": valid,
		"ssl":   ssl,
	})
}
