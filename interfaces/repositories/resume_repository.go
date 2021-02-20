package repositories

import (
	"github.com/antonve/portfolio-api/domain"
	"github.com/antonve/portfolio-api/interfaces/rdb"
	"github.com/antonve/portfolio-api/usecases"
)

// NewResumeRepository instantiates a new resume repository
func NewResumeRepository(sqlHandler rdb.SQLHandler) usecases.ResumeRepository {
	return &resumeRepository{sqlHandler: sqlHandler}
}

type resumeRepository struct {
	sqlHandler rdb.SQLHandler
}

func (r *resumeRepository) StoreResume(resume *domain.Resume) error {
	query := `
		insert into resume
		(slug, body, enabled)
		values (:slug, :body, :enabled)
	`

	_, err := r.sqlHandler.NamedExecute(query, resume)
	return domain.WrapError(err)
}

func (r *resumeRepository) StoreResumeLog(log *domain.ResumeLog) error {
	return nil
}

func (r *resumeRepository) FindBySlug(slug string) (domain.Resume, error) {
	return domain.Resume{}, nil
}
