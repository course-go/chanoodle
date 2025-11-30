package service

import (
	"fmt"

	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/interfaces/repository"
	"github.com/course-go/chanoodle/internal/domain/interfaces/service"
	"github.com/course-go/chanoodle/internal/domain/value/events"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/course-go/chanoodle/internal/domain/value/pagination"
	"github.com/rs/zerolog"
)

var _ service.EventService = &EventService{}

type EventService struct {
	log             zerolog.Logger
	eventRepository repository.EventRepository
}

func NewEventService(log zerolog.Logger, eventRepository repository.EventRepository) *EventService {
	return &EventService{
		log:             log.With().Str("component", "domain/event-service").Logger(),
		eventRepository: eventRepository,
	}
}

// Events implements [service.EventService].
func (es *EventService) Events(
	filter events.Filter,
	pagination pagination.Pagination[entity.Event],
) ([]entity.Event, error) {
	channels, err := es.eventRepository.Events(filter, pagination)
	if err != nil {
		return nil, fmt.Errorf("failed fetching events from repository: %w", err)
	}

	return channels, nil
}

// Event implements [service.EventService].
func (es *EventService) Event(id id.ID) (entity.Event, error) {
	event, err := es.eventRepository.Event(id)
	if err != nil {
		return entity.Event{}, fmt.Errorf("failed fetching event from repository: %w", err)
	}

	return event, nil
}

// CreateEvent implements [service.EventService].
func (es *EventService) CreateEvent(anonymousEvent entity.AnonymousEvent) (entity.Event, error) {
	event, err := es.eventRepository.CreateEvent(anonymousEvent)
	if err != nil {
		return entity.Event{}, fmt.Errorf("failed creating event in repository: %w", err)
	}

	return event, nil
}

// UpdateEvent implements [service.EventService].
func (es *EventService) UpdateEvent(event entity.Event) error {
	err := es.eventRepository.UpdateEvent(event)
	if err != nil {
		return fmt.Errorf("failed updating event in repository: %w", err)
	}

	return nil
}
