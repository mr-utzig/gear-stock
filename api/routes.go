package api

import (
	"github.com/labstack/echo"
)

func SetupRoutes(e *echo.Echo) {
	e.Static("/", "api/web/public")

	e.GET("/", IndexHandler)
	e.GET("/login", LoginHandler)
}
