package infra

import (
	"net/http"
	"regexp"
	"strconv"

	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/antonve/portfolio-api/domain"
	"github.com/antonve/portfolio-api/interfaces/services"
	"github.com/antonve/portfolio-api/usecases"
)

// NewRouter instantiates a router
func NewRouter(
	environment domain.Environment,
	port string,
	corsAllowedOrigins []string,
	errorReporter usecases.ErrorReporter,
	routes ...services.Route,
) services.Router {
	m := &middlewares{}
	e := newEcho(environment, m, corsAllowedOrigins, errorReporter, routes...)
	return router{e, port}
}

type middlewares struct {
}

func newEcho(
	environment domain.Environment,
	m *middlewares,
	corsAllowedOrigins []string,
	errorReporter usecases.ErrorReporter,
	routes ...services.Route,
) *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = errorHandler(errorReporter)
	e.Use(newContextMiddleware(environment))
	e.Use(sentryecho.New(sentryecho.Options{}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: corsAllowedOrigins,
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	for _, route := range routes {
		e.Add(route.Method, route.Path, wrap(route, m))
	}

	return e
}

func newContextMiddleware(environment domain.Environment) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &context{c, environment}
			return next(cc)
		}
	}
}

var errorCodeRegularExpression = regexp.MustCompile("^code=([0-9]{3}).")

func errorHandler(errorReporter usecases.ErrorReporter) func(error, echo.Context) {
	return func(err error, c echo.Context) {
		c.Logger().Error(err)

		if err == middleware.ErrJWTMissing {
			c.NoContent(http.StatusUnauthorized)
			return
		}

		if match := errorCodeRegularExpression.FindStringSubmatch(err.Error()); len(match) > 1 {
			if statusCode, errInt := strconv.Atoi(match[1]); errInt == nil {
				c.NoContent(statusCode)
				return
			}
		}

		if err == domain.ErrNotFound {
			c.NoContent(http.StatusNotFound)
			return
		}

		errorReporter.Capture(err)
		c.NoContent(http.StatusInternalServerError)
	}
}

func wrap(r services.Route, m *middlewares) echo.HandlerFunc {
	handler := func(c echo.Context) error {
		return r.HandlerFunc(c.(*context))
	}

	return handler
}

type router struct {
	*echo.Echo
	port string
}

func (r router) StartListening() error {
	return r.Start(":" + r.port)
}
