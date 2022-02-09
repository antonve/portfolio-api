package main

import (
	"fmt"
	"log"

	"github.com/antonve/portfolio-api/app"
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
		panic("migrate not yet implemented")

	default:
		panic(fmt.Sprintf("server type '%s' is not supported", serverType))
	}
}
