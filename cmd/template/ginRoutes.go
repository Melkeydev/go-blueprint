package template

// GinTemplates contains the methods used for building
// an app that uses [github.com/gin-gonic/gin]
type GinTemplates struct{}

func (g GinTemplates) Main() []byte {
	return MainTemplate()
}

func (g GinTemplates) Server() []byte {
	return MakeHTTPServer()
}

func (g GinTemplates) Routes() []byte {
	return MakeGinRoutes()
}

// MakeGinRoutes returns a byte slice that represents 
// the internal/server/routes.go file when using Gin.
func MakeGinRoutes() []byte {
	return []byte(`package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", s.helloWorldHandler)
	r.GET("/ws", s.pingPongWebsocketHandler)

	return r
}

func (s *Server) helloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) pingPongWebsocketHandler(c *gin.Context) {
	socket, err := websocket.Accept(c.Writer, c.Request, nil)

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
