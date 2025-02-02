Created project can utilizes several Go web backends to handle HTTP routing and server functionality. The chosen backends are:

1. [**Standard-library**](https://pkg.go.dev/std): Offers a vast collection of packages and functions.
2. [**Chi**](https://github.com/go-chi/chi): Lightweight and flexible router for building Go HTTP services.
3. [**Gin**](https://github.com/gin-gonic/gin): A web framework with a martini-like API, but with much better performance.
4. [**Fiber**](https://github.com/gofiber/fiber): Express-inspired web framework designed to be fast, simple, and efficient.
5. [**Gorilla/mux**](https://github.com/gorilla/mux): A powerful URL router and dispatcher for Golang.
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
│       ├── routes_test.go
│       └── server.go
├── go.mod
├── go.sum
├── Makefile
└── README.md
```
