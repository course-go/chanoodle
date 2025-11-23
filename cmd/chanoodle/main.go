package main

import (
	"fmt"
	"os"

	"github.com/course-go/chanoodle/internal/api/rest"
	"github.com/course-go/chanoodle/internal/api/rest/channels"
	"github.com/course-go/chanoodle/internal/api/rest/events"
	application "github.com/course-go/chanoodle/internal/application/service"
	domain "github.com/course-go/chanoodle/internal/domain/service"
	"github.com/course-go/chanoodle/internal/infrastructure/persistence/memory"
)

func main() {
	err := runApp()
	if err != nil {
		os.Exit(1)
	}
}

func runApp() error {
	channelRepository := memory.NewChannelRepository()
	eventRepository := memory.NewEventRepository()

	domainChannelService := domain.NewChannelService(channelRepository)
	domainEventService := domain.NewEventService(eventRepository)

	applicationChannelService := application.NewChannelService(domainChannelService)
	applicationEventService := application.NewEventService(domainEventService)

	channelAPI := channels.NewAPI(applicationChannelService)
	eventAPI := events.NewAPI(applicationEventService)
	api := rest.NewAPI(channelAPI, eventAPI)

	router := api.Router()

	err := router.Start("localhost:8080")
	if err != nil {
		return fmt.Errorf("failed running router: %w", err)
	}

	return nil
}
