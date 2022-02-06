package app

import (
	"github.com/antonve/portfolio-api/infra"
	r "github.com/antonve/portfolio-api/interfaces/repositories"
	"github.com/antonve/portfolio-api/usecases"
)

// Repositories is a collection of all repositories
type Repositories struct {
	Resume usecases.ResumeRepository
}

// NewRepositories initializes all repositories
func NewRepositories(rdb *infra.RDB) *Repositories {
	return &Repositories{
		Resume: r.NewResumeRepository(rdb),
	}
}
