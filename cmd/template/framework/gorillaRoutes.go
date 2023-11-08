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

func (g GorillaTemplates) RoutesWithDB() []byte {
	return MakeGorillaRoutesWithDB()
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

func MakeGorillaRoutesWithDB() []byte {
	return []byte(`package server

import (
	"encoding/json"
	"log"
	"net/http"
	"{{.ProjectName}}/internal/database"

	"github.com/gorilla/mux"
)

type healthHandler struct {
	s database.Service
}

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()
	h := NewHealthHandler()

	r.HandleFunc("/", s.helloWorldHandler)
	r.HandleFunc("/health", h.healthHandler)

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
