package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mr-utzig/gear-stock/internal/api/routers"
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

func (s *Server) Start() {
	e := echo.New()

	e.Use(
		middleware.Secure(),
		middleware.Recover(),
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "method=${method}, path=${path}, status=${status} err=${error}\n",
		}),
	)

	// Auth
	routers.Auth(e)
	// Users
	routers.User(e)
	// Customers
	// Orders
	// Gears

	e.Logger.Fatal(e.Start(s.port))
}
