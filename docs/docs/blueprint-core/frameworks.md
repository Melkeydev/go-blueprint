## Frameworks


Created project can utilizes several Go web frameworks to handle HTTP routing and server functionality. The chosen frameworks are:

1. [**Chi**](https://github.com/go-chi/chi): Lightweight and flexible router for building Go HTTP services.
2. [**Gin**](https://github.com/gin-gonic/gin): A web framework with a martini-like API, but with much better performance.
3. [**Fiber**](https://github.com/gofiber/fiber): Express-inspired web framework designed to be fast, simple, and efficient.
4. [**HttpRouter**](https://github.com/julienschmidt/httprouter): A high-performance HTTP request router that scales well.
5. [**Gorilla/mux**](https://github.com/gorilla/mux): A powerful URL router and dispatcher for golang.
6. [**Echo**](https://github.com/labstack/echo): High-performance, extensible, minimalist Go web framework.

## Project Structure

The project is structured with a simple layout, focusing on the cmd, internal, and tests directories:

```bash
/(Root)
├── /cmd
│   └── /api
│       └── main.go
├── /internal
│   └── /server
│       ├── routes.go
│       └── server.go
├── /tests
│   └── handler_test.go
├── go.mod
├── go.sum
├── Makefile
└── README.md
```