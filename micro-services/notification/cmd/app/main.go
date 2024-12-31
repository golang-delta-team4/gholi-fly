package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"notification-nats/api/handlers/http"
	"notification-nats/api/pb"
	notification "notification-nats/api/service"
	"notification-nats/config"
	"notification-nats/database"
	"notification-nats/models"

	"google.golang.org/grpc"
)

var configPath = flag.String("config", "config.json", "service configuration file")

func main() {
	flag.Parse() // don't forget to parse flags

	cfg := config.MustReadConfig(*configPath)

	db, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatal("error connecting to db:", err)
	}

	if err := db.AutoMigrate(&models.OutBoxMessage{}, &models.NotificationHistory{}); err != nil {
		log.Fatal("migrate error - ", err)
	}

	// 1) Start gRPC in a goroutine (so main can do more stuff)
	go func() {
		grpcServer := grpc.NewServer()

		svc := &notification.Service{DB: db, Config: cfg}
		pb.RegisterNotificationServiceServer(grpcServer, svc)

		address := fmt.Sprintf(":%d", cfg.Server.GRPCPort)
		lis, err := net.Listen("tcp", address)
		if err != nil {
			log.Fatalf("failed to listen on port %d: %v", cfg.Server.GRPCPort, err)
		}
		log.Printf("Starting gRPC server on :%d...\n", cfg.Server.GRPCPort)

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	// 2) Run Fiber (setup.go) — blocks until exit
	if err := http.RunFiber(db, cfg); err != nil {
		log.Fatalf("Fiber server error: %v", err)
	}
}
