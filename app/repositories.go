package app

import (
	"github.com/antonve/portfolio-api/interfaces/rdb"
)

// Repositories is a collection of all repositories
type Repositories struct {
}

// NewRepositories initializes all repositories
func NewRepositories(sh rdb.SQLHandler) *Repositories {
	return &Repositories{}
}
