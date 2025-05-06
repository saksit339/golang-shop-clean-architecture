package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/saksit339/isekai-shop-api-tutorial/config"
	"gorm.io/gorm"
)

var once sync.Once
var server *echoServer

type echoServer struct {
	app  *echo.Echo
	db   *gorm.DB
	conf *config.Config
}

func NewEchoServer(conf *config.Config, db *gorm.DB) *echoServer {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	once.Do(func() {
		server = &echoServer{
			app:  e,
			db:   db,
			conf: conf,
		}
	})
	return server
}

func (s *echoServer) Start() {

	//middleware
	corsMiddleware := getCoreMiddleware(s.conf.Server.AllowOrigins)
	bodyLimitMiddleware := getBodyLimitMiddleware(s.conf.Server.BodyLimit)
	timeoutmiddleware := getTimeoutMiddleware(s.conf.Server.Timeout)

	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())
	s.app.Use(corsMiddleware)
	s.app.Use(bodyLimitMiddleware)
	s.app.Use(timeoutmiddleware)

	s.app.GET("/health", s.healthCheck)

	quitch := make(chan os.Signal, 1)
	signal.Notify(quitch, syscall.SIGINT, syscall.SIGTERM)
	go s.gracefulShutdown(quitch)

	s.HttpListening()
}

func (s *echoServer) HttpListening() {
	url := fmt.Sprintf("%d", s.conf.Server.Port)

	if err := s.app.Start(":" + url); err != nil && err != http.ErrServerClosed {
		s.app.Logger.Fatal(err.Error())
	}
}

func (s *echoServer) gracefulShutdown(quite chan os.Signal) {
	<-quite
	s.app.Logger.Info("Shutting down server...")
	ctx := context.Background()
	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatal(err.Error())
	}
}

func (s *echoServer) healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Ok")
}

func getTimeoutMiddleware(timeout time.Duration) echo.MiddlewareFunc {
	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		Timeout:      timeout * time.Second,
		ErrorMessage: "Request timeout",
	})
}

func getCoreMiddleware(allowOrigin []string) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: allowOrigin,
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	})
}

func getBodyLimitMiddleware(bodyLimit string) echo.MiddlewareFunc {
	return middleware.BodyLimit(bodyLimit)
}
