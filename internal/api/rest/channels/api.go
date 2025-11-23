package channels

import services "github.com/course-go/chanoodle/internal/application/interfaces/service"

type API struct {
	channelService services.ChannelService
}

func NewAPI(channelService services.ChannelService) API {
	return API{
		channelService: channelService,
	}
}
