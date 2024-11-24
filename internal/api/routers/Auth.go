package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/mr-utzig/gear-stock/internal/api/handlers"
	"github.com/mr-utzig/gear-stock/internal/api/models"
)

func Auth(e *echo.Echo) {
	userModel := models.NewUserModel()
	handler := handlers.NewAuthHandler(userModel)

	auth := e.Group("/auth")

	// Login:
	auth.POST("/login", handler.HandleLogin)
	// Logout:
	// auth.POST("", handler.HandleLogout)
}
