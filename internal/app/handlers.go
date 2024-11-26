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

// HandleGetSong godoc
// @Summary Get song
// @Description Get song details from the database using group name and song name
// @ID get-song
// @Produce json
// @Success 200 {array} storage.Song
// @Failure 400 {object} string "Song and group parameters are required"
// @Failure 404 {object} string "Song not found"
// @Router /api/getsong [post]
func (s *Service) HandleGetSong(c echo.Context) error {
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

// HandleAddSong godoc
// @Summary Add song
// @Description Add song to the database, using group and song name, fetching details from other API
// @ID add-song
// @Produce json
// @Success 200 {array} string
// @Failure 400 {object} string "Song and group parameters are required"
// @Failure 500 {object} string "Failed to add song"
// @Router /api/addsong [post]
func (s *Service) HandleAddSong(c echo.Context) error {
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

// HandleEditSong godoc
// @Summary Edit song
// @Description Edit song details, using group and song to identify the database line, then text, link, releaseDate (of tape Date)
// @ID edit-song
// @Produce json
// @Success 200 {array} string
// @Failure 400 {object} string "Invalid request body or missing parameters"
// @Failure 500 {object} string "failed to edit song"
// @Router /api/editsong [post]
func (s *Service) HandleEditSong(c echo.Context) error {

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

// HandleDeleteSong
// @Summary Delete song
// @Description Deletes a song from the database, using song name and group name to identify it
// @ID del-song
// @Produce json
// @Success 200 {array} string
// @Failure 400 {object} string "Invalid request body or missing parameters"
// @Failure 500 {object} string "Failed to delete song"
// @Router /api/delsong [post]
func (s *Service) HandleDeleteSong(c echo.Context) error {

	ctx := c.Request().Context()
	var song storage.Song
	if err := c.Bind(&song); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}
	if song.Name == "" || song.Group == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "song and group parameters are required"})
	}

	err := s.Storage.DeleteSong(ctx, song.Group, song.Name)
	if err != nil {
		s.Logger.Error("failed to delete song: %w", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to delete song"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "song deleted successfully"})
}
