package api

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eucatur/go-toolbox/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func getPort() string {
	// For Heroku Work
	porta := os.Getenv("PORT")

	if porta == "" {
		porta = *port
	}

	return porta
}

func Run() {

	porta := getPort()

	echoServer.Logger.Fatal(echoServer.Start(":" + porta))
}

func Use(middleware ...echo.MiddlewareFunc) {
	echoServer.Use(middleware...)
}

func UseCustomHTTPErrorHandler() {
	echoServer.HTTPErrorHandler = CustomHTTPErrorHandler
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		msg  interface{}
	)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
		if he.Internal != nil {
			msg = fmt.Sprintf("%v, %v", err, he.Internal)
		}
	} else {
		msg = http.StatusText(code)
	}

	if _, ok := msg.(string); ok {
		msg = echo.Map{"message": msg}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD { // Issue #608
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, msg)
		}
		if err != nil {
			c.Echo().Logger.Error(err)
		}
	} else {
		log.File(time.Now().Format("errors/2006/01/02/15h.log"), err.Error())
	}
}

func RunWithGracefulShutdown(ctx context.Context, secondWaitToExit int) {

	go runAsync()

	settingGracefulShutdown(ctx, secondWaitToExit)

}

func settingGracefulShutdown(ctx context.Context, secondWaitToExit int) {

	ctxGracefulShutdown, cancel := gracefulShutdown(ctx, secondWaitToExit)

	if echoServer == nil {
		log.Println("Servidor da api não instanciado")
	}

	if err := echoServer.Shutdown(ctxGracefulShutdown); err != nil {
		log.Println(fmt.Sprintf("Falha ao encerrar o servidor de API. Detalhes: %s", err.Error()))
	}

	defer cancel()

}

func runAsync() {

	porta := getPort()

	if err := echoServer.Start(":" + porta); err != nil && err != http.ErrServerClosed {
		echoServer.Logger.Fatal(err)
	}
}

func gracefulShutdown(ctx context.Context, secondsWaitToExit int) (retCtx context.Context, cancel context.CancelFunc) {

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)

	<-stop

	timeToExitApp := time.Duration(secondsWaitToExit) * time.Second

	log.Println(fmt.Sprintf("Se houver algum processamento em andamento a aplicação irá aguardar por mais %s até se encerrar", timeToExitApp.String()))

	retCtx, cancel = context.WithTimeout(ctx, timeToExitApp)

	return

}
