package repositories_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/antonve/portfolio-api/domain/resume"
	"github.com/antonve/portfolio-api/repositories"
)

func TestResumeRepository_StoreResume(t *testing.T) {
	rdb, cleanup := setupTestingSuite(t)
	defer cleanup()

	repo := repositories.NewResumeRepository(rdb)

	resume, _ := resume.NewResume("test", "{}", true)

	err := repo.StoreResume(context.Background(), resume)
	assert.NoError(t, err)
}
