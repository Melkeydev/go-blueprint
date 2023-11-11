package framework

// StandardLibTemplate contains the methods used for building
// an app that uses [net/http]
type StandardLibTemplate struct{}

func (s StandardLibTemplate) Main() []byte {
	return MainTemplate()
}

func (s StandardLibTemplate) Server() []byte {
	return MakeHTTPServer()
}

func (s StandardLibTemplate) Routes() []byte {
	return MakeHTTPRoutes()
}

func (s StandardLibTemplate) ServerTest() []byte {
	return MakeHTTPServerTest()
}

func (s StandardLibTemplate) RoutesTest() []byte {
	return MakeHTTPRoutesTest()
}
// MakeHTTPServer returns a byte slice that represents
// the default internal/server/server.go file.
func MakeHTTPServer() []byte {
	return []byte(`package server

import (
	"fmt"
	"net/http"
	"time"
)

var port = 8080

type Server struct {
	port int
}

func NewServer() *http.Server {

	NewServer := &Server{
		port: port,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
`)
}

// MakeHTTPRoutes returns a byte slice that represents
// the internal/server/routes.go file when using net/http
func MakeHTTPRoutes() []byte {
	return []byte(`package server

import (
	"net/http"
	"encoding/json"
	"log"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handler)

	return mux
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Write(jsonResp)
}
`)
}

func MakeHTTPServerTest() []byte {
	return []byte(`package server

import (
	"net/http"
	"testing"
	"time"
	"fmt"
)

// TestNewServer checks if a new server is created with correct configurations
func TestNewServer(t *testing.T) {
	server := NewServer()

	// Check if server is not nil
	if server == nil {
		t.Errorf("NewServer returned nil")
	}

	// Check port configuration
	expectedAddr := fmt.Sprintf(":%d", port)
	if server.Addr != expectedAddr {
		t.Errorf("Expected address %s, got %s", expectedAddr, server.Addr)
	}

	// Check if Handler is set
	if server.Handler == nil {
		t.Errorf("Handler is not set")
	}

	// Check timeout configurations
	if server.IdleTimeout != time.Minute {
		t.Errorf("Expected IdleTimeout to be 1 minute, got %s", server.IdleTimeout)
	}

	if server.ReadTimeout != 10*time.Second {
		t.Errorf("Expected ReadTimeout to be 10 seconds, got %s", server.ReadTimeout)
	}

	if server.WriteTimeout != 30*time.Second {
		t.Errorf("Expected WriteTimeout to be 30 seconds, got %s", server.WriteTimeout)
	}
}
`)
}

func MakeHTTPRoutesTest() []byte {
	return []byte(`package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestRegisterRoutes checks if routes are registered correctly
func TestRegisterRoutes(t *testing.T) {
	s := &Server{}
	mux := s.RegisterRoutes()

	if mux == nil {
		t.Fatal("Expected non-nil mux, got nil")
	}

	// Test if the root route is registered
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := ` + "`{\"message\":\"Hello World\"}`" + `
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// TestHandlerResponse checks the response from the handler
func TestHandlerResponse(t *testing.T) {
	s := &Server{}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.handler)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, req)

	expected := map[string]string{"message": "Hello World"}
	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if response["message"] != expected["message"] {
		t.Errorf("handler returned unexpected body: got %v want %v", response["message"], expected["message"])
	}
}
`)}