package channels

import (
	application "github.com/course-go/chanoodle/internal/application/interfaces/service"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

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
	return nil
}

func (a *API) getChannelController(c echo.Context) error {
	return nil
}

func (a *API) postChannelsController(c echo.Context) error {
	return nil
}

func (a *API) putChannelController(c echo.Context) error {
	return nil
}
