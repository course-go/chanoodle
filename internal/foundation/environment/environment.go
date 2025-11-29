package environment

import "errors"

var ErrUnknownEnvironment = errors.New("unknown environment")

type Environment string

const (
	Development Environment = "development"
	Production  Environment = "production"
)

func Environments() []Environment {
	return []Environment{
		Development,
		Production,
	}
}
