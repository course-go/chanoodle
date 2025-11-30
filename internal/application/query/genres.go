package query

import (
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/pagination"
)

type Genres struct {
	Pagination pagination.Pagination[entity.Genre]
}

type GenresResult struct {
	Genres []entity.Genre
}
