package service

import (
	"context"
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
func (es *EventService) Event(
	ctx context.Context,
	q query.Event,
) (r query.EventResult, err error) {
	event, err := es.eventRepository.Event(ctx, q.ID)
	if err != nil {
		return query.EventResult{}, fmt.Errorf("failed getting event from service: %w", err)
	}

	return query.EventResult{
		Event: event,
	}, nil
}

// Events implements [service.EventService].
func (es *EventService) Events(
	ctx context.Context,
	q query.Events,
) (r query.EventsResult, err error) {
	events, err := es.eventRepository.Events(ctx, q.Filter, &q.Pagination)
	if err != nil {
		return query.EventsResult{}, fmt.Errorf("failed getting events from repository: %w", err)
	}

	return query.EventsResult{
		Events: events,
	}, nil
}

// CreateEvent implements [service.EventService].
func (es *EventService) CreateEvent(
	ctx context.Context,
	c command.CreateEvent,
) (r command.CreateEventResult, err error) {
	event, err := es.eventRepository.CreateEvent(ctx, c.Event)
	if err != nil {
		return command.CreateEventResult{}, fmt.Errorf("failed creating event in repository: %w", err)
	}

	return command.CreateEventResult{
		Event: event,
	}, nil
}

// UpdateEvent implements [service.EventService].
func (es *EventService) UpdateEvent(
	ctx context.Context,
	c command.UpdateEvent,
) (r command.UpdateEventResult, err error) {
	err = es.eventRepository.UpdateEvent(ctx, c.ID, c.Event)
	if err != nil {
		return command.UpdateEventResult{}, fmt.Errorf("failed updating event in repository: %w", err)
	}

	return command.UpdateEventResult{}, nil
}
