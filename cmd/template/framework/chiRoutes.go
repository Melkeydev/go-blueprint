package framework 

// ChiTemplates contains the methods used for building
// an app that uses [github.com/go-chi/chi]
type ChiTemplates struct{}

func (c ChiTemplates) Main() []byte {
	return MainTemplate()
}

func (c ChiTemplates) Server() []byte {
	return MakeHTTPServer()
}

func (c ChiTemplates) Routes() []byte {
	return MakeChiRoutes()
}

func (c ChiTemplates) ServerTest() []byte {
	return MakeHTTPServerTest()
}

func (c ChiTemplates) RoutesTest() []byte {
	return MakeChiRoutesTest()
}

// MakeChiRoutes returns a byte slice that represents 
// the internal/server/routes.go file when using Chi.
func MakeChiRoutes() []byte {
	return []byte(`package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.helloWorldHandler)

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

func MakeChiRoutesTest() []byte {
	return []byte(`package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

type Server struct{}

// TestRegisterRoutes tests if routes are registered correctly
func TestRegisterRoutes(t *testing.T) {
	s := &Server{}
	r := s.RegisterRoutes()
	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	defer res.Body.Close()
	var resp map[string]string
	err = json.NewDecoder(res.Body).Decode(&resp)
	assert.NoError(t, err)
	assert.Equal(t, "Hello World", resp["message"])
}

// TestHelloWorldHandler tests the response from the helloWorldHandler
func TestHelloWorldHandler(t *testing.T) {
	s := &Server{}
	r := chi.NewRouter()
	r.Get("/", s.helloWorldHandler)

	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	defer res.Body.Close()
	var resp map[string]string
	err = json.NewDecoder(res.Body).Decode(&resp)
	assert.NoError(t, err)
	assert.Equal(t, "Hello World", resp["message"])
}
`)
}