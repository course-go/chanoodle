package genres

import (
	"fmt"
	"net/http"

	"github.com/course-go/chanoodle/internal/api/rest/common"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/genres/request"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/genres/response"
	application "github.com/course-go/chanoodle/internal/application/interfaces/service"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

// API represents the Genre REST API subgroup.
type API struct {
	log          zerolog.Logger
	genreService application.GenreService
}

func NewAPI(log zerolog.Logger, genreService application.GenreService) API {
	return API{
		log:          log.With().Str("component", "api-rest/channel-api").Logger(),
		genreService: genreService,
	}
}

func (a *API) MountRoutes(e *echo.Group) {
	channels := e.Group("/genres")

	channels.GET("", a.getGenresController)
	channels.POST("", a.postGenresController)
}

func (a *API) getGenresController(c echo.Context) error {
	q, err := request.ParseGetGenres(c)
	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed parsing request: %w", err)
	}

	qr, err := a.genreService.Genres(q)
	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed querying genres: %w", err)
	}

	data := response.NewGetGenres(qr)

	_ = c.JSON(http.StatusOK, common.NewDataResponse(data))

	return nil
}

func (a *API) postGenresController(c echo.Context) error {
	cmd, err := request.ParsePostGenres(c)
	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed parsing request: %w", err)
	}

	cr, err := a.genreService.CreateGenre(cmd)
	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed creating channel: %w", err)
	}

	data := response.NewPostGenres(cr)

	_ = c.JSON(http.StatusOK, common.NewDataResponse(data))

	return nil
}
