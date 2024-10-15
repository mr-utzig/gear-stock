package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func MaintenanceListPage(c echo.Context) error {

	return c.Render(http.StatusOK, "maintenance/index", nil)
}
