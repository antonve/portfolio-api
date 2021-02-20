//go:generate gex mockgen -source=context.go -package services -destination=context_mock.go

package services

import (
	"github.com/antonve/portfolio-api/domain"
)

// based on https://github.com/labstack/echo/blob/a2d4cb9c7a629e2ee21861501690741d2374de10/context.go

// Context is a subset of the echo framework context, so we are not directly depending on it
type Context interface {
	// Param returns path parameter by name.
	Param(name string) string

	// QueryParam returns the query param for the provided name.
	QueryParam(name string) string

	// IntQueryParam returns the query param for the provided name, converted to int
	IntQueryParam(name string) (uint64, error)

	// OptionalIntQueryParam returns the query param for the provided name, converted to int with a fallback
	OptionalIntQueryParam(name string, defaultValue uint64) (result uint64)

	// Get retrieves data from the context.
	Get(key string) interface{}

	// Set saves data in the context.
	Set(key string, val interface{})

	// Bind binds the request body into provided type `i`. The default binder
	// does it based on Content-Type header.
	Bind(i interface{}) error

	// String sends a string response with status code.
	String(code int, s string) error

	// NoContent sends a response with no body and a status code.
	NoContent(code int) error

	// JSON sends a JSON response with status code.
	JSON(code int, i interface{}) error

	// GetID gets the current id in the route
	GetID() (uint64, error)

	// Parses out the id in the route and binds it to the given variable
	BindID(*uint64) error

	// Returns the environment the app is running in
	Environment() domain.Environment

	// RealIP returns the client's network address based on `X-Forwarded-For`
	// or `X-Real-IP` request header.
	RealIP() string

	// UserAgent returns the client's user agent
	UserAgent() string
}
