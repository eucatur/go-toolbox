package api

import (
	"flag"
	"os"

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
