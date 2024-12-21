package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

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
func (u *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := u.model.GetAllUsers()
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, utils.NewResponse(false, "Not Found", users))
		}
		c.Logger().Error("u.model.GetAllUsers()", err)

		return c.JSON(http.StatusInternalServerError, utils.NewResponse(false, "Internal Server Error", users))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(true, "OK", users))
}

// Create User:
func (u *UserHandler) CreateUser(c echo.Context) error {
	data := new(models.UserCreateRequest)
	if err := c.Bind(data); err != nil {
		c.Logger().Error("c.Bind(data)", &data, err)

		return c.JSON(http.StatusBadRequest, utils.NewResponse(false, "Bad Request", data))
	}

	user, err := u.model.CreateUser(data)
	if err != nil {
		c.Logger().Error("u.model.CreateUser(data)", &data, err)

		return c.JSON(http.StatusInternalServerError, utils.NewResponse(false, "Internal Server Error", data))
	}

	return c.JSON(http.StatusCreated, utils.NewResponse(true, "Created", user))
}

// Get User:
func (u *UserHandler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := u.model.GetUser(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, utils.NewResponse(false, "Not Found", user))
		}
		c.Logger().Error("u.model.GetUser(id)", id, c.Param("id"), err)

		return c.JSON(http.StatusInternalServerError, utils.NewResponse(false, "Internal Server Error", user))
	}

	userData := models.UserDataResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Status:    user.Status,
		ProfileID: user.ProfileID,
	}

	return c.JSON(http.StatusOK, utils.NewResponse(true, "OK", userData))
}

// Update User:
func (u *UserHandler) UpdateUser(c echo.Context) error {
	data := new(models.UserUpdateRequest)
	if err := c.Bind(data); err != nil {
		c.Logger().Error("c.Bind(data)", &data, err)

		return c.JSON(http.StatusBadRequest, utils.NewResponse(false, "Bad Request", data))
	}

	id, _ := strconv.Atoi(c.Param("id"))
	user, err := u.model.UpdateUser(id, data)
	if err != nil {
		c.Logger().Error("u.model.UpdateUser(data)", &data, err)

		return c.JSON(http.StatusInternalServerError, utils.NewResponse(false, "Internal Server Error", data))
	}

	userData := models.UserDataResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Status:    user.Status,
		ProfileID: user.ProfileID,
	}

	return c.JSON(http.StatusOK, utils.NewResponse(true, "OK", userData))
}
