package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handlers
func handleAPI(c echo.Context) error {
	return c.String(http.StatusOK, "API endpoint reached")
}
