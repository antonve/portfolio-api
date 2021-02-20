package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/antonve/portfolio-api/domain"
	"github.com/antonve/portfolio-api/interfaces/repositories"
)

func TestResumeRepository_StoreResume(t *testing.T) {
	sqlHandler, cleanup := setupTestingSuite(t)
	defer cleanup()

	repo := repositories.NewResumeRepository(sqlHandler)

	resume := &domain.Resume{
		Slug:    "test",
		Body:    "{}",
		Enabled: true,
	}

	err := repo.StoreResume(resume)
	assert.NoError(t, err)
}
