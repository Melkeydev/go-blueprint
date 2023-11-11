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

func (f FiberTemplates) ServerTest() []byte {
	return MakeFiberServerTest()
}

func (f FiberTemplates) RoutesTest() []byte {
	return MakeFiberRoutesTest()
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

func MakeFiberServerTest() []byte {
	return []byte(`package server

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// TestNew tests the New function for creating a new FiberServer instance
func TestNew(t *testing.T) {
	server := New()

	// Check if the server is not nil
	assert.NotNil(t, server)

	// Check if the server is of type *FiberServer
	assert.IsType(t, &FiberServer{}, server)

	// Check if the App field is correctly initialized
	assert.IsType(t, &fiber.App{}, server.App)
}
`)
}

func MakeFiberRoutesTest() []byte {
	return []byte(`package server

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// TestHelloWorldHandler tests the response from the helloWorldHandler
func TestHelloWorldHandler(t *testing.T) {
	// Use the New function to create a new FiberServer instance
	server := New()

	server.RegisterFiberRoutes()

	req := httptest.NewRequest("GET", "/", nil)
	resp, err := server.App.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Read the response body using io.ReadAll
	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)

	// Convert the body to string for assertion
	bodyString := string(body)
	expectedBody := ` + "`{\"message\":\"Hello World\"}`" + `
	assert.Equal(t, expectedBody, bodyString)

	// Close the response body
	err = resp.Body.Close()
	assert.Nil(t, err)
}
`)
}
