package template

func MakeFiberServer() []byte {
	return []byte(`

package server

import "github.com/gofiber/fiber/v2"

type FiberServer struct {
	*fiber.App
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(),
	}

	return server
}

`)
}

func MakeFiberRoutes() []byte {
	return []byte(`

package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.helloWorldHandler)
}

func (s *FiberServer) helloWorldHandler(c *fiber.Ctx) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(resp)
}
`)
}

func MakeFiberMain() []byte {
	return []byte(`

package main

import (
	"{{.ProjectName}}/internal/server"
)
	
func main() {

	server := server.New()

	server.RegisterFiberRoutes()

	err := server.Listen(":8080")
	if err != nil {
		panic("cannot start server")
	}
}
`)
}
