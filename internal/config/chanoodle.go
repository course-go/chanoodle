package config

import (
	"fmt"
	"os"
	"slices"

	"github.com/course-go/chanoodle/internal/foundation/environment"
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
)

// Chanoodle represents the main application config.
type Chanoodle struct {
	Environment   environment.Environment `yaml:"environment"`
	LogLevel      string                  `yaml:"log_level"`
	ListenAddress string                  `yaml:"listen_address"`

	Auth    Auth    `yaml:"auth"`
	Storage Storage `yaml:"storage"`
}

// Parse parses the config from the given file path.
func Parse(filePath string) (config Chanoodle, err error) {
	config = defaultChanoodle()

	data, err := os.ReadFile(filePath) //nolint: gosec
	if err != nil {
		return Chanoodle{}, fmt.Errorf("failed reading config file: %w", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Chanoodle{}, fmt.Errorf("failed parsing config file: %w", err)
	}

	err = config.validate()
	if err != nil {
		return Chanoodle{}, fmt.Errorf("failed validation config: %w", err)
	}

	return config, nil
}

func (c *Chanoodle) validate() error {
	if !slices.Contains(environment.Environments(), c.Environment) {
		return environment.ErrUnknownEnvironment
	}

	err := c.Auth.validate()
	if err != nil {
		return fmt.Errorf("auth config validation failed: %w", err)
	}

	err = c.Storage.validate()
	if err != nil {
		return fmt.Errorf("storage config validation failed: %w", err)
	}

	return nil
}

func defaultChanoodle() Chanoodle {
	return Chanoodle{
		Environment:   environment.Production,
		ListenAddress: "localhost:8080",
		LogLevel:      zerolog.InfoLevel.String(),

		Auth:    defaultAuth(),
		Storage: defaultStorage(),
	}
}
