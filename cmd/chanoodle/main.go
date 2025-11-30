package main

import (
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/course-go/chanoodle/internal/api/rest"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/channels"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/epg"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/events"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/genres"
	"github.com/course-go/chanoodle/internal/api/rest/middleware/auth"
	application "github.com/course-go/chanoodle/internal/application/service"
	"github.com/course-go/chanoodle/internal/config"
	domain "github.com/course-go/chanoodle/internal/domain/service"
	"github.com/course-go/chanoodle/internal/foundation/logger"
	"github.com/course-go/chanoodle/internal/infrastructure/persistence/memory"
	"github.com/rs/zerolog"
)

func main() {
	configPath := flag.String("config", "config/chanoodle.yaml", "path to config file")

	flag.Parse()

	config, err := config.Parse(*configPath)
	if err != nil {
		slog.Error("failed parsing config",
			"error", err,
			"config path", *configPath,
		)

		os.Exit(1)
	}

	log, err := logger.New(config.LogLevel, config.Environment)
	if err != nil {
		slog.Error("failed creating logger",
			"error", err,
			"log level", config.LogLevel,
			"environment", config.Environment,
		)

		os.Exit(1)
	}

	err = runApp(log, config)
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed running app")

		os.Exit(1)
	}
}

func runApp(log zerolog.Logger, config config.Chanoodle) error {
	channelRepository := memory.NewChannelRepository(log)
	eventRepository := memory.NewEventRepository(log)
	genresRepository := memory.NewGenreRepository(log)

	domainEPGService := domain.NewEPGService(log)

	applicationChannelService := application.NewChannelService(log, channelRepository)
	applicationEventService := application.NewEventService(log, eventRepository)
	applicationGenreService := application.NewGenreService(log, genresRepository)
	applicationEPGService := application.NewEPGService(log, domainEPGService, channelRepository, eventRepository)

	channelAPI := channels.NewAPI(log, applicationChannelService)
	eventAPI := events.NewAPI(log, applicationEventService)
	genresAPI := genres.NewAPI(log, applicationGenreService)
	epgAPI := epg.NewAPI(log, applicationEPGService)

	apiKeyAuth := auth.NewAPIKey(config.Auth)

	api := rest.NewAPI(apiKeyAuth, channelAPI, eventAPI, genresAPI, epgAPI)

	router := api.Router(log)

	err := router.Start(config.ListenAddress)
	if !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed running router: %w", err)
	}

	return nil
}
