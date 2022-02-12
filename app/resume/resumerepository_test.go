package resume_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/antonve/portfolio-api/app/resume"
	"github.com/antonve/portfolio-api/test"
)

func TestResumeRepository_StoreResume(t *testing.T) {
	rdb, cleanup := test.GetDatabase(t)
	defer cleanup()

	repo := resume.NewResumeRepository(rdb)

	resume, _ := resume.NewResume("test", "{}", true)

	err := repo.StoreResume(context.Background(), resume)
	assert.NoError(t, err)
}
