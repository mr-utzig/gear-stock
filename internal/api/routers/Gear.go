package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/mr-utzig/gear-stock/internal/api/handlers"
	"github.com/mr-utzig/gear-stock/internal/api/models"
)

func Gear(e *echo.Echo) {
	model := models.NewGearModel()
	handler := handlers.NewGearHandler(model)

	gears := e.Group("/gears")

	// get all gears
	gears.GET("", handler.GetAllGears)
	// create gear
	gears.POST("", handler.CreateGear)
	// get gear
	gears.GET("/:id", handler.GetGear)
	// update gear
	gears.PUT("/:id", handler.UpdateGear)
	// delete gear
	gears.DELETE("/:id", handler.DeleteGear)
	// get gears by order id
	gears.GET("/order/:id", handler.GetGearsByOrderID)
}
