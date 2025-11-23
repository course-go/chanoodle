package memory

import (
	"github.com/course-go/chanoodle/internal/domain/interfaces/repository"
	"github.com/rs/zerolog"
)

var _ repository.ChannelRepository = &ChannelRepository{}

type ChannelRepository struct {
	log zerolog.Logger
}

func NewChannelRepository(log zerolog.Logger) ChannelRepository {
	return ChannelRepository{
		log: log.With().Str("component", "memory/channel-repository").Logger(),
	}
}
