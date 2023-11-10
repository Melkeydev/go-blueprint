package template

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

func (s StandardLibTemplate) Config() []byte {
	return ConfigTemplate()
}

// MakeHTTPServer returns a byte slice that represents
// the default internal/server/server.go file.
func MakeHTTPServer() []byte {
	return []byte(`package server

import (
	"fmt"
	"net/http"
	"time"
	"{{.ProjectName}}/internal/config"
)


type Server struct {
	port int
	env string
}

func NewServer(config config.Config) *http.Server {
	NewServer := &Server{
		port: config.Port,
		env: config.Environment,
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
