package rest

import (
	"net/http"
	"time"

	"github.com/course-go/chanoodle/internal/api/rest/common"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/channels"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/events"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/genres"
	"github.com/course-go/chanoodle/internal/api/rest/middleware/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/ziflex/lecho/v3"
)

const requestTimeout = 30 * time.Second

type API struct {
	channelsAPI channels.API
	eventsAPI   events.API
	genresAPI   genres.API
}

func NewAPI(channelsAPI channels.API, eventsAPI events.API, genresAPI genres.API) API {
	return API{
		channelsAPI: channelsAPI,
		eventsAPI:   eventsAPI,
		genresAPI:   genresAPI,
	}
}

func (a *API) Router(log zerolog.Logger) *echo.Echo {
	e := echo.New()

	logger := lecho.From(log)
	e.Logger = logger
	e.HTTPErrorHandler = a.errorHandler

	api := e.Group("/api/v1")
	api.Use(
		middleware.Recover(),
		middleware.Secure(),
		middleware.ContextTimeout(requestTimeout),
		lecho.Middleware(lecho.Config{
			Logger: logger,
		}),
		auth.KeyAuthMiddleware(),
	)

	a.channelsAPI.MountRoutes(api)
	a.eventsAPI.MountRoutes(api)
	a.genresAPI.MountRoutes(api)

	return e
}

func (a *API) errorHandler(err error, c echo.Context) {
	if err != nil {
		return
	}

	// Check if status was already set by handler.
	status := c.Response().Status
	if status == 0 {
		status = http.StatusInternalServerError
	}

	resp := common.NewErrorResponse(err)

	err = c.JSON(status, resp)
	if err != nil {
		_ = c.NoContent(http.StatusInternalServerError)
	}
}
