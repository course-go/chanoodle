package service

import (
	"github.com/course-go/chanoodle/internal/application/interfaces/service"
	domain "github.com/course-go/chanoodle/internal/domain/interfaces/service"
	"github.com/rs/zerolog"
)

var _ service.EventService = &EventService{}

type EventService struct {
	log          zerolog.Logger
	eventService domain.EventService
}

func NewEventService(log zerolog.Logger, eventService domain.EventService) EventService {
	return EventService{
		log:          log.With().Str("component", "application/event-service").Logger(),
		eventService: eventService,
	}
}
