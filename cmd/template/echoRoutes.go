package template

import "fmt"

// EchoTemplates contains the methods used for building
// an app that uses [github.com/labstack/echo]
type EchoTemplates struct{}

func (e EchoTemplates) Main() []byte {
	return MainTemplate()
}
func (e EchoTemplates) Server() []byte {
	return MakeHTTPServer()
}

func (e EchoTemplates) Routes() []byte {
	return MakeEchoRoutes()
}

// MakeEchoRoutes returns a byte slice that represents 
// the internal/server/routes.go file when using Echo.
func MakeEchoRoutes() []byte {
	return []byte(fmt.Sprintf(`package server

	import (
		"errors"
		"log"
		"net/http"
	
		"github.com/labstack/echo/v4"
		"github.com/labstack/echo/v4/middleware"
		"nhooyr.io/websocket"
	)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.helloWorldHandler)
	e.GET("/ws", s.pingPongWebsocketHandler)

	return e
}

func (s *Server) helloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) pingPongWebsocketHandler(c echo.Context) error {
	w := c.Response().Writer
	r := c.Request()
	%s
}
`, websocketTemplate()))
}
