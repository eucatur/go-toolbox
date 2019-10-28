// It's a lib for some utility functions to help with handler in Echo framework
package handler

import (
	"reflect"

	"github.com/eucatur/go-toolbox/log"

	"github.com/eucatur/go-toolbox/validator"
	"github.com/labstack/echo"
	"github.com/mcuadros/go-defaults"
)

const PARAMETERS = "parameters"

var (
	MESSAGE = "message"
)

type Handler struct {
	Message string `json:"message" form:"message" query:"message"`
}

// Like the name Validade and bind one struct with the validador golang lib
func BindAndValidate(c echo.Context, obj interface{}, args ...interface{}) (err error) {
	obj = reflect.ValueOf(obj).Elem().Interface()
	obj = reflect.New(reflect.TypeOf(obj)).Interface()

	err = c.Bind(obj)
	if err != nil {
		return
	}

	var options []string
	if len(args) > 0 {
		group, ok := args[0].(string)
		if ok {
			options = append(options, group)
		}
	}

	vErr := validator.Validate(obj, options...)
	if vErr != nil {
		log.Error(vErr)
		err = c.JSON(422, vErr)
		if err != nil {
			return
		}
		return vErr
	}

	defaults.SetDefaults(obj)

	c.Set(PARAMETERS, obj)

	return
}

// Ok
func Ok(c echo.Context, b interface{}) error {
	return c.JSON(200, b)
}

// Created ...
func Created(b interface{}) (err error) {
	switch value := b.(type) {
	case int, int32, int64:
		return echo.NewHTTPError(201, echo.Map{"id": value})
	case []struct{}:
		return echo.NewHTTPError(201, value)
	}
	return
}

// Error ...
func Error(e error) error {
	_, ok := e.(*echo.HTTPError)
	if !ok {
		log.Println(e)
		return echo.NewHTTPError(500, "Erro interno no servidor.")
	}

	return e
}

// Message
func Message(c echo.Context, m string) error {
	return c.JSON(200, &Handler{m})
}

// ErrorMessage
func ErrorMessage(c echo.Context, m string) error {
	return c.JSON(400, &Handler{m})
}
