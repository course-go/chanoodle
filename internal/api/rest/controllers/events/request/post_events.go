package request

import (
	"fmt"

	"github.com/course-go/chanoodle/internal/api/rest/controllers/events/dto"
	"github.com/course-go/chanoodle/internal/application/command"
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/labstack/echo/v4"
)

type PostEvents struct {
	Data struct {
		Event dto.AnonymousEvent `json:"event" validate:"required"`
	} `json:"data" validate:"required"`
}

func ParsePostEvents(c echo.Context) (cmd command.CreateEvent, err error) {
	var model PostEvents

	err = c.Bind(&model)
	if err != nil {
		return command.CreateEvent{}, fmt.Errorf("failed binding request to model: %w", err)
	}

	err = c.Validate(model)
	if err != nil {
		return command.CreateEvent{}, fmt.Errorf("failed validating request: %w", err)
	}

	return command.CreateEvent{
		Event: entity.AnonymousEvent{
			Channel: id.ID(model.Data.Event.Channel),
			Name:    model.Data.Event.Name,
			From:    model.Data.Event.From,
			To:      model.Data.Event.To,
		},
	}, nil
}
