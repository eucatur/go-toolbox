package handler

import (
	"github.com/labstack/echo"
)

var claims map[string]interface{}

// Middleware for bind and validate
func MiddlewareBindAndValidate(object interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {

			return BindAndValidate(c, object)
		}
	}
}
