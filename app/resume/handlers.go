package resume

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/antonve/portfolio-api/ports/openapi"
)

type HTTPHandlers struct {
	resumeService ResumeService
}

func NewHTTPHandlers(resumeService ResumeService) HTTPHandlers {
	return HTTPHandlers{
		resumeService: resumeService,
	}
}

func (s HTTPHandlers) FindResumeBySlug(ctx echo.Context, slug string) error {
	r, err := s.resumeService.FindResumeBySlug(ctx.Request().Context(), slug)

	if err != nil {
		switch err {
		case ErrResumeNotFound:
			return ctx.NoContent(http.StatusNotFound)
		default:
			return err
		}
	}

	return ctx.JSON(http.StatusOK, openapi.Resume{Body: r.Body()})
}
