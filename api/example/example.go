package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/eucatur/go-toolbox/api"
)

func main() {
	server := api.Make()

	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	api.Run()
}
