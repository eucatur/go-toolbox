// It's a lib for some utility functions to help with handler in Echo framework
package handler

import (
	"reflect"

	"github.com/eucatur/go-toolbox/log"

	"github.com/eucatur/go-toolbox/validator"
	"github.com/labstack/echo"
)

const PARAMETERS = "parameters"

var (
	MESSAGE = "message"
)

type Handler struct {
	Message string `json:"message" form:"message" query:"message"`
}

// Like the name Validade and bind one struct with the validador golang lib
func BindAndValidate(c echo.Context, obj interface{}) (err error) {
	obj = reflect.ValueOf(obj).Elem().Interface()
	obj = reflect.New(reflect.TypeOf(obj)).Interface()

	if err := c.Bind(obj); err != nil {
		log.Error(err)
		return c.JSON(422, &Handler{err.Error()})
	}

	if err := validator.Validate(obj); err != nil {
		log.Error(err)
		return c.JSON(422, err)
	}

	// defaults.SetDefaults(obj)

	c.Set(PARAMETERS, obj)

	return
}

// Ok
func Ok(c echo.Context, b interface{}) error {
	return c.JSON(200, b)
}

// Error
func Error(e error) error {
	return echo.NewHTTPError(400, e)
}

// Message
func Message(c echo.Context, m string) error {
	return c.JSON(200, &Handler{m})
}

// ErrorMessage
func ErrorMessage(c echo.Context, m string) error {
	return c.JSON(400, &Handler{m})
}
