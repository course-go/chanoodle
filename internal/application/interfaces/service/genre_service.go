package service

import (
	"context"

	"github.com/course-go/chanoodle/internal/application/command"
	"github.com/course-go/chanoodle/internal/application/query"
)

// GenreService defines all supported Genre related use-cases.
type GenreService interface {
	Genres(ctx context.Context, q query.Genres) (r query.GenresResult, err error)
	CreateGenre(ctx context.Context, c command.CreateGenre) (r command.CreateGenreResult, err error)
}
