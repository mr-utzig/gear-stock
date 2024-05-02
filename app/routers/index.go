package routers

import (
	"github.com/labstack/echo"
	"github.com/mr-utzig/gear-stock/app/handlers"
)

func Setup(e *echo.Echo) {
	e.Static("/", "/app/web/public")

	e.GET("/", handlers.Index)
	e.GET("/login", handlers.Login)
}
