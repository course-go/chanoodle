package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/course-go/chanoodle/internal/api/rest"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/channels"
	"github.com/course-go/chanoodle/internal/api/rest/controllers/events"
	application "github.com/course-go/chanoodle/internal/application/service"
	domain "github.com/course-go/chanoodle/internal/domain/service"
	"github.com/course-go/chanoodle/internal/infrastructure/persistence/memory"
	"github.com/rs/zerolog"
)

func main() {
	log := zerolog.New(os.Stderr)

	err := runApp(log)
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed running app")

		os.Exit(1)
	}
}

func runApp(log zerolog.Logger) error {
	channelRepository := memory.NewChannelRepository(log)
	eventRepository := memory.NewEventRepository(log)

	domainChannelService := domain.NewChannelService(log, channelRepository)
	domainEventService := domain.NewEventService(log, eventRepository)

	applicationChannelService := application.NewChannelService(log, domainChannelService)
	applicationEventService := application.NewEventService(log, domainEventService)

	channelAPI := channels.NewAPI(log, applicationChannelService)
	eventAPI := events.NewAPI(log, applicationEventService)
	api := rest.NewAPI(channelAPI, eventAPI)

	router := api.Router(log)

	err := router.Start("localhost:8080")
	if !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed running router: %w", err)
	}

	return nil
}
