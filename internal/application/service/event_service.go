package service

import (
	"github.com/course-go/chanoodle/internal/application/interfaces/service"
	domain "github.com/course-go/chanoodle/internal/domain/interfaces/service"
)

var _ service.EventService = &EventService{}

type EventService struct {
	eventService domain.EventService
}

func NewEventService(eventService domain.EventService) EventService {
	return EventService{
		eventService: eventService,
	}
}
