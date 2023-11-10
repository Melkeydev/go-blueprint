package template

import (
	_ "embed"
)

//go:embed files/routes/standard_library.go.tmpl
var standardRoutesTemplate []byte

//go:embed files/server/standard_library.go.tmpl
var standardServerTemplate []byte

// StandardLibTemplate contains the methods used for building
// an app that uses [net/http]
type StandardLibTemplate struct{}

func (s StandardLibTemplate) Main() []byte {
	return mainTemplate
}

func (s StandardLibTemplate) Server() []byte {
	return standardServerTemplate
}

func (s StandardLibTemplate) Routes() []byte {
	return standardRoutesTemplate
}

// MakeHTTPServer returns a byte slice that represents 
// the default internal/server/server.go file.
func MakeHTTPServer() []byte {
	return []byte(`package server

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
}

func NewServer() *http.Server {

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		fmt.Printf("error converting variable to int, defaulting to 8080.")
		port = 8080
	}

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
