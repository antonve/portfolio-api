package app

import (
	"log"
	"sync"

	"github.com/creasty/configo"

	"github.com/antonve/portfolio-api/app/resume"
	"github.com/antonve/portfolio-api/domain"
	"github.com/antonve/portfolio-api/infra"
)

type Application interface {
	RDB() *infra.RDB
	Config() *Config

	HTTPHandlers() *HTTPHandlers
	Services() *Services
}

type HTTPHandlers struct {
	resume.HTTPHandlers
}

type Services struct {
	Resume resume.ResumeService
}

type Config struct {
	ServerToRun          string             `envconfig:"server_to_run" valid:"required" default:"http"`
	Environment          domain.Environment `envconfig:"app_env" valid:"required" default:"development"`
	DatabaseURL          string             `envconfig:"database_url" valid:"required"`
	DatabaseMaxIdleConns int                `envconfig:"database_max_idle_conns" valid:"required"`
	DatabaseMaxOpenConns int                `envconfig:"database_max_open_conns" valid:"required"`
	CORSAllowedOrigins   []string           `envconfig:"cors_allowed_origins" valid:"required"`
	// ErrorReporterDSN     string             `envconfig:"error_reporter_dsn"`
	Port     string `envconfig:"app_port" valid:"required"`
	TimeZone string `envconfig:"app_timezone" valid:"required"`
}

func NewApplication() (Application, error) {
	cfg := &Config{}
	err := configo.Load(cfg, configo.Option{})
	if err != nil {
		return nil, err
	}

	a := &app{config: cfg}
	a.Init()

	return a, nil
}

type app struct {
	config *Config

	rdb struct {
		result *infra.RDB
		once   sync.Once
	}

	services struct {
		result *Services
		once   sync.Once
	}

	httpHandlers struct {
		result *HTTPHandlers
		once   sync.Once
	}
}

func (d *app) Config() *Config {
	return d.config
}

func (d *app) Init() {}

func (d *app) RDB() *infra.RDB {
	holder := &d.rdb
	holder.once.Do(func() {
		var err error
		config := d.Config()
		holder.result, err = infra.NewRDB(config.DatabaseURL, config.DatabaseMaxIdleConns, config.DatabaseMaxOpenConns)

		if err != nil {
			log.Fatalf("failed to initialize connection pool with database: %v\n", err)
		}
	})
	return holder.result
}

func (d *app) Services() *Services {
	holder := &d.services
	holder.once.Do(func() {
		resumeRepo := resume.NewResumeRepository(d.RDB())
		holder.result = &Services{
			Resume: resume.NewResumeService(resumeRepo),
		}
	})
	return holder.result
}

func (d *app) HTTPHandlers() *HTTPHandlers {
	holder := &d.httpHandlers
	holder.once.Do(func() {
		holder.result = &HTTPHandlers{resume.NewHTTPHandlers(d.Services().Resume)}
	})
	return holder.result
}
