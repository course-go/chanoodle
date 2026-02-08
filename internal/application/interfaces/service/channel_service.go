package service

import (
	"context"

	"github.com/course-go/chanoodle/internal/application/command"
	"github.com/course-go/chanoodle/internal/application/query"
)

// ChannelService defines all supported Channel related use-cases.
type ChannelService interface {
	Channels(ctx context.Context, q query.Channels) (r query.ChannelsResult, err error)
	Channel(ctx context.Context, q query.Channel) (r query.ChannelResult, err error)
	CreateChannel(ctx context.Context, c command.CreateChannel) (r command.CreateChannelResult, err error)
	UpdateChannel(ctx context.Context, c command.UpdateChannel) (r command.UpdateChannelResult, err error)
}
