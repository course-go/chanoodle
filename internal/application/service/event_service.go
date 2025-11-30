package service

import (
	"fmt"

	"github.com/course-go/chanoodle/internal/application/command"
	"github.com/course-go/chanoodle/internal/application/interfaces/service"
	"github.com/course-go/chanoodle/internal/application/query"
	"github.com/course-go/chanoodle/internal/domain/interfaces/repository"
	"github.com/rs/zerolog"
)

var _ service.EventService = &EventService{}

type EventService struct {
	log             zerolog.Logger
	eventRepository repository.EventRepository
}

func NewEventService(log zerolog.Logger, eventRepository repository.EventRepository) *EventService {
	return &EventService{
		log:             log.With().Str("component", "application/event-service").Logger(),
		eventRepository: eventRepository,
	}
}

// Event implements [service.EventService].
func (es *EventService) Event(q query.Event) (r query.EventResult, err error) {
	event, err := es.eventRepository.Event(q.ID)
	if err != nil {
		return query.EventResult{}, fmt.Errorf("failed getting event from service: %w", err)
	}

	return query.EventResult{
		Event: event,
	}, nil
}

// Events implements [service.EventService].
func (es *EventService) Events(q query.Events) (r query.EventsResult, err error) {
	events, err := es.eventRepository.Events(q.Filter, q.Pagination)
	if err != nil {
		return query.EventsResult{}, fmt.Errorf("failed getting events from repository: %w", err)
	}

	return query.EventsResult{
		Events: events,
	}, nil
}

// CreateEvent implements [service.EventService].
func (es *EventService) CreateEvent(c command.CreateEvent) (r command.CreateEventResult, err error) {
	event, err := es.eventRepository.CreateEvent(c.Event)
	if err != nil {
		return command.CreateEventResult{}, fmt.Errorf("failed creating event in repository: %w", err)
	}

	return command.CreateEventResult{
		Event: event,
	}, nil
}

// UpdateEvent implements [service.EventService].
func (es *EventService) UpdateEvent(c command.UpdateEvent) (r command.UpdateEventResult, err error) {
	err = es.eventRepository.UpdateEvent(c.Event)
	if err != nil {
		return command.UpdateEventResult{}, fmt.Errorf("failed updating event in repository: %w", err)
	}

	return command.UpdateEventResult{}, nil
}
