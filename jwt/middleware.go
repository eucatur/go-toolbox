package jwt

import (
	"github.com/labstack/echo"
)

var claims map[string]interface{}

// Middleware for Echo framework
func Middleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {

			tokenRequest := c.Request().Header.Get(HEADER)

			if tokenRequest == "" {
				return &echo.HTTPError{Code: 401, Message: "Header Authorization não informado"}
			}

			claims, err = VerifyTokenAndGetClaims(tokenRequest, secret)

			if err != nil {
				return &echo.HTTPError{Code: 401, Message: "Token inválido"}
			}

			c.Set(IDENTICATION, claims[IDENTICATION])

			return next(c)
		}
	}
}
