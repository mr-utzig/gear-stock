package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

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
func (g *GearHandler) GetAllGears(c echo.Context) error {
	gears, err := g.model.GetAllGears()
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, utils.NewResponse(false, "Not Found", gears))
		}
		c.Logger().Error("g.model.GetAllGears()", err)

		return c.JSON(http.StatusInternalServerError, utils.NewResponse(false, "Internal Server Error", gears))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(true, "OK", gears))
}

// Create Gear:
func (g *GearHandler) CreateGear(c echo.Context) error {
	data := new(models.GearCreateRequest)
	if err := c.Bind(data); err != nil {
		c.Logger().Error("c.Bind(data)", &data, err)

		return c.JSON(http.StatusBadRequest, utils.NewResponse(false, "Bad Request", data))
	}

	gear, err := g.model.CreateGear(data)
	if err != nil {
		c.Logger().Error("g.model.CreateGear(data)", &data, err)

		return c.JSON(http.StatusInternalServerError, utils.NewResponse(false, "Internal Server Error", data))
	}

	return c.JSON(http.StatusCreated, utils.NewResponse(true, "Created", gear))
}

// Get Gear:
func (g *GearHandler) GetGear(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	gear, err := g.model.GetGear(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, utils.NewResponse(false, "Not Found", gear))
		}
		c.Logger().Error("g.model.GetGear(id)", id, c.Param("id"), err)

		return c.JSON(http.StatusInternalServerError, utils.NewResponse(false, "Internal Server Error", gear))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(true, "OK", gear))
}

// Update Gear:
func (g *GearHandler) UpdateGear(c echo.Context) error {
	data := new(models.GearUpdateRequest)
	if err := c.Bind(data); err != nil {
		c.Logger().Error("c.Bind(data)", &data, err)

		return c.JSON(http.StatusBadRequest, utils.NewResponse(false, "Bad Request", data))
	}

	id, _ := strconv.Atoi(c.Param("id"))
	gear, err := g.model.UpdateGear(id, data)
	if err != nil {
		c.Logger().Error("g.model.UpdateGear(data)", &data, err)

		return c.JSON(http.StatusInternalServerError, utils.NewResponse(false, "Internal Server Error", data))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(true, "OK", gear))
}

// Delete Gear:
func (g *GearHandler) DeleteGear(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	gear, err := g.model.DeleteGear(id)
	if err != nil {
		c.Logger().Error("g.model.GetGear(id)", id, c.Param("id"), err)

		return c.JSON(http.StatusInternalServerError, utils.NewResponse(false, "Internal Server Error", gear))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(true, "OK", gear))
}
