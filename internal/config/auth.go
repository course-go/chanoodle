package config

import "errors"

var ErrNoAPIKeySet = errors.New("no API key set")

// Auth represents authentication configuration.
type Auth struct {
	APIKey string `yaml:"api_key"`
}

func defaultAuth() Auth {
	return Auth{}
}

func (a Auth) validate() error {
	if a.APIKey == "" {
		return ErrNoAPIKeySet
	}

	return nil
}
