package app

import (
	"embed"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mr-utzig/gear-stock/app/routers"
	"github.com/mr-utzig/gear-stock/app/utils"
)

//go:embed web/public web/views
var WebFS embed.FS

func Setup() {
	e := echo.New()
	e.Use(
		middleware.Secure(),
		middleware.Recover(),
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "method=${method}, path=${path}, status=${status} err=${error}\n",
		}),
	)

	e.Renderer = &utils.Template{
		FS: WebFS,
	}

	routers.Setup(e)

	e.Logger.Fatal(e.Start(":1334"))
}
