package rest

import (
	"time"

	"github.com/course-go/chanoodle/internal/api/rest/controllers/channels"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/events"
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
}

func NewAPI(channelsAPI channels.API, eventsAPI events.API) API {
	return API{
		channelsAPI: channelsAPI,
		eventsAPI:   eventsAPI,
	}
}

func (a *API) Router(log zerolog.Logger) *echo.Echo {
	e := echo.New()

	logger := lecho.From(log)
	e.Logger = logger

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

	return e
}
