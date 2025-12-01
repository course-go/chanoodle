package service

import (
	"github.com/course-go/chanoodle/internal/application/command"
	"github.com/course-go/chanoodle/internal/application/query"
)

// ChannelService defines all supported Channel related use-cases.
type ChannelService interface {
	Channels(q query.Channels) (r query.ChannelsResult, err error)
	Channel(q query.Channel) (r query.ChannelResult, err error)
	CreateChannel(c command.CreateChannel) (r command.CreateChannelResult, err error)
	UpdateChannel(c command.UpdateChannel) (r command.UpdateChannelResult, err error)
}
