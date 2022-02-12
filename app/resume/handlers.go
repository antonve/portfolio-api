package resume

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/antonve/portfolio-api/ports/openapi"
)

type HTTPHandlers struct {
}

func NewHTTPHandlers() HTTPHandlers {
	return HTTPHandlers{}
}

func (s HTTPHandlers) FindResumeBySlug(ctx echo.Context, slug string) error {
	return ctx.JSON(http.StatusOK, openapi.Resume{Body: "hello world"})
}
