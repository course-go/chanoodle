package request

import (
	"fmt"

	"github.com/course-go/chanoodle/internal/application/command"
	"github.com/course-go/chanoodle/internal/domain/entity"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/labstack/echo/v4"
)

type PutChannel struct {
	ID   id.ID `param:"id"`
	Data struct {
		Channel entity.AnonymousChannel `json:"channel"`
	} `json:"data"`
}

func ParsePutChannel(c echo.Context) (cmd command.UpdateChannel, err error) {
	var model PutChannel

	err = c.Bind(&model)
	if err != nil {
		return command.UpdateChannel{}, fmt.Errorf("failed binding request to model: %w", err)
	}

	return command.UpdateChannel{
		Channel: entity.Channel{
			ID:   model.ID,
			Name: model.Data.Channel.Name,
		},
	}, nil
}
