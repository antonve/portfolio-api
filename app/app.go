package app

import (
	"errors"
	"log"
	"sync"

	"github.com/creasty/configo"

	"github.com/antonve/portfolio-api/domain"
	"github.com/antonve/portfolio-api/infra"
)

// ServerDependencies is a dependency container for the api
type ServerDependencies interface {
	AutoConfigure() error
	Init()

	RDB() *infra.RDB
}

// NewServerDependencies instantiates all the dependencies for the api server
func NewServerDependencies() ServerDependencies {
	return &serverDependencies{}
}

type serverDependencies struct {
	Environment          domain.Environment `envconfig:"app_env" valid:"environment" default:"development"`
	DatabaseURL          string             `envconfig:"database_url" valid:"required"`
	DatabaseMaxIdleConns int                `envconfig:"database_max_idle_conns" valid:"required"`
	DatabaseMaxOpenConns int                `envconfig:"database_max_open_conns" valid:"required"`
	CORSAllowedOrigins   []string           `envconfig:"cors_allowed_origins" valid:"required"`
	// ErrorReporterDSN     string             `envconfig:"error_reporter_dsn"`
	Port     string `envconfig:"app_port" valid:"required"`
	TimeZone string `envconfig:"app_timezone" valid:"required"`

	rdb struct {
		result *infra.RDB
		once   sync.Once
	}
}

func (d *serverDependencies) AutoConfigure() error {
	return configo.Load(d, configo.Option{})
}

func (d *serverDependencies) Init() {}

func (d *serverDependencies) RDB() *infra.RDB {
	holder := &d.rdb
	holder.once.Do(func() {
		var err error
		holder.result, err = infra.NewRDB(d.DatabaseURL, d.DatabaseMaxIdleConns, d.DatabaseMaxOpenConns)

		if err != nil {
			log.Fatalf("failed to initialize connection pool with database: %v\n", err)
		}
	})
	return holder.result
}

func RunHTTPServer(d ServerDependencies) error {
	d.Init()

	return errors.New("not yet implemented")
}
