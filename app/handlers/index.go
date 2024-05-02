package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}

func Login(c echo.Context) error {
	return c.Render(http.StatusOK, "login", nil)
}
