package service

import (
	"fmt"

	"github.com/course-go/chanoodle/internal/application/command"
	"github.com/course-go/chanoodle/internal/application/interfaces/service"
	"github.com/course-go/chanoodle/internal/application/query"
	"github.com/course-go/chanoodle/internal/domain/interfaces/repository"
	"github.com/rs/zerolog"
)

var _ service.ChannelService = &ChannelService{}

type ChannelService struct {
	log               zerolog.Logger
	channelRepository repository.ChannelRepository
}

func NewChannelService(log zerolog.Logger, channelRepository repository.ChannelRepository) *ChannelService {
	return &ChannelService{
		log:               log.With().Str("component", "application/channel-service").Logger(),
		channelRepository: channelRepository,
	}
}

// Channel implements [service.ChannelService].
func (cs *ChannelService) Channel(q query.Channel) (r query.ChannelResult, err error) {
	channel, err := cs.channelRepository.Channel(q.ID)
	if err != nil {
		return query.ChannelResult{}, fmt.Errorf("failed getting channel from repository: %w", err)
	}

	return query.ChannelResult{
		Channel: channel,
	}, nil
}

// Channels implements [service.ChannelService].
func (cs *ChannelService) Channels(q query.Channels) (r query.ChannelsResult, err error) {
	channels, err := cs.channelRepository.Channels(q.Filter, &q.Pagination)
	if err != nil {
		return query.ChannelsResult{}, fmt.Errorf("failed getting channels from repository: %w", err)
	}

	return query.ChannelsResult{
		Channels: channels,
	}, nil
}

// CreateChannel implements [service.ChannelService].
func (cs *ChannelService) CreateChannel(c command.CreateChannel) (r command.CreateChannelResult, err error) {
	channel, err := cs.channelRepository.CreateChannel(c.Channel)
	if err != nil {
		return command.CreateChannelResult{}, fmt.Errorf("failed creating channel in repository: %w", err)
	}

	return command.CreateChannelResult{
		Channel: channel,
	}, nil
}

// UpdateChannel implements [service.ChannelService].
func (cs *ChannelService) UpdateChannel(c command.UpdateChannel) (r command.UpdateChannelResult, err error) {
	err = cs.channelRepository.UpdateChannel(c.Channel)
	if err != nil {
		return command.UpdateChannelResult{}, fmt.Errorf("failed updating channel in repository: %w", err)
	}

	return command.UpdateChannelResult{}, nil
}
