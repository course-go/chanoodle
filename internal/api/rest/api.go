package rest

import (
	"github.com/course-go/chanoodle/internal/api/rest/channels"
	"github.com/course-go/chanoodle/internal/api/rest/events"
)

type API struct {
	channelsAPI channels.API
	eventsAPI   events.API
}

func NewAPI(channelsAPI channels.API, eventsAPI events.API) API {
	return API{
		channelsAPI: channelsAPI,
		eventsAPI:   eventsAPI,
	}
}
