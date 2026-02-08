package request

import (
	"fmt"

	"github.com/course-go/chanoodle/internal/application/query"
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/events"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/course-go/chanoodle/internal/domain/value/pagination"
	"github.com/labstack/echo/v4"
)

type GetEvents struct {
	Channels []id.ID `query:"channels"`
	Genres   []id.ID `query:"genres"`
	Limit    int     `query:"limit"`
	Offset   int     `query:"offset"`
}

func ParseGetEvents(c echo.Context) (q query.Events, err error) {
	var model GetEvents

	err = c.Bind(&model)
	if err != nil {
		return query.Events{}, fmt.Errorf("failed binding request to model: %w", err)
	}

	pagination, err := pagination.New[entity.Event](model.Limit, model.Offset)
	if err != nil {
		return query.Events{}, fmt.Errorf("failed creating pagination: %w", err)
	}

	return query.Events{
		Filter: events.Filter{
			Channels: model.Channels,
			Genres:   model.Genres,
		},
		Pagination: pagination,
	}, nil
}
