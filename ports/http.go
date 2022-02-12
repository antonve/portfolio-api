package ports

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/antonve/portfolio-api/app"
	"github.com/antonve/portfolio-api/ports/openapi"
)

type HttpServer struct {
	app    app.Application
	server *echo.Echo
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{
		app:    app,
		server: echo.New(),
	}
}

func (h HttpServer) Start() error {
	h.server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: h.app.Config().CORSAllowedOrigins,
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	h.server.Use(middleware.Logger())

	openapi.RegisterHandlers(h.server, h.app.HTTPHandlers())
	fmt.Println(h.app.Config())
	return h.server.Start(fmt.Sprintf(":%s", h.app.Config().Port))
}
