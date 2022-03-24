package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"

	"github.com/labstack/echo/v4"
)

// JWT will handle the JWT middleware
func (m *GoMiddleware) JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get the token from the header

		authorizationHeader := c.Request().Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		// Initialize a new instance of `Claims`
		claims := &domain.JwtCustomClaims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			fmt.Println("JWT_KEY", m.JWTKey)
			return m.JWTKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid || !tkn.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if !tkn.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		c.Set("auth:user", claims)

		return next(c)
	}
}
