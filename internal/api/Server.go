package api

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mr-utzig/gear-stock/internal/api/handlers"
	"github.com/mr-utzig/gear-stock/internal/api/models"
	"github.com/mr-utzig/gear-stock/internal/api/routers"
)

type Server struct {
	port string
	db   *sql.DB
}

func NewServer(port string, db *sql.DB) *Server {
	return &Server{
		port: port,
		db:   db,
	}
}

func (s *Server) Start() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Users
	// Customers
	// Orders
	orderModel := models.NewOrderModel(s.db)
	orderHandler := handlers.NewOrderHandler(orderModel)
	routers.Order(e, orderHandler)
	// Gears
	gearModel := models.NewGearModel(s.db)
	gearHandler := handlers.NewGearHandler(gearModel)
	routers.Gear(e, gearHandler)

	e.Logger.Fatal(e.Start(s.port))
}
