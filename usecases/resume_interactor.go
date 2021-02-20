//go:generate gex mockgen -source=resume_interactor.go -package usecases -destination=resume_interactor_mock.go

package usecases

import (
	"github.com/srvc/fail"

	"github.com/antonve/portfolio-api/domain"
)

// ErrResumeNotFound for when no resume could be found
var ErrResumeNotFound = fail.New("no resume could be found")

// ResumeInteractor contains all business logic for resumes
type ResumeInteractor interface {
	TrackedFind(slug, ipAddress, userAgent string) (*domain.Resume, error)
}

func NewResumeInteractor(resumeRepository ResumeRepository) ResumeInteractor {
	return &resumeInteractor{
		resumeRepository: resumeRepository,
	}
}

type resumeInteractor struct {
	resumeRepository ResumeRepository
}

func (i *resumeInteractor) TrackedFind(slug, ipAddress, userAgent string) (*domain.Resume, error) {
	resume, err := i.resumeRepository.FindBySlug(slug)

	if err != nil {
		if err == domain.ErrNotFound {
			return nil, ErrResumeNotFound
		}

		return nil, domain.WrapError(err)
	}

	log := &domain.ResumeLog{
		Slug:      slug,
		IPAddress: ipAddress,
		UserAgent: userAgent,
	}
	err = i.resumeRepository.StoreResumeLog(log)

	if err != nil {
		return nil, domain.WrapError(err)
	}

	return &resume, nil
}
