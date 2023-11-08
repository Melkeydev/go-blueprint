package framework

// EchoTemplates contains the methods used for building
// an app that uses [github.com/labstack/echo]
type EchoTemplates struct{}

func (e EchoTemplates) Main() []byte {
	return MainTemplate()
}
func (e EchoTemplates) Server() []byte {
	return MakeHTTPServer()
}

func (e EchoTemplates) Routes() []byte {
	return MakeEchoRoutes()
}

// MakeEchoRoutes returns a byte slice that represents
// the internal/server/routes.go file when using Echo.
func (e EchoTemplates) RoutesWithDB() []byte {
	return MakeEchoRoutesWithDB()
}

func MakeEchoRoutes() []byte {
	return []byte(`package server

	import (
		"net/http"
	
		"github.com/labstack/echo/v4"
		"github.com/labstack/echo/v4/middleware"
	)
func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.helloWorldHandler)

	return e
}

func (s *Server) helloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}
`)
}

func MakeEchoRoutesWithDB() []byte {
	return []byte(`package server

	import (
		"net/http"
		"{{.ProjectName}}/internal/database"

		"github.com/labstack/echo/v4"
		"github.com/labstack/echo/v4/middleware"
	)

type healthHandler struct {
	s database.Service
}

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	d := NewHealthHandler()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.helloWorldHandler)
	e.GET("/health", d.healthHandler)

	return e
}

func (s *Server) helloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func NewHealthHandler() *healthHandler {
	return &healthHandler{
		s: database.New(),
	}
}

func (h *healthHandler) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, h.s.Health())
}

`)
}
