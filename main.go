package main

import (
	"log"

	"github.com/antonve/portfolio-api/app"
	"github.com/antonve/portfolio-api/ports"
)

func main() {
	app, err := app.NewApplication()
	if err != nil {
		log.Fatalf("could not start application: %s", err)
	}

	httpServer := ports.NewHttpServer(app)

	httpServer.Start()
}
