package config

import (
	"fmt"
	"os"
	"slices"

	"github.com/course-go/chanoodle/internal/foundation/environment"
	"gopkg.in/yaml.v3"
)

type Chanoodle struct {
	Environment   environment.Environment `yaml:"environment"`
	LogLevel      string                  `yaml:"log_level"`
	ListenAddress string                  `yaml:"listen_address"`
}

// Parse parses the config from the given file path.
func Parse(filePath string) (config Chanoodle, err error) {
	config = newDefaultChanoodle()

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

	return nil
}

func newDefaultChanoodle() Chanoodle {
	return Chanoodle{
		Environment:   environment.Production,
		ListenAddress: "localhost:8080",
	}
}
