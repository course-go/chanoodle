package request

import (
	"fmt"

	"github.com/course-go/chanoodle/internal/application/query"
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/channels"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/course-go/chanoodle/internal/domain/value/pagination"
	"github.com/labstack/echo/v4"
)

type GetChannels struct {
	Genres []id.ID `query:"genres"`
	Limit  int     `query:"limit"`
	Offset int     `query:"offset"`
}

func ParseGetChannels(c echo.Context) (q query.Channels, err error) {
	var model GetChannels

	err = c.Bind(&model)
	if err != nil {
		return query.Channels{}, fmt.Errorf("failed binding request to model: %w", err)
	}

	pagination, err := pagination.New[entity.Channel](model.Limit, model.Offset)
	if err != nil {
		return query.Channels{}, fmt.Errorf("failed creating pagination: %w", err)
	}

	return query.Channels{
		Filter: channels.Filter{
			Genres: model.Genres,
		},
		Pagination: pagination,
	}, nil
}
