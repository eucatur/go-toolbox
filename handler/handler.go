// It's a lib for some utility functions to help with handler in Echo framework
package handler

import (
	"github.com/eucatur/go-toolbox/log"

	"bitbucket.org/eucatur/go-toolbox/lib/validator"
	"github.com/labstack/echo"
	defaults "github.com/mcuadros/go-defaults"
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
	if err := c.Bind(obj); err != nil {
		return c.JSON(422, &Handler{err.Error()})
	}

	if err := validator.Validate(obj); err != nil {
		log.Error(err)
		return c.JSON(422, err)
	}

	defaults.SetDefaults(obj)

	c.Set(PARAMETERS, obj)

	return
}
