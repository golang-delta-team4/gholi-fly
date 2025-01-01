package main

import (
	"flag"
	"fmt"
	"log"

	h "gholi-fly-maps/api/handlers/http"
	"gholi-fly-maps/app"
	"gholi-fly-maps/config"
)

func main() {
	// Parse flags for configuration file
	configPath := flag.String("config", "config.json", "Path to the configuration file")
	flag.Parse()

	// Read configuration
	cfg := config.MustReadConfig(*configPath)

	// Initialize the application container
	appContainer := app.NewMustApp(cfg)

	// Determine HTTP port
	httpPort := cfg.Server.HttpPort
	if httpPort == 0 {
		log.Fatalf("HttpPort is not set in the configuration")
	}

	// Setup the Fiber router
	router := h.SetupRouter(
		appContainer.TerminalService(),
		appContainer.PathService(),
	)

	// Start the Fiber server
	serverAddress := fmt.Sprintf(":%d", httpPort)
	log.Printf("Starting Fiber server on %s\n", serverAddress)

	if err := router.Listen(serverAddress); err != nil {
		log.Fatalf("Failed to start Fiber server: %v", err)
	}
}
