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

func (s StandardLibTemplate) RoutesWithDB() []byte {
	return MakeHTTPRoutesWithDB()
}

// MakeHTTPServer returns a byte slice that represents
// the default internal/server/server.go file.
func MakeHTTPServer() []byte {
	return []byte(`package server

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"
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

func MakeHTTPRoutesWithDB() []byte {
	return []byte(`package server

import (
	"net/http"
	"encoding/json"
	"log"
	"{{.ProjectName}}/internal/database"
)

type healthHandler struct {
	s database.Service
}

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	h := NewHealthHandler()

	mux.HandleFunc("/", s.handler)
	mux.HandleFunc("/health", h.healthHandler)

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
