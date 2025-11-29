package service

import (
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
func (es *EventService) Events(filter events.Filter, pag pagination.Pagination[entity.Event]) ([]entity.Event, error) {
	return es.eventRepository.Events(filter, pag)
}

// Event implements [service.EventService].
func (es *EventService) Event(id id.ID) (entity.Event, error) {
	return es.eventRepository.Event(id)
}

// CreateEvent implements [service.EventService].
func (es *EventService) CreateEvent(anonymousEvent entity.AnonymousEvent) (entity.Event, error) {
	return es.eventRepository.CreateEvent(anonymousEvent)
}

// UpdateEvent implements [service.EventService].
func (es *EventService) UpdateEvent(event entity.Event) error {
	return es.eventRepository.UpdateEvent(event)
}
