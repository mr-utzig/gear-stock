package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/mr-utzig/gear-stock/internal/api/handlers"
)

func Order(e *echo.Echo, handler *handlers.OrderHandler) {
	orders := e.Group("/orders")

	// Get all Orders:
	orders.GET("", handler.GetAllOrders)
	// Create Order:
	orders.POST("", handler.CreateOrder)
	// Get Order:
	orders.GET(":id", handler.GetOrder)
	// Update Order:
	orders.PUT(":id", handler.UpdateOrder)
}
