Created projects can utilize several Go web frameworks to handle HTTP routing and server functionality. The chosen frameworks are:

1. [**Chi**](https://github.com/go-chi/chi): Lightweight and flexible router for building Go HTTP services.
2. [**Echo**](https://github.com/labstack/echo): High-performance, extensible, minimalist Go web framework.
3. [**Fiber**](https://github.com/gofiber/fiber): Express-inspired web framework designed to be fast, simple, and efficient.
4. [**Gin**](https://github.com/gin-gonic/gin): A web framework with a martini-like API, but with much better performance.
5. [**Gorilla/mux**](https://github.com/gorilla/mux): A powerful URL router and dispatcher for Golang.
6. [**HttpRouter**](https://github.com/julienschmidt/httprouter): A high-performance HTTP request router that scales well.
7. [**Huma**](https://huma.rocks/): A modern framework for building REST/GraphQL APIs in Go, designed for speed and developer experience. It can be used with various routers, including Fiber.

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
