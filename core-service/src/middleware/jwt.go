package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// JWT will handle the JWT middleware
func (m *GoMiddleware) JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get the token from the header

		JWT_SIGNATURE_KEY := []byte("-----BEGIN PUBLIC KEY-----\n" +
			viper.GetString(`KEYCLOAK_PEM`) +
			"\n-----END PUBLIC KEY-----\n")

		authorizationHeader := c.Request().Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		key, err := jwt.ParseRSAPublicKeyFromPEM(JWT_SIGNATURE_KEY)
		if err != nil {
			log.Printf("validate: parse key: %v", err)
			return err
		}

		token, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
			}

			return key, nil
		})

		if err != nil {
			log.Printf("validate: parse token: %v", err)
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		_, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			log.Printf("validate: token invalid: %v", err)
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		return next(c)
	}
}
