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
