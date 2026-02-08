package service

import (
	"fmt"

	"github.com/course-go/chanoodle/internal/application/command"
	"github.com/course-go/chanoodle/internal/application/interfaces/service"
	"github.com/course-go/chanoodle/internal/application/query"
	"github.com/course-go/chanoodle/internal/domain/interfaces/repository"
	"github.com/rs/zerolog"
)

var _ service.GenreService = &GenreService{}

type GenreService struct {
	log             zerolog.Logger
	genreRepository repository.GenreRepository
}

func NewGenreService(log zerolog.Logger, genreRepository repository.GenreRepository) *GenreService {
	return &GenreService{
		log:             log.With().Str("component", "application/channel-service").Logger(),
		genreRepository: genreRepository,
	}
}

// Genres implements [service.GenreService].
func (cs *GenreService) Genres(q query.Genres) (r query.GenresResult, err error) {
	genres, err := cs.genreRepository.Genres(&q.Pagination)
	if err != nil {
		return query.GenresResult{}, fmt.Errorf("failed getting genres from repository: %w", err)
	}

	return query.GenresResult{
		Genres: genres,
	}, nil
}

// CreateGenre implements [service.GenreService].
func (cs *GenreService) CreateGenre(c command.CreateGenre) (r command.CreateGenreResult, err error) {
	genre, err := cs.genreRepository.GetOrCreateGenre(c.Genre)
	if err != nil {
		return command.CreateGenreResult{}, fmt.Errorf("failed getting or creating genre in repository: %w", err)
	}

	return command.CreateGenreResult{
		Genre: genre,
	}, nil
}
