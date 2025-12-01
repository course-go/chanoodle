package repository

// MediaRepository represents a repository for managing all media-related entities.
type MediaRepository interface {
	ChannelRepository
	EventRepository
	GenreRepository
}
