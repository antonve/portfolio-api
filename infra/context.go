package infra

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/antonve/portfolio-api/domain"
)

type context struct {
	echo.Context
	environment domain.Environment
}

func (c context) Environment() domain.Environment {
	return c.environment
}

func (c context) GetID() (uint64, error) {
	idFromRoute := c.Param("id")
	id, err := strconv.ParseUint(idFromRoute, 10, 64)

	return id, domain.WrapError(err)
}

func (c context) BindID(id *uint64) error {
	value, err := c.GetID()
	*id = value

	return domain.WrapError(err)
}

func (c context) IntQueryParam(name string) (uint64, error) {
	param := c.QueryParam(name)
	result, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return 0, domain.WrapError(err)
	}

	return result, nil
}

func (c context) OptionalIntQueryParam(name string, defaultValue uint64) uint64 {
	result, err := c.IntQueryParam(name)
	if err != nil {
		return defaultValue
	}

	return result
}
