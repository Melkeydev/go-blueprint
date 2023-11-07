package template

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

// MakeChiRoutes returns a byte slice that represents 
// the internal/server/routes.go file when using Chi.
func MakeChiRoutes() []byte {
	return []byte(`package server

import ( "encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"nhooyr.io/websocket"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.helloWorldHandler)
	r.Get("/ws", s.pingPongWebsocketHandler)

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
