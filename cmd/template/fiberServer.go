package template

import (
	_ "embed"
)

//go:embed files/routes/fiber.go.tmpl
var fiberRoutesTemplate []byte

//go:embed files/server/fiber.go.tmpl
var fiberServerTemplate []byte

//go:embed files/main/fiber_main.go.tmpl
var fiberMainTemplate []byte

// FiberTemplates contains the methods used for building
// an app that uses [github.com/gofiber/fiber]
type FiberTemplates struct{}

func (f FiberTemplates) Main() []byte {
	return fiberMainTemplate
}
func (f FiberTemplates) Server() []byte {
	return fiberServerTemplate
}

func (f FiberTemplates) Routes() []byte {
	return fiberRoutesTemplate
}

// MakeFiberServer returns a byte slice that represents 
// the internal/server/server.go file when using Fiber.
func MakeFiberServer() []byte {
	return []byte(`package server

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

// MakeFiberRoutes returns a byte slice that represents 
// the internal/server/routes.go file when using Fiber.
func MakeFiberRoutes() []byte {
	return []byte(`package server

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

// MakeHTTPRoutes returns a byte slice that represents 
// the cmd/api/main.go file when using Fiber.
func MakeFiberMain() []byte {
	return []byte(`package main

import (
	"os"
	"{{.ProjectName}}/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	server := server.New()

	server.RegisterFiberRoutes()

	err := server.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		panic("cannot start server")
	}
}
`)
}
