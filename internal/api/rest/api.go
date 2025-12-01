package rest

import (
	"net/http"
	"time"

	"github.com/course-go/chanoodle/internal/api/rest/common"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/channels"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/epg"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/events"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/genres"
	"github.com/course-go/chanoodle/internal/api/rest/middleware/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/ziflex/lecho/v3"
)

const requestTimeout = 30 * time.Second

// API represents the REST API.
type API struct {
	apiKeyAuth auth.APIKey

	channelsAPI channels.API
	eventsAPI   events.API
	genresAPI   genres.API
	epgAPI      epg.API

	validator *common.Validator
}

func NewAPI(
	apiKeyAuth auth.APIKey,
	channelsAPI channels.API,
	eventsAPI events.API,
	genresAPI genres.API,
	epgAPI epg.API,
) API {
	return API{
		apiKeyAuth: apiKeyAuth,

		channelsAPI: channelsAPI,
		eventsAPI:   eventsAPI,
		genresAPI:   genresAPI,
		epgAPI:      epgAPI,

		validator: common.NewValidator(),
	}
}

func (a *API) Router(log zerolog.Logger) *echo.Echo {
	e := echo.New()

	logger := lecho.From(log)
	e.Logger = logger
	e.HTTPErrorHandler = a.errorHandler
	e.Validator = a.validator

	api := e.Group("/api/v1")
	api.Use(
		middleware.Recover(),
		middleware.Secure(),
		a.apiKeyAuth.KeyAuthMiddleware(),
		middleware.ContextTimeout(requestTimeout),
		lecho.Middleware(lecho.Config{
			Logger: logger,
		}),
	)

	a.channelsAPI.MountRoutes(api)
	a.eventsAPI.MountRoutes(api)
	a.genresAPI.MountRoutes(api)
	a.epgAPI.MountRoutes(api)

	return e
}

func (a *API) errorHandler(err error, c echo.Context) {
	if err == nil {
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
