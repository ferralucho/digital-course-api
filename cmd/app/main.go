package main

import (
	"log"

	"github.com/ferralucho/digital-course-api/config"
	"github.com/ferralucho/digital-course-api/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
