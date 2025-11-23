package rest

import (
	"github.com/course-go/chanoodle/internal/api/rest/channels"
	"github.com/course-go/chanoodle/internal/api/rest/events"
	"github.com/labstack/echo/v4"
)

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

func (a *API) Router() *echo.Echo {
	e := echo.New()
	api := e.Group("/api/v1")

	a.channelsAPI.MountRoutes(api)
	a.eventsAPI.MountRoutes(api)

	return e
}
