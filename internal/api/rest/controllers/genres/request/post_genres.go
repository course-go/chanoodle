package request

import (
	"fmt"

	"github.com/course-go/chanoodle/internal/api/rest/controllers/genres/dto"
	"github.com/course-go/chanoodle/internal/application/command"
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/labstack/echo/v4"
)

type PostGenres struct {
	Data struct {
		Genre dto.AnonymousGenre `json:"genre"`
	} `json:"data"`
}

func ParsePostGenres(c echo.Context) (cmd command.CreateGenre, err error) {
	var model PostGenres

	err = c.Bind(&model)
	if err != nil {
		return command.CreateGenre{}, fmt.Errorf("failed binding request to model: %w", err)
	}

	return command.CreateGenre{
		Genre: entity.AnonymousGenre{
			Name: model.Data.Genre.Name,
		},
	}, nil
}
