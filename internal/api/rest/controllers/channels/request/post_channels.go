package request

import (
	"fmt"

	"github.com/course-go/chanoodle/internal/api/rest/controllers/channels/dto"
	"github.com/course-go/chanoodle/internal/application/command"
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/labstack/echo/v4"
)

type PostChannels struct {
	Data struct {
		Channel dto.AnonymousChannel `json:"channel"`
	} `json:"data"`
}

func ParsePostChannels(c echo.Context) (cmd command.CreateChannel, err error) {
	var model PostChannels

	err = c.Bind(&model)
	if err != nil {
		return command.CreateChannel{}, fmt.Errorf("failed binding request to model: %w", err)
	}

	return command.CreateChannel{
		Channel: entity.AnonymousChannel{
			Name: model.Data.Channel.Name,
		},
	}, nil
}
