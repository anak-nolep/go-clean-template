package main

import (
	"log"

	"github.com/evrone/go-clean-template/config"
	"github.com/evrone/go-clean-template/internal/app"
)

func main() {
	log.Println("Starting ...")
	
	log.Println("Loading configuration")
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	log.Println("Running")
	app.Run(cfg)
}
