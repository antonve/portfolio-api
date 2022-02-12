package resume

import (
	"net/http"

	"github.com/google/uuid"
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
	visit, err := NewVisit(uuid.New().String(), slug, ctx.RealIP(), ctx.Request().Header.Get("User-Agent"))
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	r, err := s.resumeService.GetTrackedResumeVisit(ctx.Request().Context(), visit)

	if err != nil {
		switch err {
		case ErrResumeNotFound:
			return ctx.NoContent(http.StatusNotFound)
		default:
			return err
		}
	}

	return ctx.JSON(http.StatusOK, openapi.ResumeView{Body: r.Body()})
}
