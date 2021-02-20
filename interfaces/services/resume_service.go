package services

import (
	"net/http"

	"github.com/antonve/portfolio-api/domain"
	"github.com/antonve/portfolio-api/usecases"
)

// ResumeService is responsible for handling HTTP requests
type ResumeService interface {
	Get(ctx Context) error
}

// NewResumeService initializer
func NewResumeService(resumeInteractor usecases.ResumeInteractor) ResumeService {
	return &resumeService{
		ResumeInteractor: resumeInteractor,
	}
}

type resumeService struct {
	ResumeInteractor usecases.ResumeInteractor
}

func (s *resumeService) Get(ctx Context) error {
	slug := ctx.Param("slug")
	ip := ctx.RealIP()
	agent := ctx.UserAgent()
	resume, err := s.ResumeInteractor.TrackedFind(slug, ip, agent)

	if err != nil {
		if err == usecases.ErrResumeNotFound {
			return ctx.NoContent(http.StatusNotFound)
		}
		return domain.WrapError(err)
	}

	return ctx.String(http.StatusOK, resume.Body)
}
