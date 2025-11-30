package command

import "github.com/course-go/chanoodle/internal/domain/entity"

type CreateGenre struct {
	Genre entity.AnonymousGenre
}

type CreateGenreResult struct {
	Genre entity.Genre
}
