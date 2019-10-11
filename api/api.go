package api

import (
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/eucatur/go-toolbox/log"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	port       *string
	debug      *bool
	echoServer *echo.Echo
)

func init() {
	port = flag.String("port", "9000", "port for the service HTTP")
	debug = flag.Bool("debug", false, "mod of the debug")
}

func Make() *echo.Echo {
	flag.Parse()

	echoServer = echo.New()

	// Esconde o cabeçalho do Echo
	echoServer.HideBanner = true

	echoServer.Use(middleware.CORS())
	echoServer.Use(middleware.Recover())

	if *debug {
		echoServer.Debug = true
		echoServer.Use(middleware.Logger())
		log.EnableDebug(true)
	}

	return echoServer
}

// Provides the instance of Echo
func ProvideEchoInstance(task func(e *echo.Echo)) {
	task(echoServer)
}

func Run() {
	// For Heroku Work
	porta := os.Getenv("PORT")

	if porta == "" {
		porta = *port
	}

	echoServer.Logger.Fatal(echoServer.Start(":" + porta))
}

func Use(middleware ...echo.MiddlewareFunc) {
	echoServer.Use(middleware...)
}

func UseCustomHTTPErrorHandler() {
	echoServer.HTTPErrorHandler = CustomHTTPErrorHandler
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	if err == nil {
		return
	}

	defer log.File(time.Now().Format("errors/2006/01/02/15h.log"), err.Error())

	var (
		code = http.StatusInternalServerError
		msg  interface{}
	)

	msg = echo.Map{"message": "Erro interno no servidor."}

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD { // Issue #608
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, msg)
		}
	}
}
