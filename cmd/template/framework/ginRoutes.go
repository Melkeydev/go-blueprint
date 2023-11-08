package framework

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

func (g GinTemplates) RoutesWithDB() []byte {
	return MakeGinRoutesWithDB()
}

// MakeGinRoutes returns a byte slice that represents
// the internal/server/routes.go file when using Gin.
func MakeGinRoutes() []byte {
	return []byte(`package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", s.helloWorldHandler)

	return r
}

func (s *Server) helloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

`)
}

func MakeGinRoutesWithDB() []byte {
	return []byte(`package server

import (
	"net/http"
	"{{.ProjectName}}/internal/database"

	"github.com/gin-gonic/gin"
)

type healthHandler struct {
	s database.Service
}

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	h := NewHealthHandler()
	r.GET("/", s.helloWorldHandler)
	r.GET("/health", h.healthHandler)

	return r
}

func (s *Server) helloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func NewHealthHandler() *healthHandler {
	return &healthHandler{
		s: database.New(),
	}
}

func (h *healthHandler) healthHandler(c *gin.Context) error {
	c.JSON(http.StatusOK, h.s.Health())
}
`)
}
