package app

// get_details_mock serves as a mock instead of providing address of the details API.
// I had not enough time to implement it in a proper way

import (
	"net/http"

	"github.com/gambruh/music_library/internal/storage"
	"github.com/labstack/echo/v4"
)

func (s *Service) getSongDetail(song *storage.Song) error {

	song.ReleaseDate = "19700101"
	song.Text = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas ut arcu magna. Aliquam sodales lacus non ipsum fermentum, id porttitor turpis pellentesque. Nulla at ipsum commodo, cursus mi non, tincidunt nibh. Praesent at nulla in neque ultrices congue in et sapien. Vestibulum id egestas lacus. Fusce porttitor aliquet enim non egestas. Suspendisse condimentum sit amet massa pellentesque efficitur. Donec felis mi, placerat non pretium vel, finibus tincidunt ipsum. "
	song.Link = "localhost"

	return nil
}

// not used though
func (s *Service) handleGetSongDetails(c echo.Context) error {
	songName := c.QueryParam("song")
	groupName := c.QueryParam("group")

	if songName == "" || groupName == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "song and group parameters are required"})
	}

	song := storage.Song{Name: songName, Group: groupName}

	err := s.getSongDetail(&song)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "song not found"})
	}

	return c.JSON(http.StatusOK, song)
}
