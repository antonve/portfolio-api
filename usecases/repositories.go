//go:generate gex mockgen -source=repositories.go -package usecases -destination=repositories_mock.go

package usecases

import (
	"github.com/antonve/portfolio-api/domain"
)

// ResumeRepository handles Resume related database interactions
type ResumeRepository interface {
	FindBySlug(slug string) (domain.Resume, error)
	StoreResume(resume *domain.Resume) error
	StoreResumeLog(log *domain.ResumeLog) error
}
