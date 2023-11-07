package template

// FiberTemplates contains the methods used for building
// an app that uses [github.com/gofiber/fiber]
type FiberTemplates struct{}

func (f FiberTemplates) Main() []byte {
	return MakeFiberMain()
}
func (f FiberTemplates) Server() []byte {
	return MakeFiberServer()
}

func (f FiberTemplates) Routes() []byte {
	return MakeFiberRoutes()
}

// MakeFiberServer returns a byte slice that represents 
// the internal/server/server.go file when using Fiber.
func MakeFiberServer() []byte {
	return []byte(`package server

import "github.com/gofiber/fiber/v2"

type FiberServer struct {
	*fiber.App
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(),
	}

	return server
}

`)
}

// MakeFiberRoutes returns a byte slice that represents 
// the internal/server/routes.go file when using Fiber.
func MakeFiberRoutes() []byte {
	return []byte(`package server

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.helloWorldHandler)
	s.App.Get("/ws", websocket.New(s.pingPongWebsocketHandler))
}

func (s *FiberServer) helloWorldHandler(c *fiber.Ctx) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(resp)
}

func (s *FiberServer) pingPongWebsocketHandler(con *websocket.Conn) {
	for {
		messageType, socketBytes, err := con.ReadMessage()

		if err != nil {
			log.Print("could not read from websocket")
			return
		}

		if string(socketBytes) == "PING" {
			if err := con.WriteMessage(messageType, []byte("PONG")); err != nil {
				log.Print("could not write to socket")
				return
			}
		} else {
			if err  := con.WriteMessage(messageType, []byte("HUH?")); err != nil {
				log.Print("could not write to socket")
				return
			}
		}
	}
}
`)
}

// MakeHTTPRoutes returns a byte slice that represents 
// the cmd/api/main.go file when using Fiber.
func MakeFiberMain() []byte {
	return []byte(`package main

import (
	"{{.ProjectName}}/internal/server"
)

func main() {

	server := server.New()

	server.RegisterFiberRoutes()

	err := server.Listen(":8080")
	if err != nil {
		panic("cannot start server")
	}
}
`)
}
