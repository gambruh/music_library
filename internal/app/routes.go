package app

import "github.com/labstack/echo/v4"

// adding routes
func (s *Service) addRoutes(e *echo.Echo) {
	e.GET("/api", s.handleAPI)
	e.POST("/api/getsong", s.handleGetSong)
	e.POST("/api/addsong", s.handleAddSong)
	e.POST("/api/editsong", s.handleEditSong)
	e.POST("/api/delsong", s.handleDeleteSong)
	e.GET("/info", s.handleGetSongDetails)
}
