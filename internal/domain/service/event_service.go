package services

import (
	"github.com/course-go/chanoodle/internal/domain/interfaces/repository"
	"github.com/course-go/chanoodle/internal/domain/interfaces/service"
)

var _ service.EventService = &EventService{}

type EventService struct {
	eventRepository repository.EventRepository
}

func NewEventService(eventRepository repository.EventRepository) EventService {
	return EventService{
		eventRepository: eventRepository,
	}
}
