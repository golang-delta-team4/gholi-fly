package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	h "gholi-fly-maps/api/handlers/http"

	"gholi-fly-maps/app"
	"gholi-fly-maps/config"
)

func main() {
	// Parse flags for configuration file and server port
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

	router := h.SetupRouter(
		appContainer.TerminalService(),
		appContainer.PathService(),
	)

	// Start the HTTP server
	serverAddress := fmt.Sprintf(":%d", httpPort)
	log.Printf("Starting HTTP server on %s\n", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, router))

	// Start gRPC server
	// grpcPort := 50051
	// lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	// if err != nil {
	// 	log.Fatalf("Failed to listen: %v", err)
	// }

	// grpcServer := grpc.NewServer()
	// pb.RegisterPathServiceServer(grpcServer, grpc.NewPathServiceServer(appContainer.PathService()))

	// log.Printf("gRPC server is running on port %d", grpcPort)
	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalf("Failed to serve: %v", err)
	// }
}
