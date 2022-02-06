package repositories

import (
	"github.com/antonve/portfolio-api/domain"
	"github.com/antonve/portfolio-api/infra"
	"github.com/antonve/portfolio-api/usecases"
)

// NewResumeRepository instantiates a new resume repository
func NewResumeRepository(rdb *infra.RDB) usecases.ResumeRepository {
	return &resumeRepository{rdb: rdb}
}

type resumeRepository struct {
	rdb *infra.RDB
}

func (r *resumeRepository) StoreResume(resume *domain.Resume) error {
	query := `
		insert into resume
		(slug, body, enabled)
		values (:slug, :body, :enabled)
	`

	_, err := r.rdb.NamedExec(query, resume)
	return domain.WrapError(err)
}

func (r *resumeRepository) StoreResumeLog(log *domain.ResumeLog) error {
	query := `
		insert into resume_logs
		(slug, ip_address, user_agent)
		values ($1, $2, $3)
		returning id
	`

	row := r.rdb.QueryRow(query, log.Slug, log.IPAddress, log.UserAgent)
	err := row.Scan(&log.ID)

	return domain.WrapError(err)
}

func (r *resumeRepository) FindBySlug(slug string) (domain.Resume, error) {
	var resume domain.Resume

	query := `
		select slug, body, enabled
		from resume
		where slug = $1 and enabled = true
	`

	err := r.rdb.Get(&resume, query, slug)
	if err != nil {
		return resume, domain.WrapError(err)
	}

	return resume, nil
}
