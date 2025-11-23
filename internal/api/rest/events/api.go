package events

import (
	application "github.com/course-go/chanoodle/internal/application/interfaces/service"
	"github.com/labstack/echo/v4"
)

type API struct {
	eventService application.EventService
}

func NewAPI(eventService application.EventService) API {
	return API{
		eventService: eventService,
	}
}

func (a *API) MountRoutes(e *echo.Group) {
	events := e.Group("/events")

	events.GET("", a.getEventsController)
	events.GET("/:id", a.getEventController)
	events.POST("", a.postEventsController)
	events.PUT("/:id", a.putEventController)
}

func (a *API) getEventsController(c echo.Context) error {
	return nil
}

func (a *API) getEventController(c echo.Context) error {
	return nil
}

func (a *API) postEventsController(c echo.Context) error {
	return nil
}

func (a *API) putEventController(c echo.Context) error {
	return nil
}
