package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
	"github.com/labstack/echo/v4"
)

func hasPermission(permission string, permissions []string) bool {
	for _, p := range permissions {
		if p == permission {
			return true
		}
	}
	return false
}

func CheckPermission(permission string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			au := helpers.GetAuthenticatedUser(c)
			if au.Role.ID == domain.RoleSuperAdmin {
				return next(c)
			}

			if !hasPermission(permission, au.Permissions) {
				return c.JSON(http.StatusForbidden, domain.NewErrResponse(errors.New("permission denied")))
			}

			return next(c)
		}
	}
}

func VerifyCache() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cached := helpers.GetCache(c)

			if cached != "" {
				var cacheRes domain.Cache

				err := json.Unmarshal([]byte(cached), &cacheRes)
				if err != nil {
					fmt.Println(err)
				}

				return c.JSON(http.StatusOK, &cacheRes)
			}

			return next(c)
		}
	}
}
