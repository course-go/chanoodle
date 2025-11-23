package service

import (
	"github.com/course-go/chanoodle/internal/application/interfaces/service"
	domain "github.com/course-go/chanoodle/internal/domain/interfaces/service"
	"github.com/rs/zerolog"
)

var _ service.EventService = &EventService{}

type ChannelService struct {
	log            zerolog.Logger
	channelService domain.ChannelService
}

func NewChannelService(log zerolog.Logger, channelService domain.ChannelService) ChannelService {
	return ChannelService{
		log:            log.With().Str("component", "application/channel-service").Logger(),
		channelService: channelService,
	}
}
