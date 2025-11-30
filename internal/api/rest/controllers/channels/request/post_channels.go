package request

import (
	"fmt"

	"github.com/course-go/chanoodle/internal/api/rest/controllers/channels/dto"
	"github.com/course-go/chanoodle/internal/application/command"
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/priority"
	"github.com/labstack/echo/v4"
)

type PostChannels struct {
	Data struct {
		Channel dto.AnonymousChannel `json:"channel" validate:"required"`
	} `json:"data" validate:"required"`
}

func ParsePostChannels(c echo.Context) (cmd command.CreateChannel, err error) {
	var model PostChannels

	err = c.Bind(&model)
	if err != nil {
		return command.CreateChannel{}, fmt.Errorf("failed binding request to model: %w", err)
	}

	err = c.Validate(model)
	if err != nil {
		return command.CreateChannel{}, fmt.Errorf("failed validating request: %w", err)
	}

	return command.CreateChannel{
		Channel: entity.AnonymousChannel{
			Name:     model.Data.Channel.Name,
			Priority: priority.Priority(model.Data.Channel.Priority),
		},
	}, nil
}
