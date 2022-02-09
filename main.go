package main

import (
	"fmt"
	"log"

	"github.com/antonve/portfolio-api/app"
	"github.com/antonve/portfolio-api/infra"
	"github.com/antonve/portfolio-api/ports"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		log.Fatalf("could not start application: %s", err)
	}

	serverType := app.Config().ServerToRun

	switch serverType {
	case "http":
		httpServer := ports.NewHttpServer(app)
		httpServer.Start()

	case "migrate":
		m, err := infra.NewMigrator(app.RDB(), "./migrations")
		if err != nil {
			panic(fmt.Sprintf("failed to migrate: %s", err))
		}

		err = m.Run()
		if err != nil {
			panic(fmt.Sprintf("failed to migrate: %s", err))
		}

	default:
		panic(fmt.Sprintf("server type '%s' is not supported", serverType))
	}
}
