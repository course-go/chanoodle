package channels

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/course-go/chanoodle/internal/api/rest/common"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/channels/request"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/channels/response"
	application "github.com/course-go/chanoodle/internal/application/interfaces/service"
	"github.com/course-go/chanoodle/internal/domain/value/id"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

// API represents the Channel REST API subgroup.
type API struct {
	log            zerolog.Logger
	channelService application.ChannelService
}

func NewAPI(log zerolog.Logger, channelService application.ChannelService) API {
	return API{
		log:            log.With().Str("component", "api-rest/channel-api").Logger(),
		channelService: channelService,
	}
}

func (a *API) MountRoutes(e *echo.Group) {
	channels := e.Group("/channels")

	channels.GET("", a.getChannelsController)
	channels.GET("/:id", a.getChannelController)
	channels.POST("", a.postChannelsController)
	channels.PUT("/:id", a.putChannelController)
}

func (a *API) getChannelsController(c echo.Context) error {
	q, err := request.ParseGetChannels(c)
	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed parsing request: %w", err)
	}

	qr, err := a.channelService.Channels(q)
	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed querying channels: %w", err)
	}

	data := response.NewGetChannels(qr)

	_ = c.JSON(http.StatusOK, common.NewDataResponse(data))

	return nil
}

func (a *API) getChannelController(c echo.Context) error {
	q, err := request.ParseGetChannel(c)
	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed parsing request: %w", err)
	}

	qr, err := a.channelService.Channel(q)
	if errors.Is(err, id.ErrNoSuchEntity) {
		c.Response().Status = http.StatusNotFound

		return fmt.Errorf("channel not found: %w", err)
	}

	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed querying channel: %w", err)
	}

	data := response.NewGetChannel(qr)

	_ = c.JSON(http.StatusOK, common.NewDataResponse(data))

	return nil
}

func (a *API) postChannelsController(c echo.Context) error {
	cmd, err := request.ParsePostChannels(c)
	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed parsing request: %w", err)
	}

	cr, err := a.channelService.CreateChannel(cmd)
	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed creating channel: %w", err)
	}

	data := response.ParsePostChannels(cr)

	_ = c.JSON(http.StatusOK, common.NewDataResponse(data))

	return nil
}

func (a *API) putChannelController(c echo.Context) error {
	cmd, err := request.ParsePutChannel(c)
	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed parsing request: %w", err)
	}

	cr, err := a.channelService.UpdateChannel(cmd)
	if err != nil {
		c.Response().Status = http.StatusBadRequest

		return fmt.Errorf("failed updating channel: %w", err)
	}

	data := response.ParsePutChannel(cr)

	_ = c.JSON(http.StatusOK, common.NewDataResponse(data))

	return nil
}
