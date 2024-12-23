package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

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
func (ct *CustomerHandler) GetAllCustomers(c echo.Context) error {
	customers, err := ct.model.GetAllCustomers()
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, utils.NewResponse(false, "Not Found", customers))
		}
		c.Logger().Error("ct.model.GetAllCustomers()", err)

		return c.JSON(http.StatusInternalServerError, utils.NewResponse(false, "Internal Server Error", customers))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(true, "OK", customers))
}

// Create Customer:
func (ct *CustomerHandler) CreateCustomer(c echo.Context) error {
	data := new(models.CustomerCreateRequest)
	if err := c.Bind(data); err != nil {
		c.Logger().Error("c.Bind(data)", &data, err)

		return c.JSON(http.StatusBadRequest, utils.NewResponse(false, "Bad Request", data))
	}

	customer, err := ct.model.CreateCustomer(data)
	if err != nil {
		c.Logger().Error("ct.model.CreateCustomer(data)", &data, err)

		return c.JSON(http.StatusInternalServerError, utils.NewResponse(false, "Internal Server Error", data))
	}

	return c.JSON(http.StatusCreated, utils.NewResponse(true, "Created", customer))
}

// Get Customer:
func (ct *CustomerHandler) GetCustomer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	customer, err := ct.model.GetCustomer(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, utils.NewResponse(false, "Not Found", customer))
		}
		c.Logger().Error("ct.model.GetCustomer(id)", id, c.Param("id"), err)

		return c.JSON(http.StatusInternalServerError, utils.NewResponse(false, "Internal Server Error", customer))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(true, "OK", customer))
}

// Update Customer:
func (ct *CustomerHandler) UpdateCustomer(c echo.Context) error {
	data := new(models.CustomerUpdateRequest)
	if err := c.Bind(data); err != nil {
		c.Logger().Error("c.Bind(data)", &data, err)

		return c.JSON(http.StatusBadRequest, utils.NewResponse(false, "Bad Request", data))
	}

	id, _ := strconv.Atoi(c.Param("id"))
	customer, err := ct.model.UpdateCustomer(id, data)
	if err != nil {
		c.Logger().Error("ct.model.UpdateCustomer(data)", &data, err)

		return c.JSON(http.StatusInternalServerError, utils.NewResponse(false, "Internal Server Error", data))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(true, "OK", customer))
}
