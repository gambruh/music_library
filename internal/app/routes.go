package app

import "github.com/labstack/echo/v4"

// adding routes
func (s *Service) addRoutes(e *echo.Echo) {
	e.GET("/api", s.handleAPI)
	e.POST("/api/getsong", s.HandleGetSong)
	e.POST("/api/addsong", s.HandleAddSong)
	e.POST("/api/editsong", s.HandleEditSong)
	e.POST("/api/delsong", s.HandleDeleteSong)
	e.GET("/info", s.handleGetSongDetails)
}
