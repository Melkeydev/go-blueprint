package template

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

func (e EchoTemplates) TestHandler() []byte {
    return MakeEchoTestHandler()
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

	e.GET("/", s.HelloWorldHandler)

	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}
`)
}

func MakeEchoTestHandler() []byte {
    return []byte(`

package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"{{.ProjectName}}/internal/server"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)
	s := &server.Server{}

	// Assertions
	if err := s.HelloWorldHandler(c); err != nil {
		t.Errorf("handler() error = %v", err)
		return
	}

	if resp.Code != http.StatusOK {
		t.Errorf("handler() wrong status code = %v", resp.Code)
		return
	}

	expected := map[string]string{"message": "Hello World"}
	var actual map[string]string

	// Decode the response body into the actual map
	if err := json.NewDecoder(resp.Body).Decode(&actual); err != nil {
		t.Errorf("handler() error decoding response body: %v", err)
		return
	}

	// Compare the decoded response with the expected value
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("handler() wrong response body. expected = %v, actual = %v", expected, actual)
		return
	}
}
`)
}
