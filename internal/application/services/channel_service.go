package services

import (
	"github.com/course-go/chanoodle/internal/application/interfaces/service"
	domain "github.com/course-go/chanoodle/internal/domain/interfaces/service"
)

var _ service.EventService = &EventService{}

type ChannelService struct {
	channelService domain.ChannelService
}

func NewChannelService(channelService domain.ChannelService) ChannelService {
	return ChannelService{
		channelService: channelService,
	}
}
