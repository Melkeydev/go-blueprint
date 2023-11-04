package template

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

func MakeEchoRoutes() []byte {
	return []byte(`package server

	import (
		"net/http"
	
		"github.com/labstack/echo/v4"
		"github.com/labstack/echo/v4/middleware"
	)
func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.helloWorldHandler)

	return e
}

func (s *Server) helloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}
`)
}
