package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mr-utzig/gear-stock/internal/api/models"
	"github.com/mr-utzig/gear-stock/internal/api/utils"
)

type GearHandler struct {
	model *models.GearModel
}

func NewGearHandler(model *models.GearModel) *GearHandler {
	return &GearHandler{model: model}
}

// Get all Gears:
func (o *GearHandler) GetAllGears(c echo.Context) error {
	response := utils.NewResponse(true, "OK", o.model.GetAllGears())
	return c.JSON(http.StatusOK, response)
}

// Create Gear:
func (o *GearHandler) CreateGear(c echo.Context) error {
	response := utils.NewResponse(true, "OK", o.model.CreateGear())
	return c.JSON(http.StatusOK, response)
}

// Get Gear:
func (o *GearHandler) GetGear(c echo.Context) error {
	response := utils.NewResponse(true, "OK", o.model.GetGear(1))
	return c.JSON(http.StatusOK, response)
}

// Update Gear:
func (o *GearHandler) UpdateGear(c echo.Context) error {
	response := utils.NewResponse(true, "OK", o.model.UpdateGear(1))
	return c.JSON(http.StatusOK, response)
}

// Delete Gear:
func (o *GearHandler) DeleteGear(c echo.Context) error {
	response := utils.NewResponse(true, "OK", o.model.DeleteGear(1))
	return c.JSON(http.StatusOK, response)
}
