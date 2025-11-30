package service

import (
	"fmt"

	"github.com/course-go/chanoodle/internal/application/command"
	"github.com/course-go/chanoodle/internal/application/interfaces/service"
	"github.com/course-go/chanoodle/internal/application/query"
	domain "github.com/course-go/chanoodle/internal/domain/interfaces/service"
	"github.com/rs/zerolog"
)

var _ service.EventService = &EventService{}

type EventService struct {
	log          zerolog.Logger
	eventService domain.EventService
}

func NewEventService(log zerolog.Logger, eventService domain.EventService) *EventService {
	return &EventService{
		log:          log.With().Str("component", "application/event-service").Logger(),
		eventService: eventService,
	}
}

// CreateEvent implements [service.EventService].
func (es *EventService) CreateEvent(c command.CreateEvent) (r command.CreateEventResult, err error) {
	event, err := es.eventService.CreateEvent(c.Event)
	if err != nil {
		return command.CreateEventResult{}, fmt.Errorf("failed creating event in service: %w", err)
	}

	return command.CreateEventResult{
		Event: event,
	}, nil
}

// Event implements [service.EventService].
func (es *EventService) Event(q query.Event) (r query.EventResult, err error) {
	event, err := es.eventService.Event(q.ID)
	if err != nil {
		return query.EventResult{}, fmt.Errorf("failed fetching event from service: %w", err)
	}

	return query.EventResult{
		Event: event,
	}, nil
}

// Events implements [service.EventService].
func (es *EventService) Events(q query.Events) (r query.EventsResult, err error) {
	events, err := es.eventService.Events(q.Filter, q.Pagination)
	if err != nil {
		return query.EventsResult{}, fmt.Errorf("failed fetching events from service: %w", err)
	}

	return query.EventsResult{
		Events: events,
	}, nil
}

// UpdateEvent implements [service.EventService].
func (es *EventService) UpdateEvent(c command.UpdateEvent) (r command.UpdateEventResult, err error) {
	err = es.eventService.UpdateEvent(c.Event)
	if err != nil {
		return command.UpdateEventResult{}, fmt.Errorf("failed updating event in service: %w", err)
	}

	return command.UpdateEventResult{}, nil
}
