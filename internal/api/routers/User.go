package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/mr-utzig/gear-stock/internal/api/handlers"
)

func User(e *echo.Echo, handler *handlers.UserHandler) {
	users := e.Group("/users")

	// Get all Users:
	users.GET("", handler.GetAllUsers)
	// Create User:
	users.POST("", handler.CreateUser)
	// Get User:
	users.GET(":id", handler.GetUser)
	// Update User:
	users.PUT(":id", handler.UpdateUser)
}
