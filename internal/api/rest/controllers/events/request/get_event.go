package request

import (
	"fmt"

	"github.com/course-go/chanoodle/internal/application/query"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/labstack/echo/v4"
)

type GetEvent struct {
	ID id.ID `param:"id"`
}

func ParseGetEvent(c echo.Context) (q query.Event, err error) {
	var model GetEvent

	err = c.Bind(&model)
	if err != nil {
		return query.Event{}, fmt.Errorf("failed binding request to model: %w", err)
	}

	return query.Event{
		ID: model.ID,
	}, nil
}
