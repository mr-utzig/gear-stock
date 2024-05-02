package app

import (
	"embed"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mr-utzig/gear-stock/app/routers"
	"github.com/mr-utzig/gear-stock/app/utils"
)

//go:embed web/public web/views
var WebFS embed.FS

func Setup() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = &utils.Template{
		FS: WebFS,
	}

	routers.Setup(e)

	e.Logger.Fatal(e.Start(":1334"))
}
