package service

import (
	"github.com/course-go/chanoodle/internal/domain/interfaces/repository"
	"github.com/course-go/chanoodle/internal/domain/interfaces/service"
	"github.com/rs/zerolog"
)

var _ service.EventService = &EventService{}

type EventService struct {
	log             zerolog.Logger
	eventRepository repository.EventRepository
}

func NewEventService(log zerolog.Logger, eventRepository repository.EventRepository) EventService {
	return EventService{
		log:             log.With().Str("component", "domain/event-service").Logger(),
		eventRepository: eventRepository,
	}
}
