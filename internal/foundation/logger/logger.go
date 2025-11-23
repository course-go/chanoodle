package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func New(environment string) zerolog.Logger {
	writer := newLoggerWriter(environment)

	return zerolog.New(writer).With().Timestamp().Logger()
}

func newLoggerWriter(environment string) io.Writer {
	if environment == "dev" {
		return zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC822,
		}
	}

	return os.Stderr
}
