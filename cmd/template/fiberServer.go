package template

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

func (f FiberTemplates) TestHandler() []byte {
    return MakeFiberTestHandler()
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
	s.App.Get("/", s.HelloWorldHandler)
}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
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
func MakeFiberTestHandler() []byte {
	return []byte(`

package tests

import (
	"io"
	"net/http"
	"{{.ProjectName}}/internal/server"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestHandler(t *testing.T) {
	// Create a Fiber app for testing
	app := fiber.New()

	// Inject the Fiber app into the server
	s := &server.FiberServer{App: app}

	// Define a route in the Fiber app
	app.Get("/", s.HelloWorldHandler)

	// Create a test HTTP request
	req,err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatalf("error creating request. Err: %v", err)
    }

	// Perform the request
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}

	// Your test assertions...
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}

	expected := "{\"message\":\"Hello World\"}"

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        t.Fatalf("error reading response body. Err: %v", err)
    }
    if expected != string(body) {
        t.Errorf("expected response body to be %v; got %v", expected, string(body))
    }
}
`)
}
