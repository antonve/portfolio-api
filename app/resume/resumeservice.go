package resume

import (
	"context"
	"database/sql"
	"errors"
)

var ErrResumeNotFound = errors.New("resume not found")

type ResumeService interface {
	FindResumeBySlug(ctx context.Context, slug string) (*Resume, error)
}

type resumeService struct {
	resumeRepository ResumeRepository
}

func NewResumeService(resumeRepository ResumeRepository) ResumeService {
	return &resumeService{
		resumeRepository: resumeRepository,
	}
}

func (s *resumeService) FindResumeBySlug(ctx context.Context, slug string) (*Resume, error) {
	r, err := s.resumeRepository.FindBySlug(ctx, slug)

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
