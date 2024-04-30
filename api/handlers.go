package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func IndexHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func LoginHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}
