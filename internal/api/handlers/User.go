package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mr-utzig/gear-stock/internal/api/models"
	"github.com/mr-utzig/gear-stock/internal/api/utils"
)

type UserHandler struct {
	model *models.UserModel
}

func NewUserHandler(model *models.UserModel) *UserHandler {
	return &UserHandler{model: model}
}

// Get all Users:
func (o *UserHandler) GetAllUsers(c echo.Context) error {
	response := utils.NewResponse(true, "OK", o.model.GetAllUsers())
	return c.JSON(http.StatusOK, response)
}

// Create User:
func (o *UserHandler) CreateUser(c echo.Context) error {
	data := new(models.PostUser)
	if err := c.Bind(data); err != nil {
		c.Logger().Error("c.Bind(data)", err)

		return c.JSON(http.StatusBadRequest, utils.NewResponse(false, "Bad Request", data))
	}

	user, err := o.model.CreateUser(data)
	if err != nil {
		c.Logger().Error("o.model.CreateUser", err)

		return c.JSON(http.StatusInternalServerError, utils.NewResponse(false, "Internal Server Error", data))
	}

	return c.JSON(http.StatusCreated, utils.NewResponse(true, "Created", user))
}

// Get User:
func (o *UserHandler) GetUser(c echo.Context) error {
	response := utils.NewResponse(true, "OK", o.model.GetUser(1))
	return c.JSON(http.StatusOK, response)
}

// Update User:
func (o *UserHandler) UpdateUser(c echo.Context) error {
	response := utils.NewResponse(true, "OK", o.model.UpdateUser(1))
	return c.JSON(http.StatusOK, response)
}
