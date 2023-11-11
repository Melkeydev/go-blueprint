package framework

// FiberTemplates contains the methods used for building
// an app that uses [github.com/gofiber/fiber]
type FiberTemplates struct{}

func (f FiberTemplates) Main() []byte {
	return MakeFiberMain()
}
func (f FiberTemplates) Server() []byte {
	return MakeFiberServer()
}

func (f FiberTemplates) Routes() []byte {
	return MakeFiberRoutes()
}

func (f FiberTemplates) RoutesWithDB() []byte {
	return MakeFiberRoutesWithDB()
}

// MakeFiberServer returns a byte slice that represents
// the internal/server/server.go file when using Fiber.
func MakeFiberServer() []byte {
	return []byte(`package server

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

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

func MakeFiberRoutesWithDB() []byte {
	return []byte(`package server

import (
	"github.com/gofiber/fiber/v2"
	"{{.ProjectName}}/internal/database"
)

type healthHandler struct {
	s database.Service
}

func (s *FiberServer) RegisterFiberRoutes() {
	h := NewHealthHandler()
	s.App.Get("/", s.helloWorldHandler)
	s.App.Get("/health", h.healthHandler)
}

func (s *FiberServer) helloWorldHandler(c *fiber.Ctx) error {
	resp := map[string]string{
		"message": "Hello World",
	}
	return c.JSON(resp)
}

func NewHealthHandler() *healthHandler {
	return &healthHandler{
		s: database.New(),
	}
}

func (h *healthHandler) healthHandler(c *fiber.Ctx) error {
	return c.JSON(h.s.Health())
}
`)
}

// MakeHTTPRoutes returns a byte slice that represents
// the cmd/api/main.go file when using Fiber.
func MakeFiberMain() []byte {
	return []byte(`package main

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
