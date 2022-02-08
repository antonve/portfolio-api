package domain

import "errors"

// Environment of the app
type Environment string

// Environment enum
const (
	EnvProduction  Environment = "production"
	EnvDevelopment Environment = "development"
	EnvTest        Environment = "test"
)

// AllEnvironments a collection of all valid environments
var AllEnvironments = []Environment{
	EnvProduction,
	EnvDevelopment,
	EnvTest,
}

// ErrInvalidEnvironment for when an environment is not defined in our app
var ErrInvalidEnvironment = errors.New("supplied environment is not supported")

// Validate wether or the the environment is a valid one
func (env Environment) Validate() error {
	for _, e := range AllEnvironments {
		if e == env {
			return nil
		}
	}

	return ErrInvalidEnvironment
}
