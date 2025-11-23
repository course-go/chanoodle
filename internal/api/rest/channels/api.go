package channels

import application "github.com/course-go/chanoodle/internal/application/interfaces/service"

type API struct {
	channelService application.ChannelService
}

func NewAPI(channelService application.ChannelService) API {
	return API{
		channelService: channelService,
	}
}
