package setup

import (
	"testing"

	"github.com/course-go/chanoodle/internal/config"
	"github.com/course-go/chanoodle/internal/foundation/environment"
	"github.com/rs/zerolog"
)

func Config(t *testing.T) config.Chanoodle {
	t.Helper()

	return config.Chanoodle{
		Environment: environment.Development,
		LogLevel:    zerolog.DebugLevel.String(),

		Auth: config.Auth{
			APIKey: "testKey",
		},
		Storage: config.Storage{
			Type: "memory",
		},
	}
}
