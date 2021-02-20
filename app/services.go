package app

import (
	"github.com/antonve/portfolio-api/interfaces/services"
)

// Services is a collection of all services
type Services struct {
	Health services.HealthService
}

// NewServices initializes all interactors
func NewServices(i *Interactors) *Services {
	return &Services{
		Health: services.NewHealthService(),
	}
}
