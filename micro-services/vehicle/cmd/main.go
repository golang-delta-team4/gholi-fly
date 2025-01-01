package main

import (
	"flag"
	"fmt"
	"log"

	"vehicle/api/handlers/http"
	"vehicle/app"
	"vehicle/config"
)

func main() {
	// Parse config file path
	configPath := flag.String("config", "config.json", "Path to the configuration file")
	flag.Parse()

	// Load configuration
	cfg := config.MustReadConfig(*configPath)

	// Initialize application
	appContainer, err := app.NewApp(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	// Set up Fiber router
	app := http.SetupRouter(appContainer)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.Server.HttpPort)))
}
