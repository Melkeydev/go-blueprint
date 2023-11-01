package template

type EchoTemplates struct{}

func (e EchoTemplates) Main() []byte {
	return MakeEchoMain()
}
func (e EchoTemplates) Server() []byte {
	return MakeEchoServer()
}

func (e EchoTemplates) Routes() []byte {
	return MakeEchoRoutes()
}

func MakeEchoServer() []byte {
	return []byte(`package server

import "github.com/labstack/echo/v4"

type EchoServer struct {
	*echo.Echo
}

func New() *EchoServer {
	server := &EchoServer{
		Echo: echo.New(),
	}
	
	return server
}
`)
}

func MakeEchoRoutes() []byte {
	return []byte(`package server

	import (
		"net/http"
	
		"github.com/labstack/echo/v4"
	)
func (s *EchoServer) RegisterEchoRoutes() {
	s.Echo.GET("/", s.helloWorldHandler)
}

func (s *EchoServer) helloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}
`)
}

func MakeEchoMain() []byte {
	return []byte(`package main
	
	import (
		"{{.ProjectName}}/internal/server"
	)

	func main() {
		server := server.New()
		server.RegisterEchoRoutes()
		server.Logger.Fatal(server.Start(":8080"))
	}
`)
}
