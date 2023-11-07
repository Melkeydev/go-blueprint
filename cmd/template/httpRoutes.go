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

	"nhooyr.io/websocket"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handler)
	mux.HandleFunc("/ws", s.pingPongWebsocketHandler)

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

func (s *Server) pingPongWebsocketHandler(w http.ResponseWriter, r *http.Request) {
	socket, err := websocket.Accept(w, r, nil)

	if err != nil {
		log.Print("could not open websocket")
		// pray this works for your user
		_, _ = w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx := r.Context()
	for {
		msgType, socketBytes, err := socket.Read(ctx)

		if err != nil {
			log.Print("could not read from websocket")
			return
		}

		if string(socketBytes) == "PING" {
			if err := socket.Write(ctx, msgType, []byte("PONG")); err != nil {
				log.Print("could not write to socket")
				return
			}
		} else {
			if err := socket.Write(ctx, msgType, []byte("HUH?")); err != nil {
				log.Print("could not write to socket")
				return
			}
		}
	}
}
`)
}
