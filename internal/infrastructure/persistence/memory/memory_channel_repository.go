package memory

import "github.com/course-go/chanoodle/internal/domain/interfaces/repository"

var _ repository.ChannelRepository = &ChannelRepository{}

type ChannelRepository struct{}

func NewChannelRepository() ChannelRepository {
	return ChannelRepository{}
}
