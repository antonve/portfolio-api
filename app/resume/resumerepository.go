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
	UUID      uint64 `db:"uuid"`
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
	FindBySlug(ctx context.Context, slug string) (*Resume, error)
	FindBySlugTracked(ctx context.Context, visit *Visit) (*Resume, error)
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

func (r *resumeRepository) FindBySlug(ctx context.Context, slug string) (*Resume, error) {
	return r.findBySlug(ctx, r.rdb, slug)
}

func (r *resumeRepository) FindBySlugTracked(ctx context.Context, visit *Visit) (*Resume, error) {
	tx, err := r.rdb.BeginTxx(ctx, nil)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	resume, err := r.findBySlug(ctx, tx, visit.Slug())
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	_, err = r.storeVisit(ctx, tx, visit)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()

	return resume, err
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

func (r *resumeRepository) storeVisit(ctx context.Context, querier infra.Querier, visit *Visit) (uuid string, err error) {
	query := `
		insert into resume_visits
		(uuid, slug, ip_address, user_agent)
		values ($1, $2, $3, $4)
		returning uuid
	`

	row := querier.QueryRowxContext(ctx, query, visit.UUID(), visit.Slug(), visit.IPAddress(), visit.UserAgent())
	err = row.Scan(&uuid)

	return
}

func (r *resumeRepository) findBySlug(ctx context.Context, querier infra.Querier, slug string) (*Resume, error) {
	var model ResumeModel

	query := `
		select slug, body, is_visible
		from resume
		where slug = $1 and is_visible = true
	`

	err := querier.GetContext(ctx, &model, query, slug)
	if err != nil {
		return nil, err
	}

	return resumeModelToDomain(model)
}
