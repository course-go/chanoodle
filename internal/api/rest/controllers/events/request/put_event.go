package request

import (
	"fmt"

	"github.com/course-go/chanoodle/internal/application/command"
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/labstack/echo/v4"
)

type PutEvent struct {
	ID   id.ID `param:"id" validate:"required"`
	Data struct {
		Event entity.AnonymousEvent `json:"event" validate:"required"`
	} `           validate:"required" json:"data"`
}

func ParsePutEvent(c echo.Context) (cmd command.UpdateEvent, err error) {
	var model PutEvent

	err = c.Bind(&model)
	if err != nil {
		return command.UpdateEvent{}, fmt.Errorf("failed binding request to model: %w", err)
	}

	err = c.Validate(model)
	if err != nil {
		return command.UpdateEvent{}, fmt.Errorf("failed validating request: %w", err)
	}

	return command.UpdateEvent{
		Event: entity.Event{
			ID:   model.ID,
			Name: model.Data.Event.Name,
			From: model.Data.Event.From,
			To:   model.Data.Event.To,
		},
	}, nil
}
