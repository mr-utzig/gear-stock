package routers

import (
	"encoding/base64"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/mr-utzig/gear-stock/app/handlers"
)

func Setup(e *echo.Echo) {
	e.Static("/public", "app/web/public")

	login := e.Group("login")
	login.GET("", handlers.LoginPage)
	login.POST("/validate", handlers.LoginValidation)

	SigningKey, _ := base64.StdEncoding.DecodeString(os.Getenv("JWT_SIGING_KEY"))
	logged := e.Group("")
	logged.Use(echojwt.WithConfig(echojwt.Config{
		SigningMethod: "HS512",
		SigningKey:    SigningKey,
		TokenLookup:   "cookie:tkn",
		ErrorHandler: func(c echo.Context, err error) error {
			return c.Redirect(http.StatusPermanentRedirect, "/login")
		},
	}))
	logged.GET("", handlers.MaintenanceListPage)
}
