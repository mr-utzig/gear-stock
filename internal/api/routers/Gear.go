package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/mr-utzig/gear-stock/internal/api/handlers"
)

func Gear(e *echo.Echo, handler *handlers.GearHandler) {
	gears := e.Group("/gears")

	// get all gears
	gears.GET("", handler.GetAllGears)
	// create gear
	gears.POST("", handler.CreateGear)
	// get gear
	gears.GET(":id", handler.GetGear)
	// update gear
	gears.PUT(":id", handler.UpdateGear)
	// delete gear
	gears.DELETE(":id", handler.DeleteGear)
}
