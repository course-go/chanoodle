package service

import (
	"github.com/course-go/chanoodle/internal/application/command"
	"github.com/course-go/chanoodle/internal/application/interfaces/service"
	"github.com/course-go/chanoodle/internal/application/query"
	"github.com/course-go/chanoodle/internal/domain/entity"
	domain "github.com/course-go/chanoodle/internal/domain/interfaces/service"
	"github.com/course-go/chanoodle/internal/domain/value/pagination"
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
		return r, err
	}

	r.Event = event
	return r, nil
}

// Event implements [service.EventService].
func (es *EventService) Event(q query.Event) (r query.EventResult, err error) {
	event, err := es.eventService.Event(q.ID)
	if err != nil {
		return r, err
	}

	r.Event = event
	return r, nil
}

// Events implements [service.EventService].
func (es *EventService) Events(q query.Events) (r query.EventsResult, err error) {
	pag := pagination.New[entity.Event](0, 0)
	events, err := es.eventService.Events(q.Filter, pag)
	if err != nil {
		return r, err
	}

	r.Events = events
	return r, nil
}

// UpdateEvent implements [service.EventService].
func (es *EventService) UpdateEvent(c command.UpdateEvent) (r command.UpdateEventResult, err error) {
	err = es.eventService.UpdateEvent(c.Event)
	if err != nil {
		return r, err
	}

	return r, nil
}
