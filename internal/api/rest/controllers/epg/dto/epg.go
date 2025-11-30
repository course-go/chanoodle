package dto

import "github.com/course-go/chanoodle/internal/domain/value/epg"

type EPG struct {
	Channels []Channel `json:"channels,omitzero"`
}

func NewEPGFromValue(value epg.EPG) EPG {
	channels := make([]Channel, 0, len(value.Channels))
	for _, channel := range value.Channels {
		channels = append(channels, NewChannelFromValue(channel))
	}

	return EPG{
		Channels: channels,
	}
}
