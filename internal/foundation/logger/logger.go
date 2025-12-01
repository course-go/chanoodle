package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/course-go/chanoodle/internal/foundation/environment"
	"github.com/rs/zerolog"
)

// New creates a new zerolog.Logger based on the provided log level and environment.
func New(rawLevel string, env environment.Environment) (log zerolog.Logger, err error) {
	writer, err := newLoggerWriter(env)
	if err != nil {
		return zerolog.Logger{}, fmt.Errorf("failed creating logger writer: %w", err)
	}

	level, err := zerolog.ParseLevel(rawLevel)
	if err != nil {
		return zerolog.Logger{}, fmt.Errorf("failed parsing log level: %w", err)
	}

	return zerolog.New(writer).Level(level).With().Timestamp().Logger(), nil
}

func newLoggerWriter(env environment.Environment) (writer io.Writer, err error) {
	switch env {
	case environment.Production:
		return os.Stderr, nil
	case environment.Development:
		return zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC822,
		}, nil
	default:
		return nil, environment.ErrUnknownEnvironment
	}
}
