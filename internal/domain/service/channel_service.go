package services

import (
	"github.com/course-go/chanoodle/internal/domain/interfaces/repository"
	"github.com/course-go/chanoodle/internal/domain/interfaces/service"
)

var _ service.ChannelService = &ChannelService{}

type ChannelService struct {
	channelRepository repository.ChannelRepository
}

func NewChannelService(channelRepository repository.ChannelRepository) ChannelService {
	return ChannelService{
		channelRepository: channelRepository,
	}
}
