package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

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
	orders, err := o.model.GetAllOrders()
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, utils.NewResponse(false, "Not Found", orders))
		}
		c.Logger().Error("o.model.GetAllOrders()", err)

		return c.JSON(http.StatusInternalServerError, utils.NewResponse(false, "Internal Server Error", orders))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(true, "OK", orders))
}

// Create Order:
func (o *OrderHandler) CreateOrder(c echo.Context) error {
	data := new(models.OrderCreateRequest)
	if err := c.Bind(data); err != nil {
		c.Logger().Error("c.Bind(data)", &data, err)

		return c.JSON(http.StatusBadRequest, utils.NewResponse(false, "Bad Request", data))
	}

	order, err := o.model.CreateOrder(data)
	if err != nil {
		c.Logger().Error("o.model.CreateOrder(data)", &data, err)

		return c.JSON(http.StatusInternalServerError, utils.NewResponse(false, "Internal Server Error", data))
	}

	return c.JSON(http.StatusCreated, utils.NewResponse(true, "Created", order))
}

// Get Order:
func (o *OrderHandler) GetOrder(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := o.model.GetOrder(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, utils.NewResponse(false, "Not Found", order))
		}
		c.Logger().Error("o.model.GetOrder(id)", id, c.Param("id"), err)

		return c.JSON(http.StatusInternalServerError, utils.NewResponse(false, "Internal Server Error", order))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(true, "OK", order))
}

// Update Order:
func (o *OrderHandler) UpdateOrder(c echo.Context) error {
	data := new(models.OrderUpdateRequest)
	if err := c.Bind(data); err != nil {
		c.Logger().Error("c.Bind(data)", &data, err)

		return c.JSON(http.StatusBadRequest, utils.NewResponse(false, "Bad Request", data))
	}

	id, _ := strconv.Atoi(c.Param("id"))
	order, err := o.model.UpdateOrder(id, data)
	if err != nil {
		c.Logger().Error("o.model.UpdateOrder(data)", &data, err)

		return c.JSON(http.StatusInternalServerError, utils.NewResponse(false, "Internal Server Error", data))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(true, "OK", order))
}
