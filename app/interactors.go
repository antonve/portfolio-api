package app

// Interactors is a collection of all repositories
type Interactors struct {
}

// NewInteractors initializes all repositories
func NewInteractors(
	r *Repositories,
) *Interactors {
	return &Interactors{}
}
