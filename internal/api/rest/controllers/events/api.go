package events

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/course-go/chanoodle/internal/api/rest/common"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/events/request"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/events/response"
	application "github.com/course-go/chanoodle/internal/application/interfaces/service"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

// API represents the Event REST API subgroup.
type API struct {
	log          zerolog.Logger
	eventService application.EventService
}

func NewAPI(log zerolog.Logger, eventService application.EventService) API {
	return API{
		log:          log.With().Str("component", "api-rest/event-api").Logger(),
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
	ctx := c.Request().Context()

	q, err := request.ParseGetEvents(c)
	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed parsing request: %w", err)
	}

	qr, err := a.eventService.Events(ctx, q)
	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed querying events: %w", err)
	}

	data := response.ParseGetEvents(qr)

	err = c.JSON(http.StatusOK, common.NewDataResponse(data))
	if err != nil {
		return fmt.Errorf("failed sending response: %w", err)
	}

	return nil
}

func (a *API) getEventController(c echo.Context) error {
	ctx := c.Request().Context()

	q, err := request.ParseGetEvent(c)
	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed parsing request: %w", err)
	}

	qr, err := a.eventService.Event(ctx, q)
	if errors.Is(err, id.ErrNoSuchEntity) {
		c.Response().Status = http.StatusNotFound

		return fmt.Errorf("event not found: %w", err)
	}

	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed querying event: %w", err)
	}

	data := response.ParseGetEvent(qr)

	err = c.JSON(http.StatusOK, common.NewDataResponse(data))
	if err != nil {
		return fmt.Errorf("failed sending response: %w", err)
	}

	return nil
}

func (a *API) postEventsController(c echo.Context) error {
	ctx := c.Request().Context()

	cmd, err := request.ParsePostEvents(c)
	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed parsing request: %w", err)
	}

	cr, err := a.eventService.CreateEvent(ctx, cmd)
	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed creating event: %w", err)
	}

	data := response.ParsePostEvents(cr)

	err = c.JSON(http.StatusOK, common.NewDataResponse(data))
	if err != nil {
		return fmt.Errorf("failed sending response: %w", err)
	}

	return nil
}

func (a *API) putEventController(c echo.Context) error {
	ctx := c.Request().Context()

	cmd, err := request.ParsePutEvent(c)
	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed parsing request: %w", err)
	}

	cr, err := a.eventService.UpdateEvent(ctx, cmd)
	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed updating event: %w", err)
	}

	data := response.ParsePutEvent(cr)

	err = c.JSON(http.StatusOK, common.NewDataResponse(data))
	if err != nil {
		return fmt.Errorf("failed sending response: %w", err)
	}

	return nil
}
