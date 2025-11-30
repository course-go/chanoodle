package request

import (
	"fmt"

	"github.com/course-go/chanoodle/internal/application/query"
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/pagination"
	"github.com/labstack/echo/v4"
)

type GetGenres struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

func ParseGetGenres(c echo.Context) (q query.Genres, err error) {
	var model GetGenres

	err = c.Bind(&model)
	if err != nil {
		return query.Genres{}, fmt.Errorf("failed binding request to model: %w", err)
	}

	pagination, err := pagination.New[entity.Genre](model.Limit, model.Offset)
	if err != nil {
		return query.Genres{}, fmt.Errorf("failed creating pagination: %w", err)
	}

	return query.Genres{
		Pagination: pagination,
	}, nil
}
