package framework

// GorillaTemplates contains the methods used for building
// an app that uses [github.com/gorilla/mux]
type GorillaTemplates struct{}

func (g GorillaTemplates) Main() []byte {
	return MainTemplate()
}
func (g GorillaTemplates) Server() []byte {
	return MakeHTTPServer()
}
func (g GorillaTemplates) Routes() []byte {
	return MakeGorillaRoutes()
}
func (g GorillaTemplates) ServerTest() []byte {
	return MakeHTTPServerTest()
}
func (g GorillaTemplates) RoutesTest() []byte {
	return MakeGorillaRoutesTest()
}

// MakeGorillaRoutes returns a byte slice that represents 
// the internal/server/routes.go file when using gorilla/mux.
func MakeGorillaRoutes() []byte {
	return []byte(`package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", s.helloWorldHandler)

	return r
}

func (s *Server) helloWorldHandler(w http.ResponseWriter, r *http.Request) {
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

func MakeGorillaRoutesTest() []byte {
	return []byte(`package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestRegisterRoutes tests if routes are registered correctly
func TestRegisterRoutes(t *testing.T) {
	s := &Server{}
	handler := s.RegisterRoutes()

	// Creating a request to pass to the handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Recording the response
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := ` + "`{\"message\":\"Hello World\"}`" + `
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// TestHelloWorldHandler tests the response from the helloWorldHandler
func TestHelloWorldHandler(t *testing.T) {
	s := &Server{}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.helloWorldHandler)

	// Creating a request to pass to the handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Recording the response
	handler.ServeHTTP(rr, req)

	// Check the response body
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
`)
}