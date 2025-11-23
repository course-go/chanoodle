package events

import services "github.com/course-go/chanoodle/internal/application/interfaces/service"

type API struct {
	eventService services.EventService
}

func NewAPI(eventService services.EventService) API {
	return API{
		eventService: eventService,
	}
}
