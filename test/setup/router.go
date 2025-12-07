package setup

import (
	"testing"

	"github.com/course-go/chanoodle/internal/api/rest"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/channels"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/epg"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/events"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/genres"
	"github.com/course-go/chanoodle/internal/api/rest/middleware/auth"
	application "github.com/course-go/chanoodle/internal/application/service"
	"github.com/course-go/chanoodle/internal/config"
	"github.com/course-go/chanoodle/internal/domain/interfaces/repository"
	domain "github.com/course-go/chanoodle/internal/domain/service"
	"github.com/course-go/chanoodle/internal/foundation/logger"
	"github.com/course-go/chanoodle/internal/infrastructure/persistence/memory"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

type Dependencies struct {
	Router *echo.Echo

	ChannelRepository repository.ChannelRepository
	EventRepository   repository.EventRepository
	GenreRepository   repository.GenreRepository
}

func NewDependencies(t *testing.T, config config.Chanoodle) Dependencies {
	t.Helper()

	log, err := logger.New(config.LogLevel, config.Environment)
	require.NoError(t, err)

	mediaRepository := memory.NewMediaRepository(log)

	domainEPGService := domain.NewEPGService(log)

	applicationChannelService := application.NewChannelService(log, mediaRepository)
	applicationEventService := application.NewEventService(log, mediaRepository)
	applicationGenreService := application.NewGenreService(log, mediaRepository)
	applicationEPGService := application.NewEPGService(log, domainEPGService, mediaRepository, mediaRepository)

	channelAPI := channels.NewAPI(log, applicationChannelService)
	eventAPI := events.NewAPI(log, applicationEventService)
	genresAPI := genres.NewAPI(log, applicationGenreService)
	epgAPI := epg.NewAPI(log, applicationEPGService)

	apiKeyAuth := auth.NewAPIKey(config.Auth)

	api := rest.NewAPI(log, apiKeyAuth, channelAPI, eventAPI, genresAPI, epgAPI)

	return Dependencies{
		Router: api.Router(log),

		ChannelRepository: mediaRepository,
		EventRepository:   mediaRepository,
		GenreRepository:   mediaRepository,
	}
}
