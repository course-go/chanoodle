package request

import (
	"fmt"
	"time"

	"github.com/course-go/chanoodle/internal/application/query"
	"github.com/labstack/echo/v4"
)

type GetEPG struct {
	From int `query:"from"`
	To   int `query:"to"`
}

func ParseGetEPG(c echo.Context) (q query.EPG, err error) {
	var model GetEPG

	err = c.Bind(&model)
	if err != nil {
		return query.EPG{}, fmt.Errorf("failed binding request to model: %w", err)
	}

	return query.EPG{
		From: time.Unix(int64(model.From), 0),
		To:   time.Unix(int64(model.To), 0),
	}, nil
}
