package response

import (
	"github.com/course-go/chanoodle/internal/api/rest/controllers/genres/dto"
	"github.com/course-go/chanoodle/internal/application/command"
)

type PostGenres struct {
	Genre dto.Genre `json:"genre"`
}

func NewPostGenres(cr command.CreateGenreResult) PostGenres {
	return PostGenres{
		Genre: dto.Genre{
			ID:   int(cr.Genre.ID),
			Name: cr.Genre.Name,
		},
	}
}
