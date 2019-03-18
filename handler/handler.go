// It's a lib for some utility functions to help with handler in Echo framework
package handler

import (
	"github.com/eucatur/go-toolbox/log"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

var (
	MESSAGE  = "message"
	validate *validator.Validate
)

type Handler struct {
	Message string `json:"message" form:"message" query:"message"`
}

// Like the name Validade and bind one struct with the validador golang lib
func BindAndValidate(c echo.Context, object interface{}) (err error) {
	validate = validator.New()

	if err := c.Bind(object); err != nil {
		return c.JSON(422, &Handler{err.Error()})
	}

	if err := validate.Struct(object); err != nil {
		log.Error(err)
		return c.JSON(422, &Handler{err.Error()})
	}

	return
}
