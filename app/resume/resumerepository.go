package resume

import (
	"context"

	"github.com/antonve/portfolio-api/infra"
)

type ResumeModel struct {
	Slug      string `db:"slug"`
	Body      string `db:"body"`
	IsVisible bool   `db:"is_visible"`
}

type VisitModel struct {
	ID        uint64 `db:"id"`
	Slug      string `db:"slug"`
	IPAddress string `db:"ip_address"`
	UserAgent string `db:"user_agent"`
}

// NewResumeRepository instantiates a new resume repository
func NewResumeRepository(rdb *infra.RDB) ResumeRepository {
	return &resumeRepository{rdb: rdb}
}

type ResumeRepository interface {
	StoreResume(ctx context.Context, resume *Resume) error
	StoreVisit(ctx context.Context, visit *Visit) (uuid string, err error)
	FindBySlug(ctx context.Context, slug string) (*Resume, error)
}

type resumeRepository struct {
	rdb *infra.RDB
}

func (r *resumeRepository) StoreResume(ctx context.Context, resume *Resume) error {
	query := `
		insert into resume
		(slug, body, is_visible)
		values (:slug, :body, :is_visible)
	`

	_, err := r.rdb.NamedExecContext(ctx, query, domainToResumeModel(resume))
	return err
}

func (r *resumeRepository) StoreVisit(ctx context.Context, visit *Visit) (uuid string, err error) {
	query := `
		insert into resume_logs
		(slug, ip_address, user_agent)
		values ($1, $2, $3)
		returning id
	`

	row := r.rdb.QueryRowContext(ctx, query, visit.Slug(), visit.IPAddress(), visit.UserAgent())
	err = row.Scan(&uuid)

	return
}

func (r *resumeRepository) FindBySlug(ctx context.Context, slug string) (*Resume, error) {
	var model ResumeModel

	query := `
		select slug, body, is_visible
		from resume
		where slug = $1 and is_visible = true
	`

	err := r.rdb.Get(&model, query, slug)
	if err != nil {
		return nil, err
	}

	return resumeModelToDomain(model)
}

func resumeModelToDomain(model ResumeModel) (*Resume, error) {
	return NewResume(model.Slug, model.Body, model.IsVisible)
}

func domainToResumeModel(resume *Resume) *ResumeModel {
	return &ResumeModel{
		Slug:      resume.Slug(),
		Body:      resume.Body(),
		IsVisible: resume.IsVisble(),
	}
}
