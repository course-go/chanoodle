package response

import (
	"github.com/course-go/chanoodle/internal/api/rest/controllers/genres/dto"
	"github.com/course-go/chanoodle/internal/application/query"
)

type GetGenres struct {
	Genres []dto.Genre `json:"genres"`
}

func NewGetGenres(qr query.GenresResult) GetGenres {
	genres := make([]dto.Genre, 0, len(qr.Genres))
	for _, genre := range qr.Genres {
		genres = append(genres, dto.NewGenreFromEntity(genre))
	}

	return GetGenres{
		Genres: genres,
	}
}
