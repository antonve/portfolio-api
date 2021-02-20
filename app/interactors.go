package app

import (
	"github.com/antonve/portfolio-api/usecases"
)

// Interactors is a collection of all repositories
type Interactors struct {
	Resume usecases.ResumeInteractor
}

// NewInteractors initializes all repositories
func NewInteractors(
	r *Repositories,
) *Interactors {
	return &Interactors{
		Resume: usecases.NewResumeInteractor(r.Resume),
	}
}
