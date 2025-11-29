package response

import (
	"github.com/course-go/chanoodle/internal/api/rest/controllers/channels/dto"
	"github.com/course-go/chanoodle/internal/application/query"
)

type GetChannels struct {
	Channels []dto.Channel `json:"channels"`
}

func NewGetChannels(qr query.ChannelsResult) GetChannels {
	channels := make([]dto.Channel, 0, len(qr.Channels))
	for _, channel := range qr.Channels {
		channels = append(channels, dto.NewChannelFromEntity(channel))
	}

	return GetChannels{
		Channels: channels,
	}
}
