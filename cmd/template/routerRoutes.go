package template

import "fmt"

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

// MakeRouterRoutes returns a byte slice that represents 
// the internal/server/routes.go file when using HttpRouter
func MakeRouterRoutes() []byte {
	return []byte(fmt.Sprintf(`package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"nhooyr.io/websocket"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := httprouter.New()
	r.HandlerFunc(http.MethodGet, "/", s.helloWorldHandler)
	r.HandlerFunc(http.MethodGet, "/ws", s.pingPongWebsocketHandler)

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
%s
}
`, websocketTemplate()))
}
