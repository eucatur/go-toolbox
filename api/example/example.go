package main

import (
	"net/http"

	"github.com/eucatur/go-toolbox/api"
	"github.com/labstack/echo"
)

func main() {
	server := api.Make()

	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	api.Run()
}
