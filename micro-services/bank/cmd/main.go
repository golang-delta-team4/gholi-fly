package main

import (
	"flag"
	"gholi-fly-bank/api/handlers/grpc"
	"gholi-fly-bank/app"
	"gholi-fly-bank/config"
	"log"
	"os"
	"sync"
)

var configPath = flag.String("config", "config.json", "service configuration file")

func main() {
	// Parse command-line flags
	flag.Parse()

	// Override config path from environment variable if set
	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}

	// Load configuration
	c := config.MustReadConfig(*configPath)

	// Initialize application container
	appContainer := app.NewMustApp(c)

	// Use WaitGroup to run HTTP and gRPC servers concurrently
	var wg sync.WaitGroup

	// Start gRPC server
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("Starting gRPC server on port %d", c.Server.GrpcPort)
		if err := grpc.Run(appContainer, c); err != nil {
			log.Fatalf("gRPC server error: %v", err)
		}
	}()

	// Wait for both servers to exit
	wg.Wait()
}
