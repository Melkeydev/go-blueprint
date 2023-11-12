package template

// GinTemplates contains the methods used for building
// an app that uses [github.com/gin-gonic/gin]
type GinTemplates struct{}

func (g GinTemplates) Main() []byte {
	return MainTemplate()
}

func (g GinTemplates) Server() []byte {
	return MakeHTTPServer()
}

func (g GinTemplates) Routes() []byte {
	return MakeGinRoutes()
}

func (g GinTemplates) TestHandler() []byte {
    return MakeGinTestHandler()
}

// MakeGinRoutes returns a byte slice that represents
// the internal/server/routes.go file when using Gin.
func MakeGinRoutes() []byte {
	return []byte(`package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", s.HelloWorldHandler)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

`)
}

func MakeGinTestHandler() []byte {
	return []byte(`
package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"{{.ProjectName}}/internal/server"
)

func TestHelloWorldHandler(t *testing.T) {
	s := &server.Server{}
	r := gin.New()
	r.GET("/", s.HelloWorldHandler)

	// Create a test HTTP request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Serve the HTTP request
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

    // Check the response body
    expected := "{\"message\":\"Hello World\"}"
    if rr.Body.String() != expected {
        t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}
    `)
}
