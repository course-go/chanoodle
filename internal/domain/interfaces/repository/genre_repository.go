package repository

import (
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/pagination"
)

type GenreRepository interface {
	Genres(pagination *pagination.Pagination[entity.Genre]) (genres []entity.Genre, err error)
	GetOrCreateGenre(anonymous entity.AnonymousGenre) (genre entity.Genre, err error)
}
