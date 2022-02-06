package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/antonve/portfolio-api/domain"
	"github.com/antonve/portfolio-api/interfaces/repositories"
)

func TestResumeRepository_StoreResume(t *testing.T) {
	rdb, cleanup := setupTestingSuite(t)
	defer cleanup()

	repo := repositories.NewResumeRepository(rdb)

	resume := &domain.Resume{
		Slug:    "test",
		Body:    "{}",
		Enabled: true,
	}

	err := repo.StoreResume(resume)
	assert.NoError(t, err)
}

func TestResumeRepository_StoreResumeLog(t *testing.T) {
	sqlHandler, cleanup := setupTestingSuite(t)
	defer cleanup()

	repo := repositories.NewResumeRepository(sqlHandler)

	log := &domain.ResumeLog{
		Slug:      "test",
		IPAddress: "127.0.0.1",
		UserAgent: "Firefox",
	}

	err := repo.StoreResumeLog(log)
	assert.NoError(t, err)
	assert.NotEqual(t, uint64(0), log.ID)
}

func TestResumeRepository_FindBySlug(t *testing.T) {
	sqlHandler, cleanup := setupTestingSuite(t)
	defer cleanup()

	repo := repositories.NewResumeRepository(sqlHandler)

	tests := []struct {
		name    string
		slug    string
		enabled bool
		body    string
	}{
		{"existing resume", "test", true, "{\"foo\": \"foo\"}"},
		{"disabled resume", "disabled", false, "{\"foo\": \"bar\"}"},
		{"missing resume", "missing", false, "{\"bar\": \"foo\"}"},
	}

	for _, test := range tests[:2] {
		resume := &domain.Resume{
			Slug:    test.slug,
			Body:    test.body,
			Enabled: test.enabled,
		}

		err := repo.StoreResume(resume)
		assert.NoError(t, err)
	}

	for _, test := range tests {
		resume, err := repo.FindBySlug(test.slug)
		if test.enabled {
			assert.NoError(t, err)
			assert.Equal(t, test.body, resume.Body)
			assert.Equal(t, test.enabled, resume.Enabled)
			assert.Equal(t, test.slug, resume.Slug)
		} else {
			assert.Error(t, err)
		}
	}
}
