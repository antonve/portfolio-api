package usecases_test

import (
	"testing"

	"github.com/antonve/portfolio-api/domain"
	"github.com/antonve/portfolio-api/usecases"
	"github.com/stretchr/testify/assert"

	gomock "github.com/golang/mock/gomock"
)

func setupResumeTest(t *testing.T) (
	*gomock.Controller,
	*usecases.MockResumeRepository,
	usecases.ResumeInteractor,
) {
	ctrl := gomock.NewController(t)

	repo := usecases.NewMockResumeRepository(ctrl)
	interactor := usecases.NewResumeInteractor(repo)

	return ctrl, repo, interactor
}

func TestResumeInteractor_TrackedFind_ValidResume(t *testing.T) {
	ctrl, repo, interactor := setupResumeTest(t)
	defer ctrl.Finish()

	slug := "test"
	ip := "127.0.0.1"
	agent := "Chrome"
	resume := domain.Resume{
		Slug:    slug,
		Body:    "{}",
		Enabled: true,
	}
	log := domain.ResumeLog{
		Slug:      slug,
		IPAddress: ip,
		UserAgent: agent,
	}

	repo.EXPECT().FindBySlug(slug).Return(resume, nil)
	repo.EXPECT().StoreResumeLog(&log).Return(nil)

	found, err := interactor.TrackedFind(slug, ip, agent)
	assert.NoError(t, err)
	assert.Equal(t, &resume, found)
}

func TestResumeInteractor_TrackedFind_ResumeNotFound(t *testing.T) {
	ctrl, repo, interactor := setupResumeTest(t)
	defer ctrl.Finish()

	slug := "test"
	ip := "127.0.0.1"
	agent := "Chrome"

	repo.EXPECT().FindBySlug(slug).Return(domain.Resume{}, domain.ErrNotFound)

	found, err := interactor.TrackedFind(slug, ip, agent)
	assert.Error(t, err)
	assert.Nil(t, found)
}
