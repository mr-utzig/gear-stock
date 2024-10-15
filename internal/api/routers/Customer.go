package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/mr-utzig/gear-stock/internal/api/handlers"
)

func Customer(e *echo.Echo, handler *handlers.CustomerHandler) {
	customers := e.Group("/customers")

	// Get all Customers:
	customers.GET("", handler.GetAllCustomers)
	// Create Customer:
	customers.POST("", handler.CreateCustomer)
	// Get Customer:
	customers.GET(":id", handler.GetCustomer)
	// Update Customer:
	customers.PUT(":id", handler.UpdateCustomer)
}
