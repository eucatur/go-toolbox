package handler

import (
	"github.com/labstack/echo/v4"
)

// Middleware for bind and validate
func MiddlewareBindAndValidate(object interface{}, args ...interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			err = BindAndValidate(c, object, args...)
			if err != nil {
				return err
			}
			return next(c)
		}
	}
}
