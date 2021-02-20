package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/antonve/portfolio-api/domain"
	"github.com/antonve/portfolio-api/interfaces/services"
	"github.com/antonve/portfolio-api/usecases"
)

func TestResumeService_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	slug := "test"
	body := "{}"
	ip := "127.0.0.1"
	agent := "Chrome"

	ctx := services.NewMockContext(ctrl)
	ctx.EXPECT().JSON(200, body)
	ctx.EXPECT().QueryParam("slug").Return(slug)
	ctx.EXPECT().RealIP().Return(ip)
	ctx.EXPECT().UserAgent().Return(agent)

	i := usecases.NewMockResumeInteractor(ctrl)
	i.EXPECT().TrackedFind(slug, ip, agent).Return(&domain.Resume{slug, body, true}, nil)

	s := services.NewResumeService(i)
	err := s.Get(ctx)

	assert.NoError(t, err)
}
