package service

import (
	"github.com/course-go/chanoodle/internal/application/command"
	"github.com/course-go/chanoodle/internal/application/query"
)

type GenreService interface {
	Genres(q query.Genres) (r query.GenresResult, err error)
	CreateGenre(c command.CreateGenre) (r command.CreateGenreResult, err error)
}
