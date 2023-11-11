package framework

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

func (e EchoTemplates) ServerTest() []byte {
	return MakeHTTPServerTest()
}

func (e EchoTemplates) RoutesTest() []byte {
	return MakeEchoRoutesTest()
}

// MakeEchoRoutes returns a byte slice that represents 
// the internal/server/routes.go file when using Echo.
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

func MakeEchoRoutesTest() []byte {
	return []byte(`package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type Server struct{}

// TestRegisterRoutes tests if routes are registered correctly
func TestRegisterRoutes(t *testing.T) {
	e := echo.New()
	s := &Server{}
	h := s.RegisterRoutes()
	e.Server.Handler = h

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	e.Server.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	expected := ` + "`{\"message\":\"Hello World\"}`" + `
	assert.Equal(t, expected, rec.Body.String())
}

// TestHelloWorldHandler tests the response from the helloWorldHandler
func TestHelloWorldHandler(t *testing.T) {
	e := echo.New()
	s := &Server{}
	e.GET("/", s.helloWorldHandler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, s.helloWorldHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		expected := ` + "`{\"message\":\"Hello World\"}`" + `
		assert.Equal(t, expected, rec.Body.String())
	}
}
`)
}