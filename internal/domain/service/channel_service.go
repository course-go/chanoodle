package service

import (
	"github.com/course-go/chanoodle/internal/domain/interfaces/repository"
	"github.com/course-go/chanoodle/internal/domain/interfaces/service"
	"github.com/rs/zerolog"
)

var _ service.ChannelService = &ChannelService{}

type ChannelService struct {
	log               zerolog.Logger
	channelRepository repository.ChannelRepository
}

func NewChannelService(log zerolog.Logger, channelRepository repository.ChannelRepository) ChannelService {
	return ChannelService{
		log:               log.With().Str("component", "domain/channel-service").Logger(),
		channelRepository: channelRepository,
	}
}
