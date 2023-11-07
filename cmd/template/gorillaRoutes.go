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
	socket, err := websocket.Accept(w, r, nil)

	if err != nil {
		log.Print("could not open websocket")
		w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx := c.Request.Context()
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
