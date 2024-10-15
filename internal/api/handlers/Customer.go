package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mr-utzig/gear-stock/internal/api/models"
	"github.com/mr-utzig/gear-stock/internal/api/utils"
)

type CustomerHandler struct {
	model *models.CustomerModel
}

func NewCustomerHandler(model *models.CustomerModel) *CustomerHandler {
	return &CustomerHandler{model: model}
}

// Get all Customers:
func (o *CustomerHandler) GetAllCustomers(c echo.Context) error {
	response := utils.NewResponse(true, "OK", o.model.GetAllCustomers())
	return c.JSON(http.StatusOK, response)
}

// Create Customer:
func (o *CustomerHandler) CreateCustomer(c echo.Context) error {
	response := utils.NewResponse(true, "OK", o.model.CreateCustomer())
	return c.JSON(http.StatusOK, response)
}

// Get Customer:
func (o *CustomerHandler) GetCustomer(c echo.Context) error {
	response := utils.NewResponse(true, "OK", o.model.GetCustomer(1))
	return c.JSON(http.StatusOK, response)
}

// Update Customer:
func (o *CustomerHandler) UpdateCustomer(c echo.Context) error {
	response := utils.NewResponse(true, "OK", o.model.UpdateCustomer(1))
	return c.JSON(http.StatusOK, response)
}
