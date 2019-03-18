package jwt

import (
	"strings"

	"github.com/labstack/echo"
)

var claims map[string]interface{}

// Middleware for Echo framework
func MiddlewareForHeader(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {

			tokenRequest := c.Request().Header.Get(HEADER)

			if tokenRequest == "" {
				return &echo.HTTPError{Code: 401, Message: "Header Authorization não informado"}
			}

			claims, err = VerifyTokenAndGetClaims(strings.Replace(tokenRequest, "Bearer ", "", -1), secret)

			if err != nil {
				return &echo.HTTPError{Code: 401, Message: "Token inválido"}
			}

			c.Set("claims", claims)

			return next(c)
		}
	}
}
