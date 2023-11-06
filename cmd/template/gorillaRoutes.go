package template

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

// MakeGorillaRoutes returns a byte slice that represents 
// the internal/server/routes.go file when using gorilla/mux.
func MakeGorillaRoutes() []byte {
	return []byte(`package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"nhooyr.io/websocket"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", s.helloWorldHandler)
	r.HandleFunc("/ws", s.pingPongWebsocketHandler)

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

func (s *Server) pingPongWebsocketHandler(w http.ResponseWriter, r *http.Request) {
	socket, _ := websocket.Accept(w, r, nil)
	ctx := r.Context()
	for {
		_, socketBytes, _ := socket.Read(ctx)
		if string(socketBytes) == "PING" {
			_ = socket.Write(ctx, websocket.MessageText, []byte("PONG"))
		} else {
			_ = socket.Write(ctx, websocket.MessageText, []byte("HUH?"))
		}
	}
}
`)
}
