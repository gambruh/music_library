package app

import "github.com/labstack/echo/v4"

// adding routes
func addRoutes(e *echo.Echo) {
	e.GET("/", handleAPI)
}
