package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mr-utzig/gear-stock/internal/api/models"
	"github.com/mr-utzig/gear-stock/internal/api/utils"
)

type OrderHandler struct {
	model *models.OrderModel
}

func NewOrderHandler(model *models.OrderModel) *OrderHandler {
	return &OrderHandler{model: model}
}

// Get all Orders:
func (o *OrderHandler) GetAllOrders(c echo.Context) error {
	response := utils.NewResponse(true, "OK", o.model.GetAllOrders())
	return c.JSON(http.StatusOK, response)
}

// Create Order:
func (o *OrderHandler) CreateOrder(c echo.Context) error {
	response := utils.NewResponse(true, "OK", o.model.CreateOrder())
	return c.JSON(http.StatusOK, response)
}

// Get Order:
func (o *OrderHandler) GetOrder(c echo.Context) error {
	response := utils.NewResponse(true, "OK", o.model.GetOrder(1))
	return c.JSON(http.StatusOK, response)
}

// Update Order:
func (o *OrderHandler) UpdateOrder(c echo.Context) error {
	response := utils.NewResponse(true, "OK", o.model.UpdateOrder(1))
	return c.JSON(http.StatusOK, response)
}
