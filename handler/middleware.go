package handler

import (
	"github.com/labstack/echo"
)

// Middleware for bind and validate
func MiddlewareBindAndValidate(object interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			err = BindAndValidate(c, object)
			if err != nil {
				return err
			}
			return next(c)
		}
	}
}
