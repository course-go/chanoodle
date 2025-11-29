package service

import (
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/interfaces/repository"
	"github.com/course-go/chanoodle/internal/domain/interfaces/service"
	"github.com/course-go/chanoodle/internal/domain/value/channels"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/course-go/chanoodle/internal/domain/value/pagination"
	"github.com/rs/zerolog"
)

var _ service.ChannelService = &ChannelService{}

type ChannelService struct {
	log               zerolog.Logger
	channelRepository repository.ChannelRepository
}

func NewChannelService(log zerolog.Logger, channelRepository repository.ChannelRepository) *ChannelService {
	return &ChannelService{
		log:               log.With().Str("component", "domain/channel-service").Logger(),
		channelRepository: channelRepository,
	}
}

// Channels implements [service.ChannelService].
func (cs *ChannelService) Channels(
	filter channels.Filter,
	pag pagination.Pagination[entity.Channel],
) ([]entity.Channel, error) {
	return cs.channelRepository.Channels(filter, pag)
}

// Channel implements [service.ChannelService].
func (cs *ChannelService) Channel(id id.ID) (entity.Channel, error) {
	return cs.channelRepository.Channel(id)
}

// CreateChannel implements [service.ChannelService].
func (cs *ChannelService) CreateChannel(anonymousChannel entity.AnonymousChannel) (entity.Channel, error) {
	return cs.channelRepository.CreateChannel([]entity.AnonymousChannel{anonymousChannel})
}

// UpdateChannel implements [service.ChannelService].
func (cs *ChannelService) UpdateChannel(channel entity.Channel) error {
	return cs.channelRepository.UpdateChannel(channel)
}
