package request

import (
	"fmt"

	"github.com/course-go/chanoodle/internal/application/query"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/labstack/echo/v4"
)

type GetChannel struct {
	ID id.ID `param:"id" validate:"required"`
}

func ParseGetChannel(c echo.Context) (q query.Channel, err error) {
	var model GetChannel

	err = c.Bind(&model)
	if err != nil {
		return query.Channel{}, fmt.Errorf("failed binding request to model: %w", err)
	}

	err = c.Validate(model)
	if err != nil {
		return query.Channel{}, fmt.Errorf("failed validating request: %w", err)
	}

	return query.Channel{
		ID: model.ID,
	}, nil
}
