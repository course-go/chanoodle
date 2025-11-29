package memory

import (
	"sync"

	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/interfaces/repository"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/course-go/chanoodle/internal/domain/value/pagination"
	"github.com/rs/zerolog"
)

var _ repository.GenreRepository = &GenreRepository{}

type GenreRepository struct {
	log zerolog.Logger

	mu      sync.Mutex
	genreID id.ID
	genres  map[id.ID]entity.Genre
}

func NewGenreRepository(log zerolog.Logger) *GenreRepository {
	return &GenreRepository{
		log:    log.With().Str("component", "memory/genre-repository").Logger(),
		genres: make(map[id.ID]entity.Genre),
	}
}

// Genres implements [repository.GenreRepository].
func (e *GenreRepository) Genres(pagination pagination.Pagination[entity.Genre]) (genres []entity.Genre, err error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	genres = make([]entity.Genre, 0, len(e.genres))
	for _, genre := range e.genres {
		genres = append(genres, genre)
	}

	genres = pagination.Paginate(genres)

	return genres, nil
}

// GetOrCreateGenre implements [repository.GenreRepository].
func (e *GenreRepository) GetOrCreateGenre(anonymousGenre entity.AnonymousGenre) (genre entity.Genre, err error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	for _, genre := range e.genres {
		if genre.Name == anonymousGenre.Name {
			return genre, nil
		}
	}

	e.genreID++
	genre = anonymousGenre.ToGenre(e.genreID)
	e.genres[e.genreID] = genre

	return genre, nil
}
