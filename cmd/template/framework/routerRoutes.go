package framework

// RouterTemplates contains the methods used for building
// an app that uses [github.com/julienschmidt/httprouter]
type RouterTemplates struct{}

func (r RouterTemplates) Main() []byte {
	return MainTemplate()
}
func (r RouterTemplates) Server() []byte {
	return MakeHTTPServer()
}
func (r RouterTemplates) Routes() []byte {
	return MakeRouterRoutes()
}

func (r RouterTemplates) RoutesWithDB() []byte {
	return MakeRouterRoutesWithDB()
}

// MakeRouterRoutes returns a byte slice that represents
// the internal/server/routes.go file when using HttpRouter
func MakeRouterRoutes() []byte {
	return []byte(`package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := httprouter.New()
	r.HandlerFunc(http.MethodGet, "/", s.helloWorldHandler)

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

func MakeRouterRoutesWithDB() []byte {
	return []byte(`package server

import (
	"encoding/json"
	"log"
	"net/http"
	"{{.ProjectName}}/internal/database"

	"github.com/julienschmidt/httprouter"
)

type healthHandler struct {
	s database.Service
}

func (s *Server) RegisterRoutes() http.Handler {
	r := httprouter.New()
	h := NewHealthHandler()

	r.HandlerFunc(http.MethodGet, "/", s.helloWorldHandler)
	r.HandlerFunc(http.MethodGet, "/health", h.healthHandler)

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

func NewHealthHandler() *healthHandler {
	return &healthHandler{
		s: database.New(),
	}
}

func (h *healthHandler) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(h.s.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Write(jsonResp)
}

`)
}
