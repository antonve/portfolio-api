package resume

import (
	"context"
	"database/sql"
	"errors"
)

var ErrResumeNotFound = errors.New("resume not found")

type ResumeService interface {
	GetTrackedResumeVisit(ctx context.Context, visit *Visit) (*Resume, error)
}

type resumeService struct {
	resumeRepository ResumeRepository
}

func NewResumeService(resumeRepository ResumeRepository) ResumeService {
	return &resumeService{
		resumeRepository: resumeRepository,
	}
}

func (s *resumeService) GetTrackedResumeVisit(ctx context.Context, visit *Visit) (*Resume, error) {
	r, err := s.resumeRepository.FindBySlugTracked(ctx, visit)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrResumeNotFound
		default:
			return nil, err
		}
	}

	return r, nil
}
