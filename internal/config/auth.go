package config

import "errors"

var ErrNoAPIKeySet = errors.New("no API key set")

type Auth struct {
	APIKey string `yaml:"api_key"`
}
