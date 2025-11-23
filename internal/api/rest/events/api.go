package events

import application "github.com/course-go/chanoodle/internal/application/interfaces/service"

type API struct {
	eventService application.EventService
}

func NewAPI(eventService application.EventService) API {
	return API{
		eventService: eventService,
	}
}
