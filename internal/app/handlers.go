package app

import (
	"fmt"
	"net/http"

	"github.com/gambruh/music_library/internal/storage"
	"github.com/labstack/echo/v4"
)

// Handlers
func (s *Service) handleAPI(c echo.Context) error {
	return c.String(http.StatusOK, "API endpoint reached")
}

// HandleGetSong retrieves a song based on the provided query parameters
func (s *Service) handleGetSong(c echo.Context) error {
	type reqbody struct {
		Name  string `json:"song"`
		Group string `json:"group"`
	}

	var rsong reqbody
	ctx := c.Request().Context()

	if err := c.Bind(&rsong); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "song and group parameters are required"})
	}

	fmt.Println(rsong)

	song, err := s.Storage.GetSong(ctx, rsong.Group, rsong.Name)
	if err != nil {
		fmt.Println("failed to get song: ", err)
		return c.JSON(http.StatusNotFound, echo.Map{"error": "song not found"})
	}

	s.Logger.Infof("get song: %s,%s", song.Group, song.Name)
	return c.JSON(http.StatusOK, song)
}

// HandleAddSong adds a new song to the database
func (s *Service) handleAddSong(c echo.Context) error {
	var song storage.Song

	ctx := c.Request().Context()

	if err := c.Bind(&song); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "song and group parameters are required"})
	}

	//mock for the test task
	s.getSongDetail(&song)

	err := s.Storage.AddSong(ctx, &song)
	if err != nil {
		fmt.Println("failed to add song: %w", err)
		s.Logger.Error("failed to add song: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to add song"})
	}
	s.Logger.Infof("added song: %s,%s", song.Group, song.Name)
	return c.JSON(http.StatusCreated, echo.Map{"message": "song added successfully"})
}

// HandleEditSong edits an existing song's details
func (s *Service) handleEditSong(c echo.Context) error {

	ctx := c.Request().Context()
	var song storage.Song
	if err := c.Bind(&song); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	err := s.Storage.EditSong(ctx, &song)
	if err != nil {
		s.Logger.Error("failed to edit song: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to edit song"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "song edited successfully"})
}

// HandleDeleteSong deletes a song from the database
func (s *Service) handleDeleteSong(c echo.Context) error {
	songName := c.QueryParam("song")
	groupName := c.QueryParam("group")
	ctx := c.Request().Context()

	if songName == "" || groupName == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "song and group parameters are required"})
	}

	err := s.Storage.DeleteSong(ctx, groupName, songName)
	if err != nil {
		s.Logger.Error("failed to delete song: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to delete song"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "song deleted successfully"})
}
